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
type CloseTrade struct {
	Base *Base
}

func (w CloseTrade) Request(params map[string]string) *payloads.UnionPaymentResult {
	uri := w.Base.GetFullGatewayUrl("pay/closeorder")

	//api

	resp, err := w.Base.Request(uri, w.BuildParams(params))
	return payloads.NewUnionPaymentResult(err != nil, fmt.Sprintf("%s", err), resp)
}

func (w CloseTrade) BuildParams(params map[string]string) map[string]string {
	ret := map[string]string{
		"out_trade_no":   params["out_trade_no"],
	}
	for k, v := range params {
		ret[k] = v
	}
	return ret
}
