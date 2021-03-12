package alipay

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

type PayWeb struct {
	Base *Base
}

func (a PayWeb) Request(params map[string]string) *payloads.UnionPaymentResult {
	resp, err := a.Base.Request("alipay.trade.page.pay", a.BuildParams(params))
	return payloads.NewUnionPaymentResult(err == nil, fmt.Sprintf("%s", err), resp)
}

func (a PayWeb) BuildParams(params map[string]string) map[string]string {

	ret := map[string]string{
		"out_trade_no": params["trade_no"],
		"product_code": "FAST_INSTANT_TRADE_PAY", // 销售产品码，商家和支付宝签约的产品码，为固定值QUICK_WAP_PAY
		"total_amount": params["amount"],
		"subject":      params["subject"],
		"body":         params["body"],
		// 'goods_detail' => '', // @TODO 订单包含的商品列表信息 待实现
		"passback_params": params["return_param"],
		// 'extend_params => '', // @TODO 业务扩展参数，待支持
		"goods_type":           params["goods_type"],
		"disable_pay_channels": params["limit_pay"],
		"store_id":             params["store_id"],
		"qr_pay_mode":          params["qr_pay_mode"],
		// 'qrcode_width' => '' // @TODO 设置二维码宽度，qr_pay_mode = 4时有效。设置二维码宽度
	}

	if len(params["timeout_express"]) > 0 {
		// 超时时间 统一使用分钟计算
		ret["timeout_express"] = fmt.Sprintf("%sm", params["timeout_express"])
	}

	return ret
}
