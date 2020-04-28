package examples

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

type NotifyService struct {
}

func (s NotifyService) PayNotify(notify payloads.UnionPaymentNotify, notifyData interface{}) (isSuccess bool) {
	fmt.Println(notify)
	fmt.Println(notifyData)

	return true
}
