package alipay

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
	"github.com/bennya8/go-union-payment/utils/crypt"
	"github.com/bennya8/go-union-payment/utils/date"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
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
	PublicKey  *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
}

func NewBase(config *Config) *Base {
	b := &Base{}
	b.Config = config
	b.Http = http.Client{}

	b.GatewayUrl = "https://openapi.alipay.com/gateway.do"
	if b.Config.UseSandbox {
		b.GatewayUrl = "https://openapi.alipaydev.com/gateway.do"
	}

	// init rsa certs
	var err error
	b.PrivateKey, err = crypt.InitPrivateKey(b.Config.RsaPrivateKey, false)
	if err != nil {
		log.Fatal(err)
	}

	b.PublicKey, err = crypt.InitPublicKey(b.Config.AliPublicKey, false)
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func (b *Base) GetFullGatewayUrl(params string) string {
	return fmt.Sprintf("%s?%s", b.GatewayUrl, params)
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
	signData["notify_url"] = b.Config.NotifyUrl
	signData["timestamp"] = date.TimeFormat(time.Now())

	// biz params
	bizContent, _ := json.Marshal(bizContentParams)
	signData["biz_content"] = string(bizContent)

	if uri == "alipay.trade.page.pay" || uri == "alipay.trade.wap.pay" {
		signData["return_url"] = b.Config.ReturnUrl
	}

	signStr := b.BuildSign(signData)
	signData["sign"] = b.MakeSign(signStr)

	if uri == "alipay.trade.app.pay" {
		return NewBaseResponse(b.BuildUrlEncode(signData)), nil
	} else if uri == "alipay.trade.wap.pay" ||
		uri == "alipay.trade.page.pay" ||
		uri == "alipay.user.certify.open.certify" {
		return NewBaseResponse(b.GetFullGatewayUrl(b.BuildUrlEncode(signData))), nil
	}

	resp, err := b.Http.Post(b.GatewayUrl, "application/form-data", bytes.NewBufferString(b.BuildUrlEncode(signData)))
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
		if len(signData[k]) > 0 {
			sign += k + "=" + signData[k] + "&"
		}
	}
	return strings.TrimRight(sign, "&")
}

func (b *Base) BuildUrlEncode(kvPairs map[string]string) string {
	// sort the key
	var keys []string
	for k := range kvPairs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	params := ""
	for _, k := range keys {
		if len(kvPairs[k]) > 0 {
			if k == "sign" {
				continue
			}
			params += k + "=" + url.QueryEscape(kvPairs[k]) + "&"
		}
	}
	return params + "sign=" + url.QueryEscape(kvPairs["sign"])
}

func (b *Base) MakeSign(signStr string) string {
	var sign string

	if b.Config.SignType == SignTypeRsa {
		h := sha1.New()
		io.WriteString(h, signStr)
		hashed := h.Sum(nil)
		rsaSign, err := rsa.SignPKCS1v15(rand.Reader, b.PrivateKey, crypto.SHA1, hashed)
		if err != nil {
			log.Fatal(err)
		}
		sign = base64.StdEncoding.EncodeToString(rsaSign)
	} else if b.Config.SignType == SignTypeRsa2 {
		h := crypto.SHA256.New()
		io.WriteString(h, signStr)

		hashed := h.Sum(nil)
		rsaSign, err := rsa.SignPKCS1v15(rand.Reader, b.PrivateKey, crypto.SHA256, hashed)
		if err != nil {
			log.Fatal(err)
		}
		sign = base64.StdEncoding.EncodeToString(rsaSign)
	}
	return sign
}
