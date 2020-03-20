package go_union_payment

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/bennya8/go-union-payment/gateways/wechat"
	"github.com/bennya8/go-union-payment/payloads"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestWechatPayment(t *testing.T) {
	var rawJson = `{
"use_sandbox":false,
"app_id":"<公众账号ID>",
"sub_appid":"<公众子商户账号ID>"
}`

	config, err := wechat.NewConfigWithJson([]byte(rawJson))
	fmt.Println(err)
	fmt.Println(config)

	NewUnionPayment().Pay(payloads.WxChannelApp, config)
}

func TestCaClient(t *testing.T) {
	caCert, err := ioutil.ReadFile("/Users/Benny/Workspaces/php-projects/benny/go-union-payment/certs/wx_cacert.pem")
	if err != nil {
		log.Fatal(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	cert := "/Users/Benny/Workspaces/php-projects/diggmind/saas_cloud/common/certs/diggme_wx/apiclient_cert.pem"
	key := "/Users/Benny/Workspaces/php-projects/diggmind/saas_cloud/common/certs/diggme_wx/apiclient_key.pem"
	// LoadX509KeyPair reads files, so we give it the paths
	clientCert, err := tls.LoadX509KeyPair(cert, key)
	tlsConfig := tls.Config{
		RootCAs:      pool,
		Certificates: []tls.Certificate{clientCert},
	}
	transport := http.Transport{
		TLSClientConfig: &tlsConfig,
	}
	client := http.Client{
		Transport: &transport,
	}

	resp, err := client.Get("https://api.mch.weixin.qq.com/pay/downloadbill?appid=xxxxxx")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}

}

func TestAlipayPayment(t *testing.T) {

}

func TestCmbPayment(t *testing.T) {

}

func TestPaypalPayment(t *testing.T) {

}
