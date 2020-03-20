package wechat

import (
	"encoding/json"
)

type Config struct {
	UseSandbox     bool     `json:"use_sandbox"`
	UseBackup      bool     `json:"use_backup"`
	AppId          string   `json:"app_id"`
	SubAppId       string   `json:"sub_app_id"`
	SubMchId       string   `json:"sub_mch_id"`
	Md5Key         string   `json:"md_5_key"`
	AppCertPem     string   `json:"app_cert_pem"`
	AppCertPemPath string   `json:"app_cert_pem_path"`
	AppKeyPem      string   `json:"app_key_pem"`
	AppKeyPemPath  string   `json:"app_key_pem_path"`
	SignType       string   `json:"sign_type"`
	LimitPay       []string `json:"limit_pay"`
	FeeType        string   `json:"fee_type"`
	ReturnRaw      bool     `json:"return_raw"`
	NotifyUrl      string   `json:"notify_url"`
	RedirectUrl    string   `json:"redirect_url"`
}

func NewConfigWithJson(rawJson []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(rawJson, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
