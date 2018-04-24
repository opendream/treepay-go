package treepay

type PaymentRequest struct {
	// return customer to web page at this url
	ReturnUrl string `json:"ret_url,omitempty"`

	UserID     string `json:"user_id,omitempty"`
	OrderNo    string `json:"order_no,omitempty"`
	GoodName   string `json:"good_name,omitempty"`
	TradeMoney uint64 `json:"trade_mony,omitempty"`
	// order
	OrderFirstName string `json:"order_first_name,omitempty"`
	OrderLastName  string `json:"order_last_name,omitempty"`
	OrderCity      string `json:"order_city,omitempty"`
	OrderAddress   string `json:"order_addr,omitempty"`
	OrderCountry   string `json:"order_country,omitempty"`
	OrderEmail     string `json:"order_email,omitempty"`
	OrderTel       string `json:"order_tel,omitempty"`
	OrderPostCode  string `json:"order_post_code,omitempty"`
	// receiver
	ReceiverFirstName string `json:"recv_first_name,omitempty"`
	ReceiverLastName  string `json:"recv_last_name,omitempty"`
	ReceiverCity      string `json:"recv_city,omitempty"`
	ReceiverAddress   string `json:"recv_addr,omitempty"`
	ReceiverCountry   string `json:"recv_country,omitempty"`
	ReceiverEmail     string `json:"recv_email,omitempty"`
	ReceiverTel       string `json:"recv_tel,omitempty"`
	ReceiverPostCode  string `json:"recv_post_code,omitempty"`
	// site info
	PaymentType         PaymentType `json:"pay_type,omitempty"`
	Currency            int         `json:"currency,omitempty"`
	TreepayLanguageFlag string      `json:"tp_langFlag,omitempty"`
	SiteCode            string      `json:"site_cd,omitempty"`
	HashData            string      `json:"hash_data,omitempty"`
	// additional
	AgencyGroupCode string `json:"agency_cd,omitempty"`
	StationCode     string `json:"station_cd,omitempty"`
	// data feed url
	BackURL string `json:"back_url,omitempty"`
}

type PaymentType string

const (
	PaymentTypeCreditDebitCard  PaymentType = "PACA"
	PaymentTypeInternetBanking  PaymentType = "PABK"
	PaymentTypeInstallment      PaymentType = "PAIN"
	PaymentTypeRecurringPayment PaymentType = "PARC"
	PaymentTypeOverTheCounter   PaymentType = "PAOF"
)

type PaymentStatus struct {
	OrderNo    string `json:"order_no"`
	TradeNo    string `json:"tno"`
	TradeMoney string `json:"trade_mony"`
	// The date of the transaction approval
	TradeYMD string `json:"trade_ymd"`
	// The time of the transaction approval
	TradeHMS string `json:"trade_hms"`

	AuthNo  string `json:"auth_no"`
	AuthYMD string `json:"auth_ymd"`
	AuthHMS string `json:"auth_hms"`
}
