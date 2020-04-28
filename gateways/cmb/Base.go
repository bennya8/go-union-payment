package cmb

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

type Gateway struct {
	Base *Base
}

func (a *Gateway) Request(api payloads.UnionPaymentApi, params map[string]string) *payloads.UnionPaymentResult {
	switch api {
	case payloads.CmbApiPayApp:
	case payloads.CmbApiPayLite:
	case payloads.CmbApiPayQr:
	case payloads.CmbApiPayWap:
	case payloads.AliApiPayWeb:
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
	b.GatewayUrl = "https://payment.ebank.cmbchina.com/NetPayment/"
	if b.Config.UseSandbox {
		b.GatewayUrl = "http://121.15.180.66:801/NetPayment_dl/"
	}

	return b
}

func (b *Base) GetFullGatewayUrl(method string) string {
	return fmt.Sprintf(b.GatewayUrl, method)
}

func (b *Base) Request(uri string, params map[string]string) (*BaseResponse, error) {

	return NewBaseResponse(""), nil
}
