package go_union_payment

import (
	"github.com/bennya8/go-union-payment/contracts"
	"github.com/bennya8/go-union-payment/gateways/alipay"
	"github.com/bennya8/go-union-payment/gateways/cmb"
	"github.com/bennya8/go-union-payment/gateways/qpay"
	"github.com/bennya8/go-union-payment/gateways/wechat"
	"github.com/bennya8/go-union-payment/payloads"
	"net/http"
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

func (u *UnionPayment) Invoke(channel payloads.UnionPaymentChannel, config contracts.IGatewayConfig) *payloads.UnionPaymentResult {
	gateway := u.gatewayFactory(channel, config)

	return gateway.Request()
}

func (u *UnionPayment) ParserNotify(req *http.Response, notify contracts.IPaymentNotify) {
	//notify.PayNotify()
}

func (u *UnionPayment) gatewayFactory(channel payloads.UnionPaymentChannel, config contracts.IGatewayConfig) contracts.IGateway {
	if channel == payloads.WxChannelApp ||
		channel == payloads.WxChannelWap ||
		channel == payloads.WxChannelBar ||
		channel == payloads.WxChannelPub ||
		channel == payloads.WxChannelQr {

		return wechat.Factory(channel, config)
	} else if channel == payloads.AliChannelApp ||
		channel == payloads.AliChannelWap ||
		channel == payloads.AliChannelWeb ||
		channel == payloads.AliChannelQr ||
		channel == payloads.AliChannelBar {

		return alipay.Factory(channel, config)
	} else if channel == payloads.QpayChannelApp ||
		channel == payloads.QpayChannelWap ||
		channel == payloads.QpayChannelQr {

		return qpay.Factory(channel, config)
	} else if channel == payloads.CmbChannelApp ||
		channel == payloads.CmbChannelWap ||
		channel == payloads.CmbChannelWeb ||
		channel == payloads.CmbChannelQr ||
		channel == payloads.CmbChannelLite {

		return cmb.Factory(channel, config)
	}

	panic("unknown gateway")
}
