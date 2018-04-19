package overthecounter

import (
	"encoding/json"
	"fmt"
	"github.com/opendream/treepay-go"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient(&treepay.BackendConfiguration{}, "sitetest", "secret-key")
	assert.Equal(t, "secret-key", c.Key)
}

func TestClient_RequestSuccess(t *testing.T) {
	count := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		p := treepay.PaymentRequest{}
		if err := json.Unmarshal(body, &p); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		assert.NotEmpty(t, p.HashData)
		assert.Equal(t, "tp-20180418-0001", p.OrderNo)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"tno":"10002000300%d"}`, count)))
		count++
	}))
	b := treepay.NewBackendConfiguration(ts.URL)
	c := NewClient(&b, "sitetest", "secret-key")

	// success
	res, err := c.Request(&treepay.PaymentRequest{OrderNo: "tp-20180418-0001"})
	assert.NoError(t, err)
	if assert.NotNil(t, res) {
		assert.Equal(t, "100020003000", res.TradeNo)
	}

	// second time
	res, err = c.Request(&treepay.PaymentRequest{OrderNo: "tp-20180418-0001"})
	assert.NoError(t, err)
	if assert.NotNil(t, res) {
		assert.Equal(t, "100020003001", res.TradeNo)
	}
}

func TestClient_RequestError(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"res_cd":"A101","res_msg":"Invalid request data"}`))
	}))

	b := treepay.NewBackendConfiguration(ts.URL)
	c := NewClient(&b, "sitetest", "secret-key")

	_, err := c.Request(&treepay.PaymentRequest{})
	if assert.Error(t, err) {
		assert.IsType(t, &treepay.APIError{}, err)

		apiError := err.(*treepay.APIError)
		assert.Equal(t, "A101: Invalid request data", apiError.Error())
	}
}
