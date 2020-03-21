package wechat

import (
	"encoding/json"
	"encoding/xml"
	"github.com/bennya8/go-union-payment/contracts"
)

type BaseResponse struct {
	Resp string
}

func NewBaseResponse(resp string) *BaseResponse {
	return &BaseResponse{Resp: resp}
}

func (w *BaseResponse) ToMap() (map[string]interface{}, error) {
	var retKvMap map[string]string
	err := xml.Unmarshal([]byte(w.Resp), (*contracts.XmlParams)(&retKvMap))

	var ret map[string]interface{}
	retBytes, err := json.Marshal(retKvMap)
	if err != nil {
		return ret, nil
	}
	err = json.Unmarshal(retBytes, &ret)
	if err != nil {
		return ret, nil
	}
	return ret, err
}

func (w *BaseResponse) ToJson() (string, error) {
	dict, err := w.ToMap()
	if err != nil {
		return "", err
	}
	ret, err := json.Marshal(dict)
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

func (w *BaseResponse) ToXml() (string, error) {
	// default is xml, no need to convert
	return w.Resp, nil
}
