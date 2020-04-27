package alipay

import (
	"github.com/bennya8/go-union-payment/payloads"
)

type PayApp struct {
	Base *Base

}

func (p PayApp) Request(params map[string]string) *payloads.UnionPaymentResult {
	panic("implement me")
}

func (p PayApp) BuildParams(params map[string]string) map[string]string {
	panic("implement me")
}

