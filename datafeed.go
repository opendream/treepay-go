package treepay

type DataFeedItem struct {
	SiteName   string `form:"site_name"`
	PayType    string `form:"pay_type"`
	PayBrand   string `form:"pay_brand"`
	AuthDate   string `form:"auth_date"`
	TradeMoney string `form:"trade_mony"`
	Currency   string `form:"currency"`
	AuthNo     string `form:"auth_no"`
	OrderNo    string `form:"order_no"`
	TradeNo    string `form:"tno"`
	OrderName  string `form:"order_name"`
	GoodName   string `form:"good_name"`
	OrderTel   string `form:"order_tel"`

	// Added field in refund API (Only supports Credit/Debit Card)
	CancelYN        Status `form:"cancel_yn"`
	PartCancelYN    Status `form:"part_cancel_yn"`
	CancelYMD       string `form:"cancel_ymd"`
	PartCancelMoney string `form:"part_cancel_mony"`

	// Additional Field in Installment Service
	InstallmentPeriod string `form:"installment_period"`

	// Additional Field in Over the Counter Service
	IncludedCustomerFee string `form:"in_cust_fee"`
	// Reference number of Barcode
	ReferenceNo string `form:"reference_no"`
	ExpiryDate  string `form:"expire_date"`
}
