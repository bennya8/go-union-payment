package cmb_example

import (
	"fmt"
	"github.com/bennya8/go-union-payment/gateways/cmb"
	"io/ioutil"
	"testing"
)

func init() {

}

func TestCmbConfigWithYaml(t *testing.T) {
	config, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		t.Error(err)
	}
	wxConfig, err := cmb.NewConfigWithYaml(config)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(wxConfig)
}

func TestCmbConfigWithJson(t *testing.T) {
	config, err := ioutil.ReadFile("./config.json")
	if err != nil {
		t.Error(err)
	}
	wxConfig, err := cmb.NewConfigWithJson(config)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(wxConfig)
}