package alipay

import (
	"errors"
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

type PayBar struct {
	Base *Base
}

func (a PayBar) Request(params map[string]string) *payloads.UnionPaymentResult {
	uri := a.Base.GetFullGatewayUrl("alipay.trade.pay")

	err := a.CheckParams(params)
	if err != nil {
		return payloads.NewUnionPaymentResult(false, fmt.Sprintf("%s", err), nil)
	}

	resp, err := a.Base.Request(uri, a.BuildParams(params))
	return payloads.NewUnionPaymentResult(err == nil, fmt.Sprintf("%s", err), resp)
}

func (a PayBar) BuildParams(params map[string]string) map[string]string {
	ret := map[string]string{
		"out_trade_no": params["trade_no"],
		"scene":        params["scene"],
		"auth_code":    params["auth_code"],
		"product_code": "FACE_TO_FACE_PAYMENT",
		"subject":      params["subject"],
		// TODO 支付宝用户ID
		// 'seller_id' => $this->partner,
		"body":         params["body"],
		"total_amount": params["amount"],
		// TODO 折扣金额
		// 'discountable_amount' => '',
		// TODO  业务扩展参数 订单商品列表信息，待支持
		// 'extend_params => '',
		// 'goods_detail' => '',
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

func (a PayBar) CheckParams(params map[string]string) error {

	if len(params["scene"]) <= 0 || (params["scene"] == "bar_code" || params["scene"] == "wave_code") {
		return errors.New("支付场景 scene 必须设置 条码支付：bar_code 声波支付：wave_code")
	}

	if len(params["auth_code"]) <= 0 {
		return errors.New("支付授权码 必须设置")
	}

	return nil
}
