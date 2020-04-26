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

	// 分账API
	// 请求单次分账 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_1&index=1
	WxApiProfitSharing = "wx_api_profit_sharing"
	// 请求多次分账 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_6&index=2
	WxApiMultiProfitSharing = "wx_api_multi_profit_sharing"
	// 查询分账结果 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_2&index=3
	WxApiProfitSharingQuery = "wx_api_profit_sharing_query"
	// 添加分账接收方 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_3&index=4
	WxApiProfitSharingAddReceiver = "wx_api_profit_sharing_add_receiver"
	// 删除分账接收方 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_4&index=5
	WxApiProfitSharingRemoveReceiver = "wx_api_profit_sharing_remove_receiver"
	// 完结分账 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_5&index=6
	WxApiProfitSharingFinish = "wx_api_profit_sharing_finish"
	// 分账回退 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_7&index=7
	WxApiProfitSharingReturn = "wx_api_profit_sharing_return"
	// 回退结果查询 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_8&index=8
	WxApiProfitSharingReturnQuery ="wx_api_profit_sharing_return_query"
	// 分账动账通知 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_9&index=9

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
