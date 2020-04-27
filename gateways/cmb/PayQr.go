package cmb

import (
	"github.com/bennya8/go-union-payment/payloads"
)

type PayQr struct {
	Base *Base
}

func (p PayQr) Request(params map[string]string) *payloads.UnionPaymentResult {
	panic("implement me")
}

func (p PayQr) BuildParams(params map[string]string) map[string]string {
	panic("implement me")
}
