package wechat

import (
	"encoding/json"
	"encoding/xml"
	"github.com/bennya8/go-union-payment/contracts"
)

type WxNotifyPayResponse struct {
	AppId         string `json:"appid" xml:"appid"`
	Attach        string `json:"attach" xml:"attach"`
	BankType      string `json:"bank_type" xml:"bank_type"`
	CouponCount   string `json:"coupon_count" xml:"coupon_count"`
	CouponFee     string `json:"coupon_fee" xml:"coupon_fee"`
	CouponId      string `json:"coupon_id" xml:"coupon_id"`
	CouponType    string `json:"coupon_type" xml:"coupon_type"`
	IsSubscribe   string `json:"is_subscribe" xml:"is_subscribe"`
	MchId         string `json:"mch_id" xml:"mch_id"`
	NonceStr      string `json:"nonce_str" xml:"nonce_str"`
	Openid        string `json:"openid" xml:"openid"`
	OutTradeNo    string `json:"out_trade_no" xml:"out_trade_no"`
	ResultCode    string `json:"result_code" xml:"result_code"`
	ReturnCode    string `json:"return_code" xml:"return_code"`
	Sign          string `json:"sign" xml:"sign"`
	TimeEnd       string `json:"time_end" xml:"time_end"`
	TotalFee      string `json:"total_fee" xml:"total_fee"`
	TradeType     string `json:"trade_type" xml:"trade_type"`
	TransactionId string `json:"transaction_id" xml:"transaction_id"`
}

type WxNotifyRefundResponse struct {
	AppId      string `json:"appid" xml:"app_id"`
	MchId      string `json:"mch_id" xml:"mch_id"`
	NonceStr   string `json:"nonce_str" xml:"nonce_str"`
	ReturnCode string `json:"return_code" xml:"return_code"`
	ReqInfo    string `json:"req_info" xml:"req_info"`
}

type BaseResponse struct {
	Resp string
}

func NewBaseResponse(resp string) *BaseResponse {
	return &BaseResponse{Resp: resp}
}

func (w *BaseResponse) ToMap() (map[string]interface{}, error) {
	var retKvMap map[string]string
	err := xml.Unmarshal([]byte(w.Resp), (*contracts.XmlParams)(&retKvMap))

	var ret map[string]interface{}
	retBytes, err := json.Marshal(retKvMap)
	if err != nil {
		return ret, nil
	}
	err = json.Unmarshal(retBytes, &ret)
	if err != nil {
		return ret, nil
	}
	return ret, err
}

func (w *BaseResponse) ToJson() (string, error) {
	dict, err := w.ToMap()
	if err != nil {
		return "", err
	}
	ret, err := json.Marshal(dict)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func (w *BaseResponse) ToXml() (string, error) {
	// default is xml, no need to convert
	return w.Resp, nil
}
