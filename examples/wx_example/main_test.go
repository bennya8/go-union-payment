package wx_example

import (
	"fmt"
	go_union_payment "github.com/bennya8/go-union-payment"
	"github.com/bennya8/go-union-payment/contracts"
	"github.com/bennya8/go-union-payment/gateways/wechat"
	"github.com/bennya8/go-union-payment/payloads"
	"io/ioutil"
	"testing"
)

func initWxConfig() contracts.IGatewayConfig {
	config, err := ioutil.ReadFile("./config.local.json")
	if err != nil {
		panic(err)
	}
	wxConfig, err := wechat.NewConfigWithJson(config)
	if err != nil {
		panic(err)
	}
	return wxConfig
}

func TestWxConfigWithYaml(t *testing.T) {
	config, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		t.Error(err)
	}
	wxConfig, err := wechat.NewConfigWithYaml(config)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(wxConfig)
}

func TestWxConfigWithJson(t *testing.T) {
	config, err := ioutil.ReadFile("./config.json")
	if err != nil {
		t.Error(err)
	}
	wxConfig, err := wechat.NewConfigWithJson(config)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(wxConfig)
}

func TestWxPayApp(t *testing.T) {
	config := initWxConfig()
	payment := go_union_payment.NewUnionPayment(payloads.WechatGateway, config)
	fmt.Println(payment)
}
