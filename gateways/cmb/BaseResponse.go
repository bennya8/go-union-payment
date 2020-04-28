package cmb

type BaseResponse struct {
	Resp string
}

func (b BaseResponse) ToMap() (map[string]interface{}, error) {
	panic("implement me")
}

func (b BaseResponse) ToJson() (string, error) {
	panic("implement me")
}

func (b BaseResponse) ToXml() (string, error) {
	panic("implement me")
}

func NewBaseResponse(resp string) *BaseResponse {
	return &BaseResponse{Resp: resp}
}
