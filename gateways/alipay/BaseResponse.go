package alipay

import (
	"encoding/json"
	"encoding/xml"
	"github.com/bennya8/go-union-payment/contracts"
)

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