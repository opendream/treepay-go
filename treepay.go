package treepay

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	apiURL     = "https://pay.treepay.co.th/api"
	apiURLTest = "https://paytest.treepay.co.th/api"
)

const (
	APIURL     = "https://pay.treepay.co.th/api"
	APIURLTest = "https://paytest.treepay.co.th/api"
)

const apiVersion = "1.5.0"

const clientVersion = "1.0.0"

// default http timeout when communicate with Treepay api server
const defaultHTTPTimeout = 60 * time.Second

const DefaultAgencyGroupCode = "AGMF"

type Backend interface {
	Call(method, path string, params *Params, v interface{}) error
}

type BackendConfiguration struct {
	URL        string
	HTTPClient *http.Client
}

type APIResponse struct {
	Code    string `json:"res_cd"`
	Message string `json:"res_msg"`
}

func (b *BackendConfiguration) Call(method, path string, params *Params, v interface{}) error {
	req, err := b.NewRequest(method, path, params)
	if err != nil {
		return err
	}

	res, err := b.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return &APIConnectionError{
			treepayErr: &Error{
				HTTPResponse: res,
			},
		}
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// try to parse body as Error and detect whether if error or not.
	treepayErr := Error{
		HTTPResponse: res,
	}
	if err := json.Unmarshal(body, &treepayErr); err == nil && treepayErr.Code != "" {
		if treepayErr.Code != "0000" {
			return &APIError{
				treepayErr: &treepayErr,
			}
		}
	}

	if err := json.Unmarshal(body, v); err != nil {
		return err
	}

	return nil
}

func (b *BackendConfiguration) NewRequest(method, path string, params *Params) (*http.Request, error) {
	url := b.URL + path

	body := &strings.Reader{}
	if params != nil && params.PaymentRequest != nil {
		if params.PaymentRequest.AgencyGroupCode == "" {
			params.PaymentRequest.AgencyGroupCode = DefaultAgencyGroupCode
		}

		if params.ShouldSignRequest {
			hashString, err := b.Hash(params)
			if err != nil {
				return nil, err
			}

			params.PaymentRequest.HashData = hashString
		}

		jsonString, err := json.Marshal(params.PaymentRequest)
		if err != nil {
			return nil, err
		}

		body = strings.NewReader(string(jsonString))
	}

	r, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	if params != nil && params.Context != nil {
		r = r.WithContext(params.Context)
	}

	r.Header.Set("Content-Type", "application/json; charset=utf8")

	return r, nil
}

func (b *BackendConfiguration) Hash(params *Params) (string, error) {
	p := params.PaymentRequest
	key := params.SecureKey

	hashString := ""

	hashString += fmt.Sprintf("%s", p.PaymentType) +
		p.OrderNo +
		fmt.Sprintf("%d", p.TradeMoney) +
		p.SiteCode +
		key +
		p.UserID +
		p.AgencyGroupCode
	signed := sha256.Sum256([]byte(hashString))

	return fmt.Sprintf("%x", signed), nil
}

func NewBackendConfiguration(url string) BackendConfiguration {
	return BackendConfiguration{
		URL: url,
		HTTPClient: &http.Client{
			Timeout: defaultHTTPTimeout,
		},
	}
}
