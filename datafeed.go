package treepay

type DataFeedItem struct {
	SiteName   string `json:"site_name"`
	PayType    string `json:"pay_type"`
	PayBrand   string `json:"pay_brand"`
	AuthDate   string `json:"auth_date"`
	TradeMoney string `json:"trade_mony"`
	Currency   string `json:"currency"`
	AuthNo     string `json:"auth_no"`
	OrderNo    string `json:"order_no"`
	TradeNo    string `json:"tno"`
	OrderName  string `json:"order_name"`
	GoodName   string `json:"good_name"`
	OrderTel   string `json:"order_tel"`

	// Added field in refund API (Only supports Credit/Debit Card)
	CancelYN        Status `json:"cancel_yn"`
	PartCancelYN    Status `json:"part_cancel_yn"`
	CancelYMD       string `json:"cancel_ymd"`
	PartCancelMoney string `json:"part_cancel_mony"`

	// Additional Field in Installment Service
	InstallmentPeriod string `json:"installment_period"`

	// Additional Field in Over the Counter Service
	IncludedCustomerFee string `json:"in_cust_fee"`
	// Reference number of Barcode
	ReferenceNo string `json"reference_no"`
	ExpiryDate  string `json:"expire_date"`
}
