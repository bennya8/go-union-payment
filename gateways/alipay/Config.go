package alipay

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type Config struct {
	UseSandbox    bool     `json:"use_sandbox" yaml:"use_sandbox"`
	AppId         string   `json:"app_id" yaml:"app_id"`
	SignType      string   `json:"sign_type" yaml:"sign_type"`
	AliPublicKey  string   `json:"ali_public_key" yaml:"ali_public_key"`
	RsaPrivateKey string   `json:"rsa_private_key" yaml:"rsa_private_key"`
	LimitPay      []string `json:"limit_pay" yaml:"limit_pay"`
	NotifyUrl     string   `json:"notify_url" yaml:"notify_url"`
	ReturnUrl     string   `json:"return_url" yaml:"return_url"`
	FeeType       string   `json:"fee_type" yaml:"fee_type"`
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
