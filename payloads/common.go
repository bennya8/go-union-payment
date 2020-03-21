package payloads

type (
	UnionPaymentGateway string
	UnionPaymentChannel string
)

const (
	Version = "1.0.0.b1"

	// gateways
	AlipayGateway UnionPaymentGateway = "Alipay"
	CmbGateway    UnionPaymentGateway = "CMBank"
	WechatGateway UnionPaymentGateway = "Wechat"
	QpayGateway   UnionPaymentGateway = "Qpay"

	// wx
	WxChannelWap UnionPaymentChannel = "wx_wap"
	WxChannelApp UnionPaymentChannel = "wx_app"
	WxChannelPub UnionPaymentChannel = "wx_pub"
	WxChannelQr  UnionPaymentChannel = "wx_qr"
	WxChannelBar UnionPaymentChannel = "wx_bar"

	// ali
	Ali

	// paypal

	// cmb
)

type IGatewayResponse interface {
	ToMap() (map[string]interface{}, error)
	ToJson() (string, error)
	ToXml() (string, error)
}

type UnionPaymentResult struct {
	State bool
	Msg   string
	Data  IGatewayResponse
}

func NewUnionPaymentResult(state bool, msg string, data IGatewayResponse) *UnionPaymentResult {
	return &UnionPaymentResult{
		State: state,
		Msg:   msg,
		Data:  data,
	}
}
