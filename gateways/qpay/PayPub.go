package qpay

import (
	"github.com/bennya8/go-union-payment/payloads"
)

type PayPub struct {
	Base *Base
}

func (p PayPub) Request(params map[string]string) *payloads.UnionPaymentResult {
	panic("implement me")
}

func (p PayPub) BuildParams(params map[string]string) map[string]string {
	panic("implement me")
}
