package overthecounter

import (
	"github.com/opendream/treepay-go"
	"net/http"
)

type Client struct {
	Backend treepay.Backend
	Key     string
}

func (c Client) Request(req *treepay.PaymentRequest) (*treepay.OverTheCounterAPIResponse, error) {
	params := treepay.Params{
		PaymentRequest:    req,
		ShouldSignRequest: true,
		SecureKey:         c.Key,
	}

	resp := treepay.OverTheCounterAPIResponse{}
	if err := c.Backend.Call(http.MethodPost, "/offlineReq.api", &params, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func NewClient(b treepay.Backend, key string) Client {
	return Client{Backend: b, Key: key}
}
