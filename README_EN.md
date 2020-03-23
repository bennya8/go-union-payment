# go-union-payment

English | [中文](/README.md)

---


Payment is a Go version of the payment aggregation third-party SDK, which integrates WeChat, Alipay, Qpay, China Merchant Bank payment and maybe more. 

Inspired by [https://github.com/helei112g/payment] which is a php-version of unified payment lib.

**PR** is always Welcome!

**payment support list**

- Wechat [Working]
- Alipay [TODO]
- Qpay [TODO]
- China Merchant Bank [TODO]
- Palpal [TODO]

---

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

**b) yaml **

```golang
configJson, err: = ioutil.ReadFile("./config.json")
if err != nil {
    t.Error(err)
}
wechatConfig, err: = wechat.NewConfigWithYaml(configJson)
if err != nil {
    t.Error(err)
}
```

**c) struct **

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

**Build-in net.http Server - http.Request **

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

No extental libs requirement yet.
