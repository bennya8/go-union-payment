package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
	"github.com/bennya8/go-union-payment/utils/date"
	"strconv"
	"time"
)

type PayPub struct {
	Base *Base
}

func (w PayPub) Request(params map[string]string) *payloads.UnionPaymentResult {
	uri := w.Base.GetFullGatewayUrl("pay/unifiedorder")
	resp, err := w.Base.Request(uri, w.BuildParams(params))
	if err != nil {
		return payloads.NewUnionPaymentResult(false, fmt.Sprintf("%s", err), nil)
	}
	return payloads.NewUnionPaymentResult(true, "ok", w.ParseResult(resp))
}

func (w PayPub) ParseResult(response payloads.IGatewayResponse) map[string]string {
	var retData = map[string]string{}
	retMap, err := response.ToMap()
	if err != nil {
		return retData
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	signature := w.Base.BuildSign(timestamp, retMap)

	return map[string]string{
		"appId":     retMap["appid"].(string),
		"nonceStr":  retMap["nonce_str"].(string),
		"package":   "prepay_id=" + retMap["prepay_id"].(string),
		"paySign":   w.Base.MakeSign(signature),
		"signType":  w.Base.Config.SignType,
		"timeStamp": timestamp,
	}
}

func (w PayPub) BuildParams(params map[string]string) map[string]string {
	var notifyUrl = ""
	if len(w.Base.Config.NotifyUrl) > 0 {
		notifyUrl = w.Base.Config.NotifyUrl
	}

	var limitPay = ""
	if len(w.Base.Config.LimitPay) > 0 {
		limitPay = w.Base.Config.LimitPay[0]
	}

	var feeType = "CNY"
	if len(w.Base.Config.FeeType) > 0 {
		feeType = w.Base.Config.FeeType
	}

	amount, err := strconv.ParseFloat(params["amount"], 64)
	if err != nil {
		panic("union.payment.error:" + err.Error())
	}
	totalFee := int(amount * 100)

	timeStart := date.TimeFormat(time.Now(), "YYYYMMDDhhiiss")
	timeExpire := date.TimeFormat(time.Now().Add(30*time.Minute), "YYYYMMDDhhiiss")
	if len(params["time_expire"]) > 0 {
		expire, _ := strconv.ParseInt(params["time_expire"], 10, 64)
		expireTime := time.Unix(expire, 0)
		timeExpire = date.TimeFormat(expireTime, "YYYYMMDDhhiiss")
	}

	receipt, err := strconv.ParseBool(params["receipt"])
	if err != nil {
		receipt = false
	}
	receiptStr := ""
	if receipt {
		receiptStr = "Y"
	}

	var sceneInfo interface{}
	storeInfoBytes := ""

	if len(params["scene_info"]) > 0 {
		_ = json.Unmarshal([]byte(params["scene_info"]), &sceneInfo)
		storeInfo := map[string]interface{}{
			"scene_info": sceneInfo,
		}
		storeInfoBytesTmp, err := json.Marshal(storeInfo)
		if err != nil {
			storeInfoBytes = string(storeInfoBytesTmp)
		}
	}

	ret := map[string]string{
		"device_info":      params["device_info"],
		"body":             params["subject"],
		"detail":           params["body"],
		"attach":           params["return_param"],
		"out_trade_no":     params["trade_no"],
		"fee_type":         feeType,
		"total_fee":        fmt.Sprintf("%d", totalFee),
		"spbill_create_ip": params["client_ip"],
		"time_start":       timeStart,
		"time_expire":      timeExpire,
		"goods_tag":        params["goods_tag"],
		"notify_url":       notifyUrl,
		"trade_type":       "JSAPI",
		"limit_pay":        limitPay,
		"receipt":          receiptStr,
		"scene_info":       storeInfoBytes,
		"openid":           params["openid"],
	}

	return ret
}
