package wechat

import "github.com/bennya8/go-union-payment/contracts"

const WapChargeMethod = "pay/unifiedorder"

type WapCharge struct {
	Base *Base
}

func (w WapCharge) Request() {

	params := w.BuildParams()
	xmlParams := contracts.XmlParams(params)


	w.Base.Request()
}

func (w WapCharge) Response() {
	panic("implement me")
}

func (w WapCharge) BuildParams() map[string]string {
	ret := map[string]string{
		"app_id":     w.Base.Config.AppId,
		"sub_appid":  w.Base.Config.SubAppId,
		"mch_id":     w.Base.Config.MchId,
		"sub_mch_id": w.Base.Config.SubMchId,
		"nonce_str":  w.Base.NonceStr,
		"sign_type":  w.Base.SignType,
	}



	return ret
}
