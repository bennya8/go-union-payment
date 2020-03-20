package contracts

import "github.com/bennya8/go-union-payment/payloads"

type IPaymentNotify interface {
	PayNotify(gateway payloads.UnionPaymentGateway, channel payloads.UnionPaymentChannel, notifyData string)
}
