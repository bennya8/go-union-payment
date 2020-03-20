package contracts

type IGateway interface {
	Request()
	ResponseRaw()
	ResponseToJson()
	ResponseToXml()
}
