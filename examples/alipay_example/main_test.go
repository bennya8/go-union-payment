package alipay_example

import (
	"fmt"
	go_union_payment "github.com/bennya8/go-union-payment"
	"github.com/bennya8/go-union-payment/gateways/wechat"
	"github.com/bennya8/go-union-payment/payloads"
	"io/ioutil"
	"testing"
)

func TestWapCharge(t *testing.T) {
	configJson, err := ioutil.ReadFile("./config.json")
	if err != nil {
		t.Error(err)
	}
	wxConfig, err := wechat.NewConfigWithJson(configJson)
	if err != nil {
		t.Error(err)
	}

	// instance of go-union-payment
	payment := go_union_payment.NewUnionPayment()

	// call the gateway channel api
	result := payment.Invoke(payloads.AliChannelWeb, wxConfig)

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
