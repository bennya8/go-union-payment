package payloads

type (
	UnionPaymentGateway string
	UnionPaymentApi     string
)

const (
	Version = "1.0.0.b1"

	// gateways
	AlipayGateway UnionPaymentGateway = "Alipay"
	CmbGateway                        = "CMBank"
	WechatGateway                     = "Wechat"
	QpayGateway                       = "Qpay"

	// wx apis
	WxApiPayApp         UnionPaymentApi = "wx_api_pay_app"
	WxApiPayWap                         = "wx_api_pay_wap"
	WxApiPayPub                         = "wx_api_pay_pub"
	WxApiPayLite                        = "wx_api_pay_lite"
	WxApiPayQr                          = "wx_api_pay_qr"
	WxApiPayBar                         = "wx_api_pay_bar"
	WxApiCancelTrade                    = "wx_api_cancel_trade"
	WxApiCloseTrade                     = "wx_api_close_trade"
	WxApiQueryTrade                     = "wx_api_query_trader"
	WxApiRefund                         = "wx_api_refund"
	WxApiQueryRefund                    = "wx_api_query_refund"
	WxApiTransfer                       = "wx_api_transfer"
	WxApiQueryTransfer                  = "wx_api_query_transfer"
	WxApiBillDownload                   = "wx_api_bill_download"
	WxApiSettleDownload                 = "wx_api_settle_download"

	// alipay apis
	AliApiPayApp UnionPaymentApi = "ali_api_pay_app"
	AliApiPayWap                 = "ali_api_pay_wap"
	AliApiPayBar                 = "ali_api_pay_bar"
	AliApiPayWeb                 = "ali_api_pay_web"

	// qpay apis

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
