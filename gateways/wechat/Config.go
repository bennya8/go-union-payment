package wechat

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type Config struct {
	UseSandbox  bool     `json:"use_sandbox" yaml:"use_sandbox"`
	UseBackup   bool     `json:"use_backup"`
	AppId       string   `json:"app_id"`
	SubAppId    string   `json:"sub_app_id"`
	SubMchId    string   `json:"sub_mch_id"`
	MchId       string   `json:"mch_id"`
	Md5Key      string   `json:"md5_key"`
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

func NewConfigWithJson(content []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func NewConfigWithYaml(content []byte) (*Config, error) {
	var config Config
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
