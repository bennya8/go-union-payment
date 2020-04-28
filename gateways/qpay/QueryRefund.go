package qpay

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

type QueryRefund struct {
	Base *Base
}

func (q QueryRefund) Request(params map[string]string) *payloads.UnionPaymentResult {
	uri := q.Base.GetFullGatewayUrl("pay/qpay_refund_query.cgi")

	//api

	resp, err := q.Base.Request(uri, q.BuildParams(params))
	return payloads.NewUnionPaymentResult(err == nil, fmt.Sprintf("%s", err), resp)
}

func (q QueryRefund) BuildParams(params map[string]string) map[string]string {
	return map[string]string{

	}
}
