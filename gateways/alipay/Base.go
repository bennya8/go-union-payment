package alipay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
	"github.com/bennya8/go-union-payment/utils/date"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
)

type Gateway struct {
	Base *Base
}

func (a *Gateway) Request(api payloads.UnionPaymentApi, params map[string]string) *payloads.UnionPaymentResult {
	switch api {
	case payloads.AliApiPayApp:
		return (&PayApp{Base: a.Base}).Request(params)
	case payloads.AliApiPayBar:
		return (&PayBar{Base: a.Base}).Request(params)
	case payloads.AliApiPayQr:
		return (&PayQr{Base: a.Base}).Request(params)
	case payloads.AliApiPayWap:
		return (&PayWap{Base: a.Base}).Request(params)
	case payloads.AliApiPayWeb:
		return (&PayWeb{Base: a.Base}).Request(params)
	}
	return payloads.NewUnionPaymentResult(false, "unknown gateway api", nil)
}

type Base struct {
	Config *Config
	Http   http.Client
	// Common Properties
	GatewayUrl string
}

func NewBase(config *Config) *Base {
	b := &Base{}
	b.Config = config
	b.Http = http.Client{}

	b.GatewayUrl = "https://openapi.alipay.com/gateway.do?%s"
	if b.Config.UseSandbox {
		b.GatewayUrl = "https://openapi.alipaydev.com/gateway.do?%s"
	}

	return b
}

func (b *Base) GetFullGatewayUrl(method string) string {
	return fmt.Sprintf(b.GatewayUrl, method)
}

func (b *Base) Request(uri string, bizContentParams map[string]string) (*BaseResponse, error) {

	var signData = map[string]string{}

	// common params
	signData["app_id"] = b.Config.AppId
	signData["method"] = uri
	signData["charset"] = b.Config.Charset
	signData["format"] = b.Config.Format
	signData["sign_type"] = b.Config.SignType
	signData["version"] = b.Config.Version
	signData["timestamp"] = date.TimeFormat(time.Now())

	// biz params
	bizContent, _ := json.Marshal(bizContentParams)
	signData["biz_content"] = string(bizContent)

	if uri == "alipay.trade.page.pay" || uri == "alipay.trade.wap.pay" {
		signData["return_url"] = b.Config.ReturnUrl
	}

	signStr := b.BuildSign(signData)
	signData["sign"] = b.MakeSign(signStr)

	contentBody, _ := json.Marshal(signData)
	contentBodyBuf := bytes.NewBuffer(contentBody)
	resp, err := b.Http.Post(uri, "application/json", contentBodyBuf)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return NewBaseResponse(string(body)), nil
}

func (b *Base) BuildSign(signData map[string]string) string {
	// remove the sign key
	if _, exist := signData["sign"]; exist {
		delete(signData, "sign")
	}

	// sort the key
	var keys []string
	for k := range signData {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// build sign
	var sign string

	for _, k := range keys {
		sign += signData[k]
	}

	return sign
}

func (b *Base) MakeSign(signStr string) string {
	var sign string
	if b.Config.SignType == SignTypeRsa {
		// @todo load private key and encrypt with rsa util
	} else if b.Config.SignType == SignTypeRsa2 {
		// @todo load private key and encrypt with rsa2 util
	}
	return sign
}
