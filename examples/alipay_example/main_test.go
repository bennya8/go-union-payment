package alipay_example

import (
	"fmt"
	go_union_payment "github.com/bennya8/go-union-payment"
	"github.com/bennya8/go-union-payment/gateways/wechat"
	"github.com/bennya8/go-union-payment/payloads"
	"io/ioutil"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestAliApiPayApp(t *testing.T) {
	configJson, err := ioutil.ReadFile("./config.json")
	if err != nil {
		t.Error(err)
	}
	aliConfig, err := wechat.NewConfigWithJson(configJson)
	if err != nil {
		t.Error(err)
	}

	// instance of go-union-payment
	payment := go_union_payment.NewUnionPayment(payloads.AlipayGateway, aliConfig)

	// call the gateway channel api
	payData := map[string]string{
		"body":         "goods information",
		"subject":      "debug goods app payment",
		"trade_no":     "xxx" + strconv.Itoa(rand.Intn(99999999)),
		"time_expire":  strconv.Itoa(int(time.Now().Unix()) + 600),
		"amount":       "0.01",
		"return_param": "anything_you_want",
		"goods_type":   "1", // 0 虚拟，1 实物,
		"store_id":     "",
	}

	result := payment.Invoke(payloads.AliApiPayApp, payData)

	if !result.State {
		t.Error(result.Msg)
	}
	m, e := result.Data.ToMap()
	fmt.Println(m, e)

	j, e := result.Data.ToJson()
	fmt.Println(j, e)

	x, e := result.Data.ToXml()
	fmt.Println(x, e)
}

type AliPaymentService struct {

}

func (a *AliPaymentService) PayNotify(gateway payloads.UnionPaymentGateway, notifyData string) {
	if gateway == payloads.AlipayGateway {
		// @todo parse notifyData
		// and maybe change the order status
	}
}

func TestNotify(t *testing.T) {

	service := &AliPaymentService{}

	payment := go_union_payment.NewUnionPayment(payloads.AlipayGateway, nil)
	payment.ParserNotify(nil,service)


}