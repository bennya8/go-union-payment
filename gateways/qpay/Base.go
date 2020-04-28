package qpay

import (
	"github.com/bennya8/go-union-payment/payloads"
)

type Gateway struct {
	Base *Base
}

func (a *Gateway) Request(api payloads.UnionPaymentApi, params map[string]string) *payloads.UnionPaymentResult {
	switch api {
	case payloads.QpayApiPayApp:
	case payloads.QpayApiPayBar:
	case payloads.QpayApiPayPub:
	case payloads.QpayApiPayQr:
	}
	return payloads.NewUnionPaymentResult(false, "unknown gateway api", nil)
}

type Base struct {
	Config *Config
}

func NewBase(config *Config) *Base {
	b := &Base{}
	b.Config = config

	return b
}
