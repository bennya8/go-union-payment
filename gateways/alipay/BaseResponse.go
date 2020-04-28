package alipay

type AliNotifyPayResponse struct {
	NotifyType             string `json:"notify_type"`
	NotifyId               string `json:"notify_id"`
	NotifyTime             string `json:"notify_time"`
	SignType               string `json:"sign_type"`
	Sign                   string `json:"sign"`
	TransCurrency          string `json:"trans_currency"`
	SettleCurrency         string `json:"settle_currency"`
	PayCurrency            string `json:"pay_currency"`
	PayAmount              string `json:"pay_amount"`
	SettleTransRate        string `json:"settle_trans_rate"`
	TransPayRate           string `json:"trans_pay_rate"`
	RefundResetPaytoolList string `json:"refund_reset_paytool_list"`
	ChargeAmount           string `json:"charge_amount"`
	ChargeFlags            string `json:"charge_flags"`
	SettlementId           string `json:"settlement_id"`
	AdvanceAmount          string `json:"advance_amount"`
	NotifyActionType       string `json:"notify_action_type"`
	DiscountAmount         string `json:"discount_amount"`
	MdiscountAmount        string `json:"mdiscount_amount"`
	UnfreezeAmount         string `json:"unfreeze_amount"`
	AuthTradePayMode       string `json:"auth_trade_pay_mode"`
}
