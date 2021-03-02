package alipay

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

const SignTypeRsa = "RSA"
const SignTypeRsa2 = "RSA2"

type Config struct {
	UseSandbox    bool     `json:"use_sandbox" yaml:"use_sandbox"`
	AppId         string   `json:"app_id" yaml:"app_id"`
	Charset       string   `json:"charset" yaml:"charset"`
	Format        string   `json:"format" yaml:"format"`
	Version       string   `json:"version" yaml:"version"`
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

func (c Config) CheckConfig() error {

	if len(c.Charset) <= 0 {
		c.Charset = "UTF-8"
	}
	if len(c.Format) <= 0 {
		c.Format = "JSON"
	}
	if len(c.Version) <= 0 {
		c.Format = "1.0"
	}

	return nil
}

func NewConfigWithJson(content []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	err = config.CheckConfig()
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

	err = config.CheckConfig()
	if err != nil {
		return nil, err
	}

	return &config, nil
}
