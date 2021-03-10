package contracts

import "github.com/bennya8/go-union-payment/payloads"

type IGatewayApi interface {
	Request(params map[string]string) *payloads.UnionPaymentResult
	BuildParams(params map[string]string) map[string]string
}
