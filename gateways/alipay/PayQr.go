package alipay

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

type PayQr struct {
	Base *Base
}

func (a PayQr) Request(params map[string]string) *payloads.UnionPaymentResult {
	resp, err := a.Base.Request("alipay.trade.precreate", a.BuildParams(params))
	return payloads.NewUnionPaymentResult(err == nil, fmt.Sprintf("%s", err), resp)
}

func (a PayQr) BuildParams(params map[string]string) map[string]string {
	ret := map[string]string{
		"out_trade_no": params["trade_no"],
		// TODO 卖家支付宝id
		// 'seller_id' => '',
		"total_amount": params["amount"],
		// TODO 折扣金额
		// 'discountable_amount' => '',
		// TODO  业务扩展参数 订单商品列表信息，待支持
		// 'extend_params => '',
		// 'goods_detail' => '',
		"subject":     params["subject"],
		"body":        params["body"],
		"operator_id": params["operator_id"],
		"store_id":    params["store_id"],
		"terminal_id": params["terminal_id"],
	}

	if len(params["timeout_express"]) > 0 {
		// 超时时间 统一使用分钟计算
		ret["timeout_express"] = fmt.Sprintf("%sm", params["timeout_express"])
	}

	return ret
}
