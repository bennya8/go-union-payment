package alipay

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

type PayApp struct {
	Base *Base
}

func (a PayApp) Request(params map[string]string) *payloads.UnionPaymentResult {
	resp, err := a.Base.Request("alipay.trade.app.pay", a.BuildParams(params))
	return payloads.NewUnionPaymentResult(err == nil, fmt.Sprintf("%s", err), resp)
}

func (a PayApp) BuildParams(params map[string]string) map[string]string {
	ret := map[string]string{
		"body":         params["body"],
		"subject":      params["subject"],
		"out_trade_no": params["trade_no"],
		"total_amount": params["amount"],
		// 销售产品码，商家和支付宝签约的产品码，为固定值QUICK_MSECURITY_PAY
		"product_code":    "QUICK_MSECURITY_PAY",
		"goods_type":      params["goods_type"],
		"passback_params": params["return_param"],
		// TODO 优惠信息待支持  业务扩展参数，待支持
		// 'promo_params' => '',
		// 'extend_params => '',
		"disable_pay_channels": params["limit_pay"],
		"store_id":             params["store_id"],
	}

	if len(params["timeout_express"]) > 0 {
		// 超时时间 统一使用分钟计算
		ret["timeout_express"] = fmt.Sprintf("%sm", params["timeout_express"])
	}

	return ret
}
