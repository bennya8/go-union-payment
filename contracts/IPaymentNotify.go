package contracts

import "github.com/bennya8/go-union-payment/payloads"

type IPaymentNotify interface {
	PayNotify(notify payloads.UnionPaymentNotify, notifyData interface{})
}
