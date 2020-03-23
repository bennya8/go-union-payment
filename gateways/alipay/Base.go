package alipay

import (
	"fmt"
	"github.com/bennya8/go-union-payment/contracts"
	"github.com/bennya8/go-union-payment/payloads"
)

func Factory(channel payloads.UnionPaymentChannel, config contracts.IGatewayConfig) contracts.IGateway {
	cfg := config.ParseConfig().(Config)
	b := NewBase(cfg)
	fmt.Println(b)

	switch channel {
	case payloads.AliChannelApp:
	case payloads.AliChannelWap:
	case payloads.AliChannelWeb:
	case payloads.AliChannelQr:
	case payloads.AliChannelBar:
	}
	return nil
}

type Base struct {
	Config Config
}

func NewBase(config Config) *Base {
	b := &Base{}
	b.Config = config

	return b
}
