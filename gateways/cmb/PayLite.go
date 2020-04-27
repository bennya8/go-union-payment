package cmb

import (
	"github.com/bennya8/go-union-payment/payloads"
)

type PayLite struct {
	Base *Base
}

func (p PayLite) Request(params map[string]string) *payloads.UnionPaymentResult {
	panic("implement me")
}

func (p PayLite) BuildParams(params map[string]string) map[string]string {
	panic("implement me")
}
