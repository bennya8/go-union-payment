package wechat

import (
	"fmt"
	"github.com/bennya8/go-union-payment/payloads"
	"github.com/bennya8/go-union-payment/utils/date"
	"time"
)

/**
 * @package gateway.wechat
 * @author  : Benny
 * @email   : benny_a8@qq.com
 * @date    : 2020/04/11
 **/
type CommentQuery struct {
	Base *Base
}

func (w CommentQuery) Request(params map[string]string) *payloads.UnionPaymentResult {
	uri := w.Base.GetFullGatewayUrl("billcommentsp/batchquerycomment")

	//api

	resp, err := w.Base.Request(uri, w.BuildParams(params))
	return payloads.NewUnionPaymentResult(err != nil, fmt.Sprintf("%s", err), resp)
}

func (w CommentQuery) BuildParams(params map[string]string) map[string]string {
	ret := map[string]string{
		"begin_time": date.TimeFormat(time.Now().AddDate(0, 0, -1), "YYYYMMDDhhiiss"),
		"end_time":   date.TimeFormat(time.Now(), "YYYYMMDDhhiiss"),
		"offset":     "0",
		"limit":      "200",
	}
	for k, v := range params {
		ret[k] = v
	}
	return ret
}
