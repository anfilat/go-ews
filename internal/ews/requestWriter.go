package ews

import (
	"github.com/anfilat/go-ews/enumerations/dateTimePrecision"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
	"github.com/anfilat/go-ews/internal/requests"
	"github.com/anfilat/go-ews/internal/xmlWriter"
)

type RequestWriter struct {
	w       *xmlWriter.Writer
	Service *Data
	request requests.Request
}

func NewRequestWriter(service *Data, request requests.Request) *RequestWriter {
	return &RequestWriter{
		w:       xmlWriter.New(),
		Service: service,
		request: request,
	}
}

func (w *RequestWriter) WriteXML() ([]byte, error) {
	w.w.WriteDoc()
	w.writeRoot()

	w.writeHeader()
	w.writeBody()

	w.w.WriteEndElement()
	w.w.WriteEndDoc()

	w.w.Flush()

	if err := w.w.Err(); err != nil {
		return nil, err
	}
	return w.w.Bytes(), nil
}

func (w *RequestWriter) writeRoot() {
	w.w.StartRoot()
	w.w.WriteStartElement(xmlNamespace.Soap, "Envelope")
	w.w.WriteAttributeValue("xmlns", "xsi", "http://www.w3.org/2001/XMLSchema-instance")
	w.w.WriteAttributeValue("xmlns", "m", "http://schemas.microsoft.com/exchange/services/2006/messages")
	w.w.WriteAttributeValue("xmlns", "t", "http://schemas.microsoft.com/exchange/services/2006/types")
	w.w.EndRoot()
}

func (w *RequestWriter) writeHeader() {
	w.w.WriteStartElement(xmlNamespace.Soap, "Header")

	w.writeVersionHeader()
	w.writeTimeZoneHeader()
	w.writeDateTimePrecision()
	w.writeUserOrRoles()

	w.w.WriteEndElement()
}

func (w *RequestWriter) writeVersionHeader() {
	w.w.WriteStartElement(xmlNamespace.Types, "RequestServerVersion")
	w.w.WriteAttributeValueBool("Version", false, w.Service.GetRequestedServiceVersionString())
	w.w.WriteEndElement()
}

func (w *RequestWriter) writeDateTimePrecision() {
	if w.Service.DateTimePrecision != dateTimePrecision.Default {
		w.w.WriteElementValue(xmlNamespace.Types, "DateTimePrecision", w.Service.DateTimePrecision.String())
	}
}

func (w *RequestWriter) writeUserOrRoles() {
	if w.Service.ImpersonatedUserId != nil {
		w.Service.ImpersonatedUserId.WriteToXml(w.w, w.Service.Version)
	} else if w.Service.PrivilegedUserId != nil {
		w.Service.PrivilegedUserId.WriteToXml(w.w, w.Service.Version)
	} else if w.Service.ManagementRoles != nil {
		w.Service.ManagementRoles.WriteToXml(w.w)
	}
}

func (w *RequestWriter) writeTimeZoneHeader() {
	if w.Service.Exchange2007CompatibilityMode {
		return
	}

	if w.Service.Version == exchangeVersion.Exchange2007SP1 || w.isEmitTimeZoneHeader() {
		w.w.WriteStartElement(xmlNamespace.Types, "TimeZoneContext")

		w.w.WriteEndElement()
		w.w.IsTimeZoneHeaderEmitted = true
	}
}

func (w *RequestWriter) isEmitTimeZoneHeader() bool {
	if val, ok := w.request.(interface{ EmitTimeZoneHeader() bool }); ok {
		return val.EmitTimeZoneHeader()
	}
	return false
}

func (w *RequestWriter) writeBody() {
	w.w.WriteStartElement(xmlNamespace.Soap, "Body")

	w.w.WriteEndElement()
}
