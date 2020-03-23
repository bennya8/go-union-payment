package wechat

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
)

type PayWap struct {
	Base *Base
}

func (w PayWap) Request(params map[string]string) *payloads.UnionPaymentResult {
	uri := w.Base.GetFullGatewayUrl("pay/unifiedorder")

	//api

	resp, err := w.Base.Request(uri, w.BuildParams(params))
	return payloads.NewUnionPaymentResult(err != nil, fmt.Sprintf("%s", err), resp)
}

func (w PayWap) BuildParams(params map[string]string) map[string]string {
	ret := map[string]string{
		"app_id":     w.Base.Config.AppId,
		"sub_appid":  w.Base.Config.SubAppId,
		"mch_id":     w.Base.Config.MchId,
		"sub_mch_id": w.Base.Config.SubMchId,
		"nonce_str":  w.Base.NonceStr,
		"sign_type":  w.Base.SignType,
	}
	for k, v := range params {
		ret[k] = v
	}
	return ret
}
