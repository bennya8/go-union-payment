package wechat

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/bennya8/go-union-payment/certs"
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
	case payloads.WxApiFundDownload:
	}
	return payloads.NewUnionPaymentResult(false, "unknown gateway api", nil)
}

type Base struct {
	Config *Config
	Http   http.Client
	// common properties
	GatewayUrl string
}

func NewBase(config *Config) *Base {
	b := &Base{}
	b.Config = config

	// initial wechat gateway address
	b.GatewayUrl = "https://api.mch.weixin.qq.com/%s"
	if b.Config.UseSandbox {
		b.GatewayUrl = "https://api.mch.weixin.qq.com/sandboxnew/%s"
	} else if b.Config.UseBackup {
		b.GatewayUrl = "https://api2.mch.weixin.qq.com/%s"
	}

	// if using sandbox env then switch the merchant key
	if b.Config.UseBackup && len(b.Config.SandboxKey) == 0 {
		b.Config.SandboxKey = b.getSignKey()
		b.Config.Md5Key = b.Config.SandboxKey
	}

	// init http client with CA certs
	b.initHttp()

	return b
}

func (b *Base) GetFullGatewayUrl(method string) string {
	return fmt.Sprintf(b.GatewayUrl, method)
}

func (b *Base) Request(uri string, params map[string]string) (*BaseResponse, error) {
	// setting up base request params.
	baseParams := map[string]string{
		"appid":      b.Config.AppId,
		"sub_appid":  b.Config.SubAppId,
		"mch_id":     b.Config.MchId,
		"sub_mch_id": b.Config.SubMchId,
		"nonce_str":  str.GetNonce(NonceLen),
		"sign_type":  b.Config.SignType,
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

	// build xml
	x := "<xml>"
	for _, k := range keys {
		v, _ := url.QueryUnescape(params[k])
		signStr += k + "=" + v + "&"
		x += "<" + k + "><![CDATA[" + v + "]]></" + k + ">"
	}
	// strip last &
	if strings.LastIndex(signStr, "&") != -1 {
		signStr = signStr[:len(signStr)-1]
	}

	// append signature
	sign := b.MakeSign(signStr)
	x += "<sign><![CDATA[" + sign + "]]></sign>"
	x += "</xml>"

	contentBody := bytes.NewBufferString(x)
	resp, err := b.Http.Post(uri, "application/xml", contentBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if strings.Contains(string(body), "<return_code><![CDATA[FAIL]]></return_code>") {
		return nil, errors.New(string(body))
	}

	return NewBaseResponse(string(body)), nil
}

func (*Base) getSignKey() string {
	//method := "pay/getsignkey"
	// @todo request wx api is necessary when fetch the signature key with sanbox env.
	return ""
}

func (b *Base) BuildSign(timestamp string, retMap map[string]interface{}) string {
	var (
		buffer strings.Builder
	)
	buffer.WriteString("appId=")
	buffer.WriteString(retMap["appid"].(string))
	buffer.WriteString("&nonceStr=")
	buffer.WriteString(retMap["nonce_str"].(string))
	buffer.WriteString("&package=")
	buffer.WriteString("prepay_id=" + retMap["prepay_id"].(string))
	buffer.WriteString("&signType=")
	buffer.WriteString(b.Config.SignType)
	buffer.WriteString("&timeStamp=")
	buffer.WriteString(timestamp)
	return buffer.String()
}

func (b *Base) MakeSign(signStr string) string {
	var sign string
	if b.Config.SignType == SignTypeMd5 {
		signStr += "&key=" + b.Config.Md5Key
		sign = crypt.MD5(signStr)
	} else if b.Config.SignType == SignTypeSha {
		signStr += "&key=" + b.Config.Md5Key
		sign = crypt.HmacSha1(signStr, b.Config.Md5Key)
	}
	return sign
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
