package go_union_payment

import "sync"

type (
	UnionPaymentGateway = string
	UnionPaymentChannel = string
)

const (
	Version = "1.0.0.b1"

	// gateways
	AlipayGateway UnionPaymentGateway = "Alipay"
	CmbGateway    UnionPaymentGateway = "CMBank"
	WechatGateway UnionPaymentGateway = "Wechat"
	QpayGateway   UnionPaymentGateway = "Qpay"

	// wx
	WxChannelApp UnionPaymentChannel = "wx_app"
	WxChannelPub UnionPaymentChannel = "wx_pub"
	WxChannelQr  UnionPaymentChannel = "wx_qr"
	WxChannelBar UnionPaymentChannel = "wx_bar"

	// ali
	Ali

	// paypal

	// cmb
)

var (
	instance *UnionPayment
	once     sync.Once
)

type IPayNotifiable interface {
	PayNotify(gateway UnionPaymentGateway, channel UnionPaymentChannel, notifyData string)
}

func NewUnionPayment() *UnionPayment {
	once.Do(func() {
		instance = &UnionPayment{}
	})
	return instance
}

type IUnionPayment interface {
	Pay()
}
type UnionPayment struct {
}
