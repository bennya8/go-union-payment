package wechat

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/bennya8/go-union-payment/certs"
	"github.com/bennya8/go-union-payment/contracts"
	"github.com/bennya8/go-union-payment/payloads"
	"github.com/bennya8/go-union-payment/utils/str"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const NonceLen = 32
const ReqSuccess = "SUCCESS"
const SignTypeMd5 = "MD5"
const SignTypeSha = "HMAC-SHA256"

func Factory(channel payloads.UnionPaymentChannel, config contracts.IGatewayConfig) contracts.IGateway {
	cfg := config.ParseConfig().(Config)
	b := NewBase(cfg)

	switch channel {
	case payloads.WxChannelWap:
		return &WapCharge{Base: b}

	}
	return nil
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

	b.initialHttpClient()

	return b
}

func (b *Base) initialHttpClient() {
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
	// @todo request wx api
	return ""
}

func (*Base) buildParams() {

}

func (*Base) changeKeyName() {

}
