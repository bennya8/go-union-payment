package alipay

import "github.com/bennya8/go-union-payment/payloads"

type PayWeb struct {
	Base *Base
}

func (p PayWeb) Request(params map[string]string) *payloads.UnionPaymentResult {
	panic("implement me")
}

func (p PayWeb) BuildParams(params map[string]string) map[string]string {
	panic("implement me")
}
