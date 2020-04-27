package alipay

import "github.com/bennya8/go-union-payment/payloads"

type PayWap struct {
	Base *Base
}

func (p PayWap) Request(params map[string]string) *payloads.UnionPaymentResult {
	panic("implement me")
}

func (p PayWap) BuildParams(params map[string]string) map[string]string {
	panic("implement me")
}

