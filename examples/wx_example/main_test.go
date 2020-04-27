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

func TestWxPayLite(t *testing.T) {
	config := initWxConfig()

	payment := go_union_payment.NewUnionPayment(payloads.WechatGateway, config)
	rs := payment.Invoke(payloads.WxApiPayLite, map[string]string{
		"body":         "test body",
		"subject":      "test subject",
		"trade_no":     getTradeNo(),
		"time_expire":  strconv.Itoa(int(time.Now().Add(600 * time.Second).Unix())),
		"amount":       "0.01",
		"return_param": "123",
		"client_ip":    "127.0.0.1",
		"openid":       "okJMV0hE8EmcyoSBcDIuEBTIGNg8",

		//{"store_info" : {"id": "SZTX001", "name": "腾大餐厅", "area_code": "440305", "address": "科技园中一路腾讯大厦" }}
		"scene_info":   "",
	})
	if !rs.State {
		panic(rs.Msg)
	}

	fmt.Println(rs.Data.ToJson())
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

		//IOS移动应用
		//{"h5_info": {"type":"IOS","app_name": "王者荣耀","bundle_id": "com.tencent.wzryIOS"}}
		//安卓移动应用
		//{"h5_info": {"type":"Android","app_name": "王者荣耀","package_name": "com.tencent.tmgp.sgame"}}
		//WAP网站应用
		//{"h5_info": {"type":"Wap","wap_url": "https://pay.qq.com","wap_name": "腾讯充值"}}
		"scene_info":   `{"type":"Wap","wap_url":"http://yourhost","wap_name":"your host name"}`,
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
		"openid":       "okJMV0hE8EmcyoSBcDIuEBTIGNg8",

		//{"store_info" : {"id": "SZTX001", "name": "腾大餐厅", "area_code": "440305", "address": "科技园中一路腾讯大厦" }}
		"scene_info":   "",
	})
	if !rs.State {
		panic(rs.Msg)
	}

	fmt.Println(rs.Data.ToJson())
}
