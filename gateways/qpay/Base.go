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
	case payloads.QpayApiPayApp:
	case payloads.QpayApiPayBar:
	case payloads.QpayApiPayPub:
	case payloads.QpayApiPayQr:
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
