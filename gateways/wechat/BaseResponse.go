package wechat

type WxNotifyPayResponse struct {
	AppId         string `json:"appid"`
	Attach        string `json:"attach"`
	BankType      string `json:"bank_type"`
	CouponCount   int    `json:"coupon_count"`
	CouponFee     int    `json:"coupon_fee"`
	CouponId      int    `json:"coupon_id"`
	CouponType    string `json:"coupon_type"`
	IsSubscribe   string `json:"is_subscribe"`
	MchId         int    `json:"mch_id"`
	NonceStr      string `json:"nonce_str"`
	Openid        string `json:"openid"`
	OutTradeNo    string `json:"out_trade_no"`
	ResultCode    string `json:"result_code"`
	ReturnCode    string `json:"return_code"`
	Sign          string `json:"sign"`
	TimeEnd       int    `json:"time_end"`
	TotalFee      int    `json:"total_fee"`
	TradeType     string `json:"trade_type"`
	TransactionId int    `json:"transaction_id"`
}

type WxNotifyRefundResponse struct {
	AppId      string `json:"appid"`
	MchId      int    `json:"mch_id"`
	NonceStr   string `json:"nonce_str"`
	ReturnCode string `json:"return_code"`
	ReqInfo    string `json:"req_info"`
}