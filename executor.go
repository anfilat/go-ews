package ews

import (
	"context"
	"fmt"

	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
	"github.com/anfilat/go-ews/internal/requests"
	"github.com/anfilat/go-ews/internal/xmlWriter"
)

//nolint:unparam
func execute(_ context.Context, service *ExchangeService, request requests.Request) (interface{}, error) {
	if err := service.validate(); err != nil {
		return nil, err
	}
	if err := request.Validate(); err != nil {
		return nil, err
	}

	service.ensureClient()
	buf, err := writeXML(service, request)
	fmt.Println(string(buf))
	fmt.Println(err)

	return nil, nil
}

func writeXML(service *ExchangeService, request requests.Request) ([]byte, error) {
	w := xmlWriter.New()

	w.WriteDoc()
	w.WriteStartElement(xmlNamespace.Soap, "Envelope")
	w.WriteAttributeValue("xmlns", "xsi", "http://www.w3.org/2001/XMLSchema-instance")
	w.WriteAttributeValue("xmlns", "m", "http://schemas.microsoft.com/exchange/services/2006/messages")
	w.WriteAttributeValue("xmlns", "t", "http://schemas.microsoft.com/exchange/services/2006/types")

	w.WriteStartElement(xmlNamespace.Soap, "Header")

	w.WriteStartElement(xmlNamespace.Types, "RequestServerVersion")
	w.WriteAttributeValueBool("Version", false, service.getRequestedServiceVersionString())
	w.WriteEndElement()

	emitTimeZoneHeader, ok := request.(interface{ EmitTimeZoneHeader() bool })
	if (service.version == exchangeVersion.Exchange2007SP1 || (ok && emitTimeZoneHeader.EmitTimeZoneHeader())) &&
		!service.Exchange2007CompatibilityMode {
		w.WriteStartElement(xmlNamespace.Types, "TimeZoneContext")

		w.WriteEndElement()
		w.IsTimeZoneHeaderEmitted = true
	}

	w.WriteEndElement()
	w.WriteEndDoc()

	w.Flush()

	if err := w.Err(); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
