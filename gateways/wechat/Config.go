package wechat

import (
	"encoding/json"
)

type Config struct {
	UseSandbox bool   `json:"use_sandbox"`
	UseBackup  bool   `json:"use_backup"`
	AppId      string `json:"app_id"`
	SubAppId   string `json:"sub_app_id"`
	SubMchId   string `json:"sub_mch_id"`
	Md5Key     string `json:"md_5_key"`
	//CaPem       string   `json:"ca_pem"` // @todo using the build-in wx-ca.pem temporary, see certs.go -> WxCaCertPem() method
	AppCertPem  string   `json:"app_cert_pem"`
	AppKeyPem   string   `json:"app_key_pem"`
	SignType    string   `json:"sign_type"`
	LimitPay    []string `json:"limit_pay"`
	FeeType     string   `json:"fee_type"`
	ReturnRaw   bool     `json:"return_raw"`
	NotifyUrl   string   `json:"notify_url"`
	RedirectUrl string   `json:"redirect_url"`

}

func (c Config) ParseConfig() interface{} {
	return c
}

func NewConfigWithJson(rawJson []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(rawJson, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
