package wechat

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

/**
 * @package gateway.wechat
 * @author  : Benny
 * @email   : benny_a8@qq.com
 * @date    : 2020/04/11
 **/
type Coupon struct {
	Base *Base
}

func (w Coupon) Request(params map[string]string) *payloads.UnionPaymentResult {
	uri := w.Base.GetFullGatewayUrl("mmpaymkttransfers/send_coupon")

	//api

	resp, err := w.Base.Request(uri, w.BuildParams(params))
	return payloads.NewUnionPaymentResult(err == nil, fmt.Sprintf("%s", err), resp)
}

func (w Coupon) BuildParams(params map[string]string) map[string]string {
	ret := map[string]string{
		"coupon_stock_id":  "",
		"openid_count":     "1",
		"partner_trade_no": "",
		"openid":           "",
		"op_user_id":       "",
		"device_info":      "",
		"version":          "1.0",
		"type":             "XML",
	}
	for k, v := range params {
		ret[k] = v
	}
	return ret
}
