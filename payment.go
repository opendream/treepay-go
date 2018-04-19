package treepay

type PaymentRequest struct {
	// return customer to web page at this url
	ReturnUrl string `json:"ret_url"`

	UserID     string `json:"user_id,omitempty"`
	OrderNo    string `json:"order_no"`
	GoodName   string `json:"good_name"`
	TradeMoney uint64 `json:"trade_mony"`
	// order
	OrderFirstName string `json:"order_first_name"`
	OrderLastName  string `json:"order_last_name"`
	OrderCity      string `json:"order_city"`
	OrderAddress   string `json:"order_addr"`
	OrderCountry   string `json:"order_country"`
	OrderEmail     string `json:"order_email"`
	OrderTel       string `json:"order_tel"`
	OrderPostCode  string `json:"order_post_code"`
	// receiver
	ReceiverFirstName string `json:"recv_first_name"`
	ReceiverLastName  string `json:"recv_last_name"`
	ReceiverCity      string `json:"recv_city"`
	ReceiverAddress   string `json:"recv_addr"`
	ReceiverCountry   string `json:"recv_country"`
	ReceiverEmail     string `json:"recv_email"`
	ReceiverTel       string `json:"recv_tel"`
	ReceiverPostCode  string `json:"recv_post_code"`
	// site info
	PaymentType         PaymentType `json:"pay_type"`
	Currency            int         `json:"currency"`
	TreepayLanguageFlag string      `json:"tp_langFlag"`
	SiteCode            string      `json:"site_cd"`
	HashData            string      `json:"hash_data"`
	// additional
	AgencyGroupCode string `json:"agency_cd"`
	StationCode     string `json:"station_cd"`
	// data feed url
	BackURL string `json:"back_url"`
}

type PaymentType string

const (
	PaymentTypeCreditDebitCard  PaymentType = "PACA"
	PaymentTypeInternetBanking  PaymentType = "PABK"
	PaymentTypeInstallment      PaymentType = "PAIN"
	PaymentTypeRecurringPayment PaymentType = "PARC"
	PaymentTypeOverTheCounter   PaymentType = "PAOF"
)
