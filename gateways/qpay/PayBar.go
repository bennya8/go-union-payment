package qpay

import (
	"github.com/bennya8/go-union-payment/payloads"
)

type PayBar struct {
	Base *Base
}

func (p PayBar) Request(params map[string]string) *payloads.UnionPaymentResult {
	panic("implement me")
}

func (p PayBar) BuildParams(params map[string]string) map[string]string {
	panic("implement me")
}
