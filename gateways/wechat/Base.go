package wechat

import (
	"fmt"
	"github.com/bennya8/go-union-payment/utils/str"
	"github.com/gojektech/heimdall/httpclient"
	"io/ioutil"
	"time"
)

const NonceLen = 32
const ReqSuccess = "SUCCESS"
const SignTypeMd5 = "MD5"
const SignTypeSha = "HMAC-SHA256"

type Base struct {
	GatewayUrl string
	MerKey     string
	SandboxKey string
	IsSandbox  bool
	ReturnRaw  bool
	NonceStr   string
	UseBackup  bool
	SignType   string
}

func NewBase(config Config) *Base {
	base := &Base{}
	base.IsSandbox = config.UseSandbox
	base.UseBackup = config.UseBackup
	base.ReturnRaw = config.ReturnRaw
	base.MerKey = config.Md5Key
	base.SignType = config.SignType
	base.NonceStr = str.GetNonce(NonceLen)

	// initial wechat gateway address
	base.GatewayUrl = "https://api.mch.weixin.qq.com/%s"
	if base.IsSandbox {
		base.GatewayUrl = "https://api.mch.weixin.qq.com/sandboxnew/%s"
	} else if base.UseBackup {
		base.GatewayUrl = "https://api2.mch.weixin.qq.com/%s"
	}

	// if using sandbox env then switch the merchant key
	if base.IsSandbox && len(base.SandboxKey) == 0 {
		base.SandboxKey = base.getSignKey()
		base.MerKey = base.SandboxKey
	}

	return base
}

func (*Base) RequestWxApi()  {
	timeout := 1000 * time.Millisecond
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	// Use the clients GET method to create and execute the request
	res, err := client.Get("http://google.com", nil)
	if err != nil{
		panic(err)
	}

	// Heimdall returns the standard *http.Response object
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func (*Base) getSignKey() string {
	//method := "pay/getsignkey"
	// @todo request wx api
	return ""
}

func (*Base) buildParams() {

}

func (*Base) changeKeyName() {

}


