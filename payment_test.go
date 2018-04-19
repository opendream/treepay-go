package treepay

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaymentRequest_MarshalJSON(t *testing.T) {
	r := &PaymentRequest{
		OrderNo:     "tp-20180418-0001",
		PaymentType: PaymentTypeCreditDebitCard,
	}

	marshalled, err := json.Marshal(r)
	assert.NoError(t, err)
	assert.NotEmpty(t, marshalled)

	unmarshalled := &PaymentRequest{}
	err = json.Unmarshal(marshalled, unmarshalled)
	assert.NoError(t, err)

	assert.Equal(t, r.OrderNo, unmarshalled.OrderNo)
	assert.Equal(t, r.PaymentType, unmarshalled.PaymentType)
}
