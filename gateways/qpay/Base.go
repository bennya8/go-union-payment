package qpay

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

type Gateway struct {
	Base *Base
}

func (a *Gateway) Request(api payloads.UnionPaymentApi, params map[string]string) *payloads.UnionPaymentResult {
	switch api {
	// pay apis
	case payloads.QpayApiPayApp:
		return (&PayApp{Base: a.Base}).Request(params)
	case payloads.QpayApiPayBar:
		return (&PayBar{Base: a.Base}).Request(params)
	case payloads.QpayApiPayPub:
		return (&PayPub{Base: a.Base}).Request(params)
	case payloads.QpayApiPayQr:
		return (&PayApp{Base: a.Base}).Request(params)
	case payloads.QpayApiQueryTrade:
		return (&QueryTrade{Base: a.Base}).Request(params)
	// refund apis
	case payloads.QpayApiRefund:
		return (&Refund{Base: a.Base}).Request(params)
	case payloads.QpayApiQueryRefund:
		return (&QueryRefund{Base: a.Base}).Request(params)
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
	b.GatewayUrl = "https://qpay.qq.com/cgi-bin/"

	return b
}

func (b *Base) GetFullGatewayUrl(method string) string {
	return fmt.Sprintf(b.GatewayUrl, method)
}

func (b *Base) Request(uri string, params map[string]string) (*BaseResponse, error) {

	return NewBaseResponse(""), nil
}
