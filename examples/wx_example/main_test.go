package wx_example

import (
	"fmt"
	go_union_payment "github.com/bennya8/go-union-payment"
	"github.com/bennya8/go-union-payment/gateways/wechat"
	"io/ioutil"
	"testing"
)

var wxUnionPayment *go_union_payment.UnionPayment

func init() {
	//wxUnionPayment = go_union_payment.NewUnionPayment(payloads.WechatGateway, )
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

func TestWxApiWap(t *testing.T) {

	a := fmt.Sprintf("%.2f", 11.1812312)
	fmt.Println(a)
}
