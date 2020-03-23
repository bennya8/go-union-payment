package payloads

type (
	UnionPaymentGateway string
	UnionPaymentChannel string
)

const (
	Version = "1.0.0.b1"

	// gateways
	AlipayGateway UnionPaymentGateway = "Alipay"
	CmbGateway                        = "CMBank"
	WechatGateway                     = "Wechat"
	QpayGateway                       = "Qpay"

	// channels
	WxChannelWap   UnionPaymentChannel = "wx_wap"
	WxChannelApp                       = "wx_app"
	WxChannelPub                       = "wx_pub"
	WxChannelQr                        = "wx_qr"
	WxChannelBar                       = "wx_bar"
	AliChannelApp                      = "ali_app"
	AliChannelWap                      = "ali_wap"
	AliChannelWeb                      = "ali_web"
	AliChannelQr                       = "ali_qr"
	AliChannelBar                      = "ali_bar"
	QpayChannelWap                     = "qpay_wap"
	QpayChannelApp                     = "qpay_app"
	QpayChannelQr                      = "qpay_qr"
	CmbChannelApp                      = "cmb_app"
	CmbChannelWap                      = "cmb_wap"
	CmbChannelWeb                      = "cmb_web"
	CmbChannelQr                       = "cmb_qr"
	CmbChannelLite                     = "cmb_lite"
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
