package go_union_payment

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/bennya8/go-union-payment/contracts"
	"github.com/bennya8/go-union-payment/gateways/alipay"
	"github.com/bennya8/go-union-payment/gateways/cmb"
	"github.com/bennya8/go-union-payment/gateways/qpay"
	"github.com/bennya8/go-union-payment/gateways/wechat"
	"github.com/bennya8/go-union-payment/payloads"
	"io/ioutil"
	"net/http"
	"strings"
)

func NewUnionPayment(gateway payloads.UnionPaymentGateway, config contracts.IGatewayConfig) *UnionPayment {
	instance := &UnionPayment{}
	instance.Gateway = instance.gatewayFactory(gateway, config)
	instance.GatewayIdentify = gateway
	return instance
}

type UnionPayment struct {
	Gateway         contracts.IGateway
	GatewayIdentify payloads.UnionPaymentGateway
}

func (u *UnionPayment) Invoke(api payloads.UnionPaymentApi, params map[string]string) *payloads.UnionPaymentResult {
	return u.Gateway.Request(api, params)
}

func (u *UnionPayment) ParserNotify(req *http.Request, notify contracts.IPaymentNotify) error {
	if u.GatewayIdentify == payloads.WechatGateway {
		b, err := ioutil.ReadAll(req.Body)

		var retKvMap map[string]string
		err = xml.Unmarshal(b, (*contracts.XmlParams)(&retKvMap))
		if err != nil {
			return err
		}

		// refund type
		if _, ok := retKvMap["req_info"]; ok {
			var refund wechat.WxNotifyRefundResponse
			err = xml.Unmarshal(b, &refund)
			if err != nil {
				return err
			}
			notify.PayNotify(payloads.WxNotifyRefund, refund)
			return nil
		} else {
			var pay wechat.WxNotifyPayResponse
			err = xml.Unmarshal(b, &pay)
			if err != nil {
				return err
			}
			notify.PayNotify(payloads.WxNotifyPay, pay)
			return nil
		}
	} else if u.GatewayIdentify == payloads.AlipayGateway {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		var retKvMap = map[string]string{}
		queries := strings.Split(string(b), "&")
		for _, q := range queries {
			kv := strings.Split(q, "=")
			if len(kv) == 2 {
				retKvMap[kv[0]] = kv[1]
			}
		}
		retKvJson, err := json.Marshal(retKvMap)
		if err != nil {
			return err
		}

		var pay alipay.AliNotifyPayResponse
		err = json.Unmarshal(retKvJson, &pay)
		if err != nil {
			return err
		}

		notify.PayNotify(payloads.AliNotifyPay, pay)
		return nil
	} else if u.GatewayIdentify == payloads.QpayGateway {

	} else if u.GatewayIdentify == payloads.CmbGateway {

	}
	return errors.New("unknown gateway identify")
}

func (u *UnionPayment) gatewayFactory(gateway payloads.UnionPaymentGateway, config contracts.IGatewayConfig) contracts.IGateway {
	if gateway == payloads.WechatGateway {
		return &wechat.Gateway{Base: wechat.NewBase(config.(*wechat.Config))}
	} else if gateway == payloads.AlipayGateway {
		return &alipay.Gateway{Base: alipay.NewBase(config.(*alipay.Config))}
	} else if gateway == payloads.QpayGateway {
		return &qpay.Gateway{Base: qpay.NewBase(config.(*qpay.Config))}
	} else if gateway == payloads.CmbGateway {
		return &cmb.Gateway{Base: cmb.NewBase(config.(*cmb.Config))}
	}
	panic("unknown gateway")
}
