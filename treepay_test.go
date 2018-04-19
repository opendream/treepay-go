package treepay

import (
	"context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTreepay_Hash(t *testing.T) {
	b := &BackendConfiguration{}
	p := &Params{
		PaymentRequest: &PaymentRequest{
			PaymentType: PaymentTypeCreditDebitCard,
			OrderNo:     "tp-20180418-0001",
			TradeMoney:  250000,
			SiteCode:    "treepaytest",
			UserID:      "u0001",
		},
		SecureKey: "secure-key-for-test",
	}

	hashed, err := b.Hash(p)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)
	// from command line (mac) `printf "...text here..." | shasum -a 256`
	assert.Equal(t, "0dc1d5f6d7262567a44ef18b56082a7b76771c2051e9a6e2fdd29ffd21ef9254", hashed)
}

func TestTreepay_NewRequest(t *testing.T) {
	b := &BackendConfiguration{
		URL: apiURL,
	}
	r, err := b.NewRequest(http.MethodPost, "/", nil)
	assert.NoError(t, err)
	if assert.NotNil(t, r) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, apiURL+"/", r.URL.String())
	}

	r, err = b.NewRequest(http.MethodPost, "/", &Params{
		PaymentRequest:    &PaymentRequest{},
		ShouldSignRequest: true,
	})
	assert.NoError(t, err)

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(body), "hash_data")
}

func TestTreepay_RequestTimeout(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(15 * time.Millisecond)
		writer.Write([]byte(`{}`))
	}))
	b := &BackendConfiguration{
		URL: ts.URL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Millisecond,
		},
	}

	var err error
	go func() {
		res := map[string]interface{}{}
		err = b.Call(http.MethodPost, "/", &Params{}, &res)
	}()

	select {
	case <-time.After(20 * time.Millisecond):
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "request canceled")
		}
	}
}

func TestTreepay_RequestTimeoutWithParamsContext(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(10 * time.Millisecond)
		writer.Write([]byte(`{}`))
	}))
	b := &BackendConfiguration{
		URL: ts.URL,
		HTTPClient: &http.Client{
			Timeout: 20 * time.Millisecond,
		},
	}

	var err error
	go func() {
		res := map[string]interface{}{}

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Millisecond)
		err = b.Call(http.MethodPost, "/", &Params{
			Context: ctx,
		}, &res)
	}()

	select {
	case <-time.After(15 * time.Millisecond):
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "context deadline")
		}
	}
}

func TestTreepay_Response(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		switch request.URL.Path {
		case "/ok":
			writer.Write([]byte(`{"res_cd":"0000","res_msg":"success"}`))
		case "/jsonerror":
			writer.Write([]byte(`{"thisisnotvalidjson"}`))
		case "/apierror":
			writer.Write([]byte(`{"res_cd":"A101","res_msg":"Invalid request data"}`))
		}
	}))
	b := &BackendConfiguration{URL: ts.URL, HTTPClient: &http.Client{}}

	res := &APIResponse{}
	err := b.Call(http.MethodPost, "/ok", &Params{}, &res)
	assert.NoError(t, err)
	if assert.NotNil(t, res) {
		assert.Equal(t, "0000", res.Code)
		assert.Equal(t, "success", res.Message)
	}

	err = b.Call(http.MethodPost, "/jsonerror", &Params{}, &res)
	assert.Error(t, err)

	err = b.Call(http.MethodPost, "/apierror", &Params{}, &res)
	assert.Error(t, err)
	assert.IsType(t, &APIError{}, err)
}

func TestTreepay_NewBackendConfiguration(t *testing.T) {
	b := NewBackendConfiguration(apiURLTest)
	assert.Equal(t, apiURLTest, b.URL)
	assert.NotNil(t, b.HTTPClient)
}
