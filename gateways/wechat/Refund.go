package wechat

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

/**
 * @package gateway.wechat
 * @author  : Benny
 * @email   : benny_a8@qq.com
 * @date    : 2021/03/13
 **/
type Refund struct {
	Base *Base
}

func (w Refund) Request(params map[string]string) *payloads.UnionPaymentResult {
	uri := w.Base.GetFullGatewayUrl("secapi/pay/refund")

	//api

	resp, err := w.Base.Request(uri, w.BuildParams(params))
	return payloads.NewUnionPaymentResult(err == nil, fmt.Sprintf("%s", err), resp)
}

func (w Refund) BuildParams(params map[string]string) map[string]string {
	ret := map[string]string{
	}
	// @todo
	return ret
}
