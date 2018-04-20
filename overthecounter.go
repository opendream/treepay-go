package treepay

type OverTheCounterAPIResponse struct {
	TradeNo    string `json:"tno"`
	TradeMoney string `json:"trade_mony"`
	// The date of the transaction approval
	TradeYMD string `json:"trade_ymd"`
	// The time of the transaction approval
	TradeHMS string `json:"trade_hms"`
	RefNo    string `json:"ref_no"`
	OrderNo  string `json:"order_no"`
	// Payment result URL that a customer can check Reference number and Barcode
	InfoURL string `json:"info_url"`
}
