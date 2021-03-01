package alipay

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

type Gateway struct {
	Base *Base
}

func (a *Gateway) Request(api payloads.UnionPaymentApi, params map[string]string) *payloads.UnionPaymentResult {
	switch api {
	case payloads.AliApiPayApp:
		return (&PayApp{Base: a.Base}).Request(params)
	case payloads.AliApiPayBar:
		return (&PayBar{Base: a.Base}).Request(params)
	case payloads.AliApiPayQr:
		return (&PayQr{Base: a.Base}).Request(params)
	case payloads.AliApiPayWap:
		return (&PayWap{Base: a.Base}).Request(params)
	case payloads.AliApiPayWeb:
		return (&PayWeb{Base: a.Base}).Request(params)
	}
	return payloads.NewUnionPaymentResult(false, "unknown gateway api", nil)
}

type Base struct {
	Config     *Config
	GatewayUrl string
}

func NewBase(config *Config) *Base {
	b := &Base{}
	b.Config = config

	return b
}

func (b *Base) GetFullGatewayUrl(method string) string {
	return fmt.Sprintf(b.GatewayUrl, method)
}

func (b *Base) Request(uri string, params map[string]string) (*BaseResponse, error) {
	return NewBaseResponse("TODO"), nil
}
