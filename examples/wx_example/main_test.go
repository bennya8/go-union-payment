package wx_example

import (
	go_union_payment "github.com/bennya8/go-union-payment"
	"github.com/bennya8/go-union-payment/payloads"
	"testing"
)

func TestWxApiWap(t *testing.T) {


	go_union_payment.NewUnionPayment(payloads.WechatGateway, )
}
