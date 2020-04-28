package wx_example

import (
	"bytes"
	"fmt"
	go_union_payment "github.com/bennya8/go-union-payment"
	"github.com/bennya8/go-union-payment/contracts"
	"github.com/bennya8/go-union-payment/gateways/wechat"
	"github.com/bennya8/go-union-payment/payloads"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"
	"time"
)

// 自定义的实现类，用于接管notify通知
type WxPaymentService struct {
}

func (a *WxPaymentService) PayNotify(notify payloads.UnionPaymentNotify, notifyData interface{}) (isSuccess bool) {

	if notify == payloads.WxNotifyPay { // 微信支付成功的回调
		// 强转类型
		wxNotifyPayRsp := notifyData.(wechat.WxNotifyPayResponse)

		// 可自行实现业务逻辑代码。更改订单状态之类
		fmt.Println(wxNotifyPayRsp)
	} else if notify == payloads.WxNotifyRefund {
		// 强转类型
		wxNotifyPayRsp := notifyData.(wechat.WxNotifyRefundResponse)

		// 可自行实现业务逻辑代码。更改订单状态之类
		fmt.Println(wxNotifyPayRsp)
	}

	fmt.Println(notify)
	fmt.Println(notifyData)
	return false
}

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

func TestWxNotifyRefund(t *testing.T) {
	service := &WxPaymentService{}

	var req = &http.Request{}
	req.Body = ioutil.NopCloser(bytes.NewBufferString(`<xml>
        <return_code>SUCCESS</return_code>
           <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
           <mch_id><![CDATA[10000100]]></mch_id>
           <nonce_str><![CDATA[TeqClE3i0mvn3DrK]]></nonce_str>
           <req_info><![CDATA[T87GAHG17TGAHG1TGHAHAHA1Y1CIOA9UGJH1GAHV871HAGAGQYQQPOOJMXNBCXBVNMNMAJAA]]></req_info>
        </xml>`))

	config := initWxConfig()
	payment := go_union_payment.NewUnionPayment(payloads.WechatGateway, config)
	err := payment.ParserNotify(req, service)
	if err != nil {
		t.Error(err)
	}
}

func TestWxNotifyPay(t *testing.T) {
	service := &WxPaymentService{}
	var req = &http.Request{}
	req.Body = ioutil.NopCloser(bytes.NewBufferString(`<xml>
  <appid><![CDATA[wx2421b1c4370ec43b]]></appid>
  <attach><![CDATA[支付测试]]></attach>
  <bank_type><![CDATA[CFT]]></bank_type>
  <fee_type><![CDATA[CNY]]></fee_type>
  <is_subscribe><![CDATA[Y]]></is_subscribe>
  <mch_id><![CDATA[10000100]]></mch_id>
  <nonce_str><![CDATA[5d2b6c2a8db53831f7eda20af46e531c]]></nonce_str>
  <openid><![CDATA[oUpF8uMEb4qRXf22hE3X68TekukE]]></openid>
  <out_trade_no><![CDATA[1409811653]]></out_trade_no>
  <result_code><![CDATA[SUCCESS]]></result_code>
  <return_code><![CDATA[SUCCESS]]></return_code>
  <sign><![CDATA[B552ED6B279343CB493C5DD0D78AB241]]></sign>
  <time_end><![CDATA[20140903131540]]></time_end>
  <total_fee>1</total_fee>
<coupon_fee><![CDATA[10]]></coupon_fee>
<coupon_count><![CDATA[1]]></coupon_count>
<coupon_type><![CDATA[CASH]]></coupon_type>
<coupon_id><![CDATA[10000]]></coupon_id>
<coupon_fee><![CDATA[100]]></coupon_fee>
  <trade_type><![CDATA[JSAPI]]></trade_type>
  <transaction_id><![CDATA[1004400740201409030005092168]]></transaction_id>
</xml>`))
	config := initWxConfig()
	payment := go_union_payment.NewUnionPayment(payloads.WechatGateway, config)
	err := payment.ParserNotify(req, service)

	if err != nil {
		t.Error(err)
	}
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
		"scene_info": "",
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
		"scene_info": `{"type":"Wap","wap_url":"http://yourhost","wap_name":"your host name"}`,
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
		"scene_info": "",
	})
	if !rs.State {
		panic(rs.Msg)
	}

	fmt.Println(rs.Data.ToJson())
}
