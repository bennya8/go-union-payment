package wechat

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	"github.com/bennya8/go-union-payment/certs"
	"github.com/bennya8/go-union-payment/contracts"
	"github.com/bennya8/go-union-payment/payloads"
	"github.com/bennya8/go-union-payment/utils/crypt"
	"github.com/bennya8/go-union-payment/utils/str"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const NonceLen = 32
const SignTypeMd5 = "MD5"
const SignTypeSha = "HMAC-SHA256"

type Gateway struct {
	Base *Base
}

func (w *Gateway) Request(api payloads.UnionPaymentApi, params map[string]string) *payloads.UnionPaymentResult {
	switch api {
	case payloads.WxApiPayApp:
		return (&PayApp{Base: w.Base}).Request(params)
	case payloads.WxApiPayWap:
		return (&PayWap{Base: w.Base}).Request(params)
	case payloads.WxApiPayPub:
		return (&PayPub{Base: w.Base}).Request(params)
	case payloads.WxApiPayLite:
		return (&PayLite{Base: w.Base}).Request(params)
	case payloads.WxApiPayQr:
		return (&PayQr{Base: w.Base}).Request(params)
	case payloads.WxApiPayBar:
		return (&PayBar{Base: w.Base}).Request(params)
	case payloads.WxApiCancelTrade:
	case payloads.WxApiCloseTrade:
	case payloads.WxApiQueryTrade:
	case payloads.WxApiQueryRefund:
	case payloads.WxApiTransfer:
	case payloads.WxApiQueryTransfer:
	case payloads.WxApiBillDownload:
	case payloads.WxApiSettleDownload:
	}
	return payloads.NewUnionPaymentResult(false, "unknown gateway api", nil)
}

type Base struct {
	GatewayUrl string
	MerKey     string
	SandboxKey string
	IsSandbox  bool
	ReturnRaw  bool
	NonceStr   string
	UseBackup  bool
	SignType   string
	Http       http.Client
	Config     Config
}

func NewBase(config Config) *Base {
	b := &Base{}
	b.Config = config

	// @todo the following params could be remove lately.
	b.IsSandbox = config.UseSandbox
	b.UseBackup = config.UseBackup
	b.ReturnRaw = config.ReturnRaw
	b.MerKey = config.Md5Key
	b.SignType = config.SignType
	b.NonceStr = str.GetNonce(NonceLen)

	// initial wechat gateway address
	b.GatewayUrl = "https://api.mch.weixin.qq.com/%s"
	if b.IsSandbox {
		b.GatewayUrl = "https://api.mch.weixin.qq.com/sandboxnew/%s"
	} else if b.UseBackup {
		b.GatewayUrl = "https://api2.mch.weixin.qq.com/%s"
	}

	// if using sandbox env then switch the merchant key
	if b.IsSandbox && len(b.SandboxKey) == 0 {
		b.SandboxKey = b.getSignKey()
		b.MerKey = b.SandboxKey
	}

	b.initHttp()

	return b
}

func (b *Base) GetFullGatewayUrl(method string) string {
	return fmt.Sprintf(b.GatewayUrl, method)
}

func (b *Base) Request(uri string, params map[string]string) (*BaseResponse, error) {
	// setting up base request params.
	baseParams := map[string]string{
		"app_id":     b.Config.AppId,
		"sub_appid":  b.Config.SubAppId,
		"mch_id":     b.Config.MchId,
		"sub_mch_id": b.Config.SubMchId,
		"nonce_str":  b.NonceStr,
		"sign_type":  b.SignType,
	}
	for k, v := range baseParams {
		params[k] = v
	}

	// strips all empty values
	for k, v := range params {
		if len(v) <= 0 {
			delete(params, k)
		}
	}

	// sort params by keys
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// build signature
	var signStr string
	for _, k := range keys {
		v, _ := url.QueryUnescape(params[k])
		signStr += k + "=" + v + "&"
	}
	if strings.LastIndex(signStr, "&") != -1 {
		signStr = signStr[:len(signStr)-1]
	}

	params["sign"] = b.makeSign(signStr)

	x, _ := xml.MarshalIndent(contracts.XmlParams(params), "", "  ")
	contentBody := bytes.NewBuffer(x)
	resp, err := b.Http.Post(uri, "application/xml", contentBody)
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

func (b *Base) initHttp() {
	// load the ca pem via. binary
	reader := strings.NewReader(certs.WxCaCertPem())
	caCert, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	clientCert, err := tls.LoadX509KeyPair(b.Config.AppCertPem, b.Config.AppKeyPem)
	tlsConfig := tls.Config{
		RootCAs:      pool,
		Certificates: []tls.Certificate{clientCert},
	}
	transport := http.Transport{
		TLSClientConfig: &tlsConfig,
	}
	b.Http = http.Client{
		Transport: &transport,
	}
}

func (*Base) getSignKey() string {
	//method := "pay/getsignkey"
	// @todo request wx api is require when fetch the signature key with sanbox env.
	return ""
}

func (b *Base) makeSign(signStr string) string {
	var sign string
	if b.SignType == SignTypeMd5 {
		signStr += "&key=" + b.MerKey
		sign = crypt.MD5(signStr)
	} else if b.SignType == SignTypeSha {
		signStr += "&key=" + b.MerKey
		sign = crypt.HmacSha1(signStr, b.MerKey)
	}
	return sign
}

//func (*Base) modifiedParams(params map[string]interface{}) {
//	modKeys := map[string]bool{
//		"mmpaymkttransfers/promotion/transfers": true,
//		"mmpaymkttransfers/sendredpack":         true,
//	}
//	if _, ok := params[modKeys] {
//
//	}
//
//}
