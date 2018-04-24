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

func TestPaymentStatus_MarshalJSON(t *testing.T) {
	r := &PaymentStatus{
		OrderNo:     "tp-20180418-0001",
		TradeNo:     "100020003000",
		TradeMoney:  "2500.00",
		TradeYMD:    "20180418",
		TradeStatus: TradeStatusWaitForPayment,
	}

	marshalled, err := json.Marshal(r)
	assert.NoError(t, err)
	assert.NotEmpty(t, marshalled)

	unmarshalled := &PaymentStatus{}
	err = json.Unmarshal(marshalled, unmarshalled)

	assert.NoError(t, err)
	assert.Equal(t, r.OrderNo, unmarshalled.OrderNo)
	assert.Equal(t, r.TradeMoney, unmarshalled.TradeMoney)
	assert.Equal(t, r.TradeYMD, unmarshalled.TradeYMD)
	assert.Equal(t, r.TradeHMS, unmarshalled.TradeHMS)
	assert.Equal(t, r.TradeStatus, unmarshalled.TradeStatus)
}
