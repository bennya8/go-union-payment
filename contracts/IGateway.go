package contracts

import (
	"encoding/xml"
	"github.com/bennya8/go-union-payment/payloads"
	"io"
)

type IGateway interface {
	Request() *payloads.UnionPaymentResult
}

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

type XmlParams map[string]string

func (m XmlParams) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (m *XmlParams) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = XmlParams{}
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}
