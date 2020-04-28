package alipay_example

import (
	"bytes"
	"fmt"
	go_union_payment "github.com/bennya8/go-union-payment"
	"github.com/bennya8/go-union-payment/contracts"
	"github.com/bennya8/go-union-payment/examples"
	"github.com/bennya8/go-union-payment/gateways/alipay"
	"github.com/bennya8/go-union-payment/payloads"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func initAliConfig() contracts.IGatewayConfig {
	configJson, err := ioutil.ReadFile("./config.local.json")
	if err != nil {
		panic(err)
	}
	config, err := alipay.NewConfigWithJson(configJson)
	if err != nil {
		panic(err)
	}
	return config
}

func TestCmbConfigWithYaml(t *testing.T) {
	config, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		t.Error(err)
	}
	aliConfig, err := alipay.NewConfigWithYaml(config)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(aliConfig)
}

func TestCmbConfigWithJson(t *testing.T) {
	config, err := ioutil.ReadFile("./config.json")
	if err != nil {
		t.Error(err)
	}
	aliConfig, err := alipay.NewConfigWithJson(config)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(aliConfig)
}

func TestAlipayPayApp(t *testing.T) {
	config := initAliConfig()

	// instance of go-union-payment
	payment := go_union_payment.NewUnionPayment(payloads.AlipayGateway, config)

	// call the gateway channel api
	payData := map[string]string{
		"body":         "goods information",
		"subject":      "debug goods app payment",
		"trade_no":     "xxx" + strconv.Itoa(rand.Intn(99999999)),
		"time_expire":  strconv.Itoa(int(time.Now().Unix()) + 600),
		"amount":       "0.01",
		"return_param": "anything_you_want",
		"goods_type":   "1", // 0 虚拟，1 实物,
		"store_id":     "",
	}

	result := payment.Invoke(payloads.AliApiPayApp, payData)

	if !result.State {
		t.Error(result.Msg)
	}
	m, e := result.Data.ToMap()
	fmt.Println(m, e)

	j, e := result.Data.ToJson()
	fmt.Println(j, e)

	x, e := result.Data.ToXml()
	fmt.Println(x, e)
}



func TestAlipayNotifyPay(t *testing.T) {
	// 初始化接管服务
	service := &examples.NotifyService{}

	// 模拟服务端接收到的数据
	var req = &http.Request{}
	req.Body = ioutil.NopCloser(bytes.NewBufferString(`notify_type=trade_status_sync&notify_id=91722adff935e8cfa58b3aabf4dead6ibe&notify_time=2017-02-16 21:46:15&sign_type=RSA2&sign=WcO+t3D8Kg71dTlKwN7r9PzUOXeaBJwp8/FOuSxcuSkXsoVYxBpsAidprySCjHCjmaglNcjoKJQLJ28/Asl93joTW39FX6i07lXhnbPknezAlwmvPdnQuI01HZsZF9V1i6ggZjBiAd5lG8bZtTxZOJ87ub2i9GuJ3Nr/NUc9VeY=&trans_currency=USD&settle_currency=USD&settle_amount=88.88&pay_currency=CNY&pay_amount=580.04&settle_trans_rate=1&trans_pay_rate=6.5261&refund_preset_paytool_list=[{"amount":"1.00","assert_type_code":"HEMA"}]&charge_amount=8.88&charge_flags=bluesea_1&settlement_id=2018101610032004620239146945&advance_amount=11.11&notify_action_type=payByAccountAction/closeTradeAction/reverseAction/finishFPAction&discount_amount=88.88&mdiscount_amount=88.88&unfreeze_amount=18.18&auth_trade_pay_mode=CREDIT_PREAUTH_PAY`))
	config := initAliConfig()

	payment := go_union_payment.NewUnionPayment(payloads.AlipayGateway, config)
	err := payment.ParserNotify(req, service)

	if err != nil {
		t.Error(err)
	}
}
