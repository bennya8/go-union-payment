package wx_example

import (
	"fmt"
	go_union_payment "github.com/bennya8/go-union-payment"
	"github.com/bennya8/go-union-payment/contracts"
	"github.com/bennya8/go-union-payment/gateways/wechat"
	"github.com/bennya8/go-union-payment/payloads"
	"io/ioutil"
	"strconv"
	"testing"
	"time"
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

func getTradeNo() string {
	return strconv.Itoa(int(time.Now().Unix()))
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

func TestWxPayWap(t *testing.T) {
	config := initWxConfig()

	payment := go_union_payment.NewUnionPayment(payloads.WechatGateway, config)
	rs := payment.Invoke(payloads.WxApiPayWap, map[string]string{
		"body":         "test body",
		"subject":      "test subject",
		"trade_no":     getTradeNo(),
		"time_expire":  strconv.Itoa(int(time.Now().Add(600 * time.Second).Unix())),
		"amount":       "0.01",
		"return_param": "123",
		"client_ip":    "127.0.0.1",
	})
	if !rs.State {
		panic(rs.Msg)
	}

	fmt.Println(rs.Data)
}

func TestWxPayPub(t *testing.T) {
	config := initWxConfig()

	payment := go_union_payment.NewUnionPayment(payloads.WechatGateway, config)
	rs := payment.Invoke(payloads.WxApiPayPub, map[string]string{
		"body":         "test body",
		"subject":      "test subject",
		"trade_no":     getTradeNo(),
		"time_expire":  strconv.Itoa(int(time.Now().Add(600 * time.Second).Unix())),
		"amount":       "0.01",
		"return_param": "123",
		"client_ip":    "127.0.0.1",
		"scene_info":   `{"h5_info":{"type":"Wap","wap_url":"https:\/\/Leo112g.github.io\/","wap_name":"\u6d4b\u8bd5\u5145\u503c"}}`,
	})
	if !rs.State {
		panic(rs.Msg)
	}

	fmt.Println(rs.Data)
}
