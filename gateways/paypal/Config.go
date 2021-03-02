package palpal

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type Config struct {
	UseSandbox      bool   `json:"use_sandbox" yaml:"use_sandbox"`
	ReturnRaw       bool   `json:"return_raw" yaml:"return_raw"`
	ClientId        string `json:"client_id" yaml:"client_id"`
	Secret          string `json:"secret" yaml:"secret"`
	NotifyUrl       string `json:"notify_url" yaml:"notify_url"`
	RefundNotifyUrl string `json:"refund_notify_url" yaml:"refund_notify_url"`
	RedirectUrl     string `json:"redirect_url" yaml:"redirect_url"`
}

func (c Config) ParseConfig() interface{} {
	return c
}

func (c Config) CheckConfig() error {
	return nil
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
