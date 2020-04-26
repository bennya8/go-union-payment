package qpay_example

import (
	"fmt"
	"github.com/bennya8/go-union-payment/gateways/qpay"
	"io/ioutil"
	"testing"
)

func init() {

}

func TestQpayConfigWithYaml(t *testing.T) {
	config, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		t.Error(err)
	}
	wxConfig, err := qpay.NewConfigWithYaml(config)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(wxConfig)
}

func TestQpayConfigWithJson(t *testing.T) {
	config, err := ioutil.ReadFile("./config.json")
	if err != nil {
		t.Error(err)
	}
	wxConfig, err := qpay.NewConfigWithJson(config)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(wxConfig)
}
