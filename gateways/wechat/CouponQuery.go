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
type CouponQuery struct {
	Base *Base
}

func (w CouponQuery) Request(params map[string]string) *payloads.UnionPaymentResult {
	uri := w.Base.GetFullGatewayUrl("mmpaymkttransfers/querycouponsinfo")

	//api

	resp, err := w.Base.Request(uri, w.BuildParams(params))
	return payloads.NewUnionPaymentResult(err != nil, fmt.Sprintf("%s", err), resp)
}

func (w CouponQuery) BuildParams(params map[string]string) map[string]string {
	ret := map[string]string{
		"coupon_id":  "",
		"openid":     "1",
		"stock_id":   "",
		"op_user_id": "",
		"version":    "1.0",
		"type":       "XML",
	}
	for k, v := range params {
		ret[k] = v
	}
	return ret
}
