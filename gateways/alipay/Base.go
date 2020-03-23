package alipay

import (
	"github.com/bennya8/go-union-payment/payloads"
)

type Gateway struct {
	Base *Base
}

func (a *Gateway) Request(api payloads.UnionPaymentApi, params map[string]string) *payloads.UnionPaymentResult {
	panic("implement me")
}

type Base struct {
	Config Config
}

func NewBase(config Config) *Base {
	b := &Base{}
	b.Config = config

	return b
}