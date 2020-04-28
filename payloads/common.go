package payloads

type (
	UnionPaymentGateway string
	UnionPaymentNotify  string
	UnionPaymentApi     string
)

const (
	Version = "1.0.0.b1"

	// gateways
	AlipayGateway UnionPaymentGateway = "Alipay"
	CmbGateway                        = "CMBank"
	WechatGateway                     = "Wechat"
	QpayGateway                       = "Qpay"

	// notifies
	WxNotifyPay           UnionPaymentNotify = "wx_notify_pay"            // 支付结果通知 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_7&index=8
	WxNotifyRefund                           = "wx_notify_refund"         // 退款结果通知 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_16&index=10
	WxNotifyProfitSharing                    = "wx_notify_profit_sharing" // 分账动账通知 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_9&index=9
	AliNotifyPay                             = "ali_notify_pay"

	// apis
	// 微信支付 API
	WxApiPayApp       UnionPaymentApi = "wx_api_pay_app"       // 统一下单-APP支付 https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_1
	WxApiPayWap                       = "wx_api_pay_wap"       // 统一下单-WAP支付 https://pay.weixin.qq.com/wiki/doc/api/H5.php?chapter=9_20&index=1
	WxApiPayPub                       = "wx_api_pay_pub"       // 统一下单-公众号支付 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
	WxApiPayLite                      = "wx_api_pay_lite"      // 统一下单-小程序 https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1
	WxApiPayQr                        = "wx_api_pay_qr"        // 统一下单-QR https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_10&index=1
	WxApiPayBar                       = "wx_api_pay_bar"       // 统一下单-NATIVE https://pay.weixin.qq.com/wiki/doc/api/native.php?chapter=9_1
	WxApiCloseTrade                   = "wx_api_close_trade"   // 关闭订单 (App\Wap\Lite\Qr\Bar） https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
	WxApiQueryTrade                   = "wx_api_query_trader"  // 查询订单 (App\Wap\Lite\Qr\Bar） https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
	WxApiCancelTrade                  = "wx_api_cancel_trade"  // 撤销订单（Bar）https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=9_11&index=3
	WxApiRefund                       = "wx_api_refund"        // 申请退款 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
	WxApiQueryRefund                  = "wx_api_query_refund"  // 查询退款 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
	WxApiBillDownload                 = "wx_api_bill_download" // 下载交易账单 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_6
	WxApiFundDownload                 = "wx_api_fund_download" // 下载资金账单 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_18&index=7

	// 微信支付 - 企业付款 API
	WxApiTransfer          = "wx_api_transfer"            // 企业付款 零钱 https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
	WxApiQueryTransfer     = "wx_api_query_transfer"      // 查询企业付款 https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_3
	WxApiTransferBank      = "wx_api_transfer_bank"       // 企业付款 银行卡 https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_2
	WxApiQueryTransferBank = "wx_api_query_transfer_bank" // 查询企业付款 https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_3

	// 微信支付 - 分账 API
	WxApiProfitSharing               = "wx_api_profit_sharing"                 // 请求单次分账 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_1&index=1
	WxApiMultiProfitSharing          = "wx_api_multi_profit_sharing"           // 请求多次分账 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_6&index=2
	WxApiProfitSharingQuery          = "wx_api_profit_sharing_query"           // 查询分账结果 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_2&index=3
	WxApiProfitSharingAddReceiver    = "wx_api_profit_sharing_add_receiver"    // 添加分账接收方 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_3&index=4
	WxApiProfitSharingRemoveReceiver = "wx_api_profit_sharing_remove_receiver" // 删除分账接收方 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_4&index=5
	WxApiProfitSharingFinish         = "wx_api_profit_sharing_finish"          // 完结分账 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_5&index=6
	WxApiProfitSharingReturn         = "wx_api_profit_sharing_return"          // 分账回退 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_7&index=7
	WxApiProfitSharingReturnQuery    = "wx_api_profit_sharing_return_query"    // 回退结果查询 https://pay.weixin.qq.com/wiki/doc/api/allocation.php?chapter=27_8&index=8

	// 支付宝 - API
	AliApiPayApp = "ali_api_pay_app" // app支付接口2.0 https://opendocs.alipay.com/apis/api_1/alipay.trade.app.pay
	AliApiPayWap = "ali_api_pay_wap" // 手机网站支付接口2.0 https://opendocs.alipay.com/apis/api_1/alipay.trade.wap.pay
	AliApiPayWeb = "ali_api_pay_web" // PC网站支付接口 https://opendocs.alipay.com/apis/api_1/alipay.trade.page.pay
	AliApiPayBar = "ali_api_pay_bar" // 统一收单交易支付接口 https://opendocs.alipay.com/apis/api_1/alipay.trade.pay
	AliApiPayQr  = "ali_api_pay_qr"  // 统一收单线下交易预创建 https://opendocs.alipay.com/apis/api_1/alipay.trade.precreate

	// QQ钱包 - API
	QpayApiPayApp      = "qpay_api_pay_app" // 统一下单-App支付 https://qpay.qq.com/buss/wiki/38/1203
	QpayApiPayPub      = "qpay_api_pay_pub" // 统一下单-公众号支付 https://qpay.qq.com/buss/wiki/38/1203
	QpayApiPayQr       = "qpay_api_pay_qr"  // 统一下单-扫码支付 https://qpay.qq.com/buss/wiki/38/1203
	QpayApiPayBar      = "qpay_api_pay_bar" // 付款码支付 https://qpay.qq.com/buss/wiki/1/1122
	QpayApiRefund      = "qpay_api_refund"
	QpayApiQueryRefund = "qpay_api_query_refund"

	// 招商银行 - API
	CmbApiPayApp  = "cmb_api_pay_app"  // 招行 APP
	CmbApiPayWap  = "cmb_api_pay_wap"  // 招行 H5支付
	CmbApiPayWeb  = "cmb_api_pay_web"  // 招行 PC支付
	CmbApiPayQr   = "cmb_api_pay_qr"   // 招行 QR
	CmbApiPayLite = "cmb_api_pay_lite" // 招行 小程序
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
