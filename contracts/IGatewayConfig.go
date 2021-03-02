package contracts

type IGatewayConfig interface {
	ParseConfig() interface{}
	CheckConfig() error
}
