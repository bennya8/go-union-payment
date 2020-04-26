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
type CancelTrade struct {
	Base *Base
}

func (w CancelTrade) Request(params map[string]string) *payloads.UnionPaymentResult {
	uri := w.Base.GetFullGatewayUrl("secapi/pay/reverse")

	//api

	resp, err := w.Base.Request(uri, w.BuildParams(params))
	return payloads.NewUnionPaymentResult(err == nil, fmt.Sprintf("%s", err), resp)
}

func (w CancelTrade) BuildParams(params map[string]string) map[string]string {
	ret := map[string]string{
		"transaction_id": params["transaction_id"],
		"out_trade_no":   params["out_trade_no"],
	}
	for k, v := range params {
		ret[k] = v
	}
	return ret
}
