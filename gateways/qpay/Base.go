package qpay

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
	case payloads.QpayChannelApp:
	case payloads.QpayChannelWap:
	case payloads.QpayChannelQr:
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
