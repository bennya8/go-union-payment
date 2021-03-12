# go-union-payment

English | [中文](/README.md)

---


Payment is a Go version of the payment aggregation third-party SDK, which integrates WeChat, Alipay, Qpay, China Merchant Bank payment and maybe more. 

Inspired by [https://github.com/helei112g/payment] which is a php-version of unified payment lib.

**PR** is always Welcome!

#### payment support list

- Wechat [Working]
- Alipay [TODO]
- Qpay [TODO]
- China Merchant Bank [TODO]
- Palpal [TODO]

#### Wechat API

| API | Finished Status | Documentation |
| --- | --- | --- |
| WxApiPayApp | ✅ | https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1 |
| WxApiPayWap | ✅ | https://pay.weixin.qq.com/wiki/doc/api/H5.php?chapter=9_20&index=1 | 
| WxApiPayPub | ✅ | https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1 |
| WxApiPayLite | ✅ | https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1 |
| WxApiPayQr | ✅ | https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1 |
| WxApiPayBar | ✅ | https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=9_1 |
| WxApiCloseTrade | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3 |
| WxApiQueryTrade | 🕓 |  https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2 | 
| WxApiCancelTrade | 🕓 |https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3 |
| WxApiRefund | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4 |
| WxApiQueryRefund | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5 |
| WxApiBillDownload | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6 |
| WxApiFundDownload | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7 |
| WxNotifyPay | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_7&index=8 | 
| WxNotifyRefund | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=10 |
| WxApiTransfer | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2 |
| WxApiQueryTransfer | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3 | 
| WxApiTransferBank | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_2 | 
| WxApiQueryTransferBank | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_3 | 
| WxApiProfitSharing | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_1&index=1 | 
| WxApiMultiProfitSharing | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_6&index=2 |  
| WxApiProfitSharingQuery | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_2&index=3 | 
| WxApiProfitSharingAddReceiver | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_3&index=4 |   
| WxApiProfitSharingRemoveReceiver | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_4&index=5 |
| WxApiProfitSharingFinish | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_5&index=6 | 
| WxApiProfitSharingReturn | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_7&index=7 | 
| WxApiProfitSharingReturnQuery | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_8&index=8 | 
| WxNotifyProfitSharing | 🕓 | https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_9&index=9 | 

#### Alipay API

| API | Finished Status | Documentation |
| --- | --- | --- |
| AliApiPayApp | ✅ | https://opendocs.alipay.com/apis/api_1/alipay.trade.app.pay |
| AliApiPayBar | ✅ | https://opendocs.alipay.com/apis/api_1/alipay.trade.pay |
| AliApiPayQr | ✅ | https://opendocs.alipay.com/apis/api_1/alipay.trade.precreate |
| AliApiPayWap | ✅ | https://opendocs.alipay.com/apis/api_1/alipay.trade.wap.pay |
| AliApiPayWeb | ✅ | https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay?scene=API002020081300013629 |

#### Qpay API

| API | Finished Status | Documentation |
| --- | --- | --- |
| QpayApiPayApp | 🕓 | |
| QpayApiPayPub | 🕓 | |
| QpayApiPayQr | 🕓 | |
| QpayApiPayBar | 🕓 | |

#### China Merchant Bank API

@todo

#### PalPal API

@todo

---

### Installation

> go get github.com/bennya8/go-union-payment

### Usage

#### First step 

-- Init config with one of following ways.

**a) json**

```golang
configJson, err: = ioutil.ReadFile("./config.json")
if err != nil {
    t.Error(err)
}
wechatConfig, err: = wechat.NewConfigWithJson(configJson)
if err != nil {
    t.Error(err)
}
```

**b) yaml**

```golang
configYaml, err: = ioutil.ReadFile("./config.yaml")
if err != nil {
    t.Error(err)
}
wechatConfig, err: = wechat.NewConfigWithYaml(configYaml)
if err != nil {
    t.Error(err)
}
```

**c) struct**

```golang
wechatConfig: = wechat.Config {
    UseSandbox: false,
    UseBackup: false,
    AppId: "",
    SubAppId: "",
    SubMchId: "",
    MchId: "",
    Md5Key: "",
    AppCertPem: "",
    AppKeyPem: "",
    SignType: "",
    LimitPay: nil,
    FeeType: "",
    ReturnRaw: false,
    NotifyUrl: "",
    RedirectUrl: "",
}
```


#### Second step 

-- Call gateway api with params to get the unified response, and there you go!


```golang
// call the gateway channel api
payData: = map[string] string {
    "body": "goods information",
    "subject": "debug goods app payment",
    "trade_no": "xxx" + strconv.Itoa(rand.Intn(99999999)),
    "time_expire": strconv.Itoa(int(time.Now().Unix()) + 600),
    "amount": "0.01",
    "return_param": "anything_you_want",
    "goods_type": "1", // 0 虚拟，1 实物,
    "store_id": "",
}

result: = payment.Invoke(payloads.AliApiPayApp, payData)

if !result.State {
    t.Error(result.Msg)
}
m, e: = result.Data.ToMap()
fmt.Println(m, e)

j, e: = result.Data.ToJson()
fmt.Println(j, e)

x, e: = result.Data.ToXml()
fmt.Println(x, e)
```

### Step three

-- Setting up gateway notify callback

#### MUST implement IPaymentNotify interface for handler the callback event. 

```golang
type AliPaymentService struct {}
func (a *AliPaymentService) PayNotify(gateway payloads.UnionPaymentGateway, notifyData string) {
	if gateway == payloads.AlipayGateway {
		// @todo parse notifyData
		// and maybe change the order status
	} 
}
```


**Iris Framework - Route Context**

```golang
package notify

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type PaymentController struct {
}

func (c *PaymentController) BeforeActivation(b mvc.BeforeActivation) {
	b.HandleMany("POST", "/payment/callback", "Callback")
}

func (c *PaymentController) AfterActivation(a mvc.AfterActivation) {
}

func (c *PaymentController) Callback(ctx iris.Context) mvc.Result {
	req := ctx.Request()
    service := &AliPaymentService{}
    payment := go_union_payment.NewUnionPayment(payloads.AlipayGateway, aliConfig)
    payment.ParserNotify(req,service)
}
```

**Build-in net.http Server - http.Request**

```golang
http.HandleFunc("/your_notify_url", func(w http.ResponseWriter, r *http.Request) {        
    service := &AliPaymentService{}
    payment := go_union_payment.NewUnionPayment(payloads.AlipayGateway, aliConfig)
    payment.ParserNotify(r, service)
})
```



### Examples

You can find all the unit test in the examples folder. 

> /examples/

#### 3rd dependencies

- gopkg.in/yaml.v2

# LICENSE

MIT