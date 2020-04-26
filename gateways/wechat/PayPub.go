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

	//api

	resp, err := w.Base.Request(uri, w.BuildParams(params))
	return payloads.NewUnionPaymentResult(err == nil, fmt.Sprintf("%s", err), resp)
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
	_ = json.Unmarshal([]byte(params["scene_info"]), &sceneInfo)

	storeInfo := map[string]interface{}{
		"scene_info": sceneInfo,
	}
	storeInfoBytes, _ := json.Marshal(storeInfo)

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
		"scene_info":       string(storeInfoBytes),
		"openid":           params["openid"],
	}

	return ret
}
