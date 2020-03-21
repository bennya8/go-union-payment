package go_union_payment

import (
	"github.com/bennya8/go-union-payment/contracts"
	"github.com/bennya8/go-union-payment/gateways/wechat"
	"github.com/bennya8/go-union-payment/payloads"
	"sync"
)

var (
	instance *UnionPayment
	once     sync.Once
)

func NewUnionPayment() *UnionPayment {
	once.Do(func() {
		instance = &UnionPayment{}
	})
	return instance
}

type UnionPayment struct {
}

func (u *UnionPayment) Pay(channel payloads.UnionPaymentChannel, config contracts.IGatewayConfig) *payloads.UnionPaymentResult {
	gateway := u.gatewayFactory(channel, config)

	return gateway.Request()
}

func (u *UnionPayment) ParserNotify(notify contracts.IPaymentNotify) {
	//notify.PayNotify()
}

func (u *UnionPayment) gatewayFactory(channel payloads.UnionPaymentChannel, config contracts.IGatewayConfig) contracts.IGateway {
	if channel == payloads.WxChannelApp ||
		channel == payloads.WxChannelWap ||
		channel == payloads.WxChannelBar ||
		channel == payloads.WxChannelPub ||
		channel == payloads.WxChannelQr {

		return wechat.Factory(channel, config)
	}

	panic("unknown gateway")
}
