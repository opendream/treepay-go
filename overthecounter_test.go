package treepay

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOverTheCounterAPIResponse_String(t *testing.T) {
	r := OverTheCounterAPIResponse{
		TradeNo:    "100020003000",
		TradeMoney: "2500.00",
		TradeYMD:   "20180418",
		TradeHMS:   "170459",
	}

	marshalled, err := json.Marshal(r)
	assert.NoError(t, err)
	if assert.NotEmpty(t, marshalled) {
		unmarshalled := OverTheCounterAPIResponse{}
		err := json.Unmarshal(marshalled, &unmarshalled)

		assert.NoError(t, err)
		assert.Equal(t, r.TradeNo, unmarshalled.TradeNo)
		assert.Equal(t, r.TradeMoney, unmarshalled.TradeMoney)
		assert.Equal(t, r.TradeYMD, unmarshalled.TradeYMD)
		assert.Equal(t, r.TradeHMS, unmarshalled.TradeHMS)
	}
}
