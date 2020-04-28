package cmb

import (
	"github.com/bennya8/go-union-payment/payloads"
)

type PayApp struct {
	Base *Base
}

func (c PayApp) Request(params map[string]string) *payloads.UnionPaymentResult {
	// no need to request

	j := `{"version":"1.0","charset":"UTF-8","signType":"SHA-256","reqData":{"dateTime":"20200428184425","branchNo":"xxxx","merchantNo":"xxxxxx","date":"20200428","orderNo":"15880706655477","amount":1,"expireTimeSpan":1000,"payNoticeUrl":"https:\/\/dayutalk.cn\/notify\/cmb","payNoticePara":"","clientIP":"","cardType":"A","subMerchantNo":"","subMerchantName":"","subMerchantTPCode":"","subMerchantTPName":"","payModeType":"","agrNo":"","merchantSerialNo":"","userID":"","mobile":"","lon":"","lat":"","riskLevel":"","signNoticeUrl":"https:\/\/dayutalk.cn\/notify\/cmb","signNoticePara":""},"sign":"65633262316138313362376361616433313761333763383333393133646262643432316261316137643666393862633139363038373331663532316439616563"}`

	return payloads.NewUnionPaymentResult(true, "", NewBaseResponse(j))
}

func (c PayApp) BuildParams(params map[string]string) map[string]string {
	return map[string]string{
	}
}
