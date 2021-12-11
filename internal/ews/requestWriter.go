package ews

import (
	"github.com/anfilat/go-ews/enumerations/dateTimePrecision"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/internal/base"
	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
	"github.com/anfilat/go-ews/internal/xmlWriter"
)

type RequestWriter struct {
	w                       *xmlWriter.Writer
	sd                      *ServiceData
	request                 base.Request
	isTimeZoneHeaderEmitted bool
}

func NewRequestWriter(sd *ServiceData, request base.Request) *RequestWriter {
	return &RequestWriter{
		w:       xmlWriter.New(),
		sd:      sd,
		request: request,
	}
}

func (w *RequestWriter) GetService() base.Service {
	return w.sd
}

func (w *RequestWriter) WriteStartElement(namespace xmlNamespace.Enum, localName string) {
	w.w.WriteStartElement(namespace, localName)
}

func (w *RequestWriter) WriteEndElement() {
	w.w.WriteEndElement()
}

func (w *RequestWriter) WriteElementValue(namespace xmlNamespace.Enum, localName string, value interface{}) {
	w.w.WriteElementValue(namespace, localName, value)
}

func (w *RequestWriter) WriteAttributeValueBool(localName string, alwaysWriteEmptyString bool, value interface{}) {
	w.w.WriteAttributeValueBool(localName, alwaysWriteEmptyString, value)
}

func (w *RequestWriter) WriteAttributeValue(namespacePrefix, localName string, value interface{}) {
	w.w.WriteAttributeValue(namespacePrefix, localName, value)
}

func (w *RequestWriter) WriteAttributeString(namespacePrefix, localName, value string) {
	w.w.WriteAttributeString(namespacePrefix, localName, value)
}

func (w *RequestWriter) GetIsTimeZoneHeaderEmitted() bool {
	return w.isTimeZoneHeaderEmitted
}

func (w *RequestWriter) SetIsTimeZoneHeaderEmitted(val bool) {
	w.isTimeZoneHeaderEmitted = val
}

func (w *RequestWriter) WriteXML() ([]byte, error) {
	w.w.WriteDoc()
	w.writeRoot()

	w.writeHeader()
	w.writeBody()

	w.WriteEndElement()
	w.w.WriteEndDoc()

	w.w.Flush()

	if err := w.w.Err(); err != nil {
		return nil, err
	}
	return w.w.Bytes(), nil
}

func (w *RequestWriter) writeRoot() {
	w.w.StartRoot()
	w.WriteStartElement(xmlNamespace.Soap, "Envelope")
	w.WriteAttributeValue("xmlns", "xsi", "http://www.w3.org/2001/XMLSchema-instance")
	w.WriteAttributeValue("xmlns", "m", "http://schemas.microsoft.com/exchange/services/2006/messages")
	w.WriteAttributeValue("xmlns", "t", "http://schemas.microsoft.com/exchange/services/2006/types")
	w.w.EndRoot()
}

func (w *RequestWriter) writeHeader() {
	w.WriteStartElement(xmlNamespace.Soap, "Header")

	w.writeVersionHeader()
	// w.writeTimeZoneHeader()
	w.writeDateTimePrecision()
	w.writeUserOrRoles()

	w.WriteEndElement()
}

func (w *RequestWriter) writeVersionHeader() {
	w.WriteStartElement(xmlNamespace.Types, "RequestServerVersion")
	w.WriteAttributeValueBool("Version", false, w.sd.GetRequestedServiceVersionString())
	w.WriteEndElement()
}

func (w *RequestWriter) writeDateTimePrecision() {
	if w.sd.DateTimePrecision != dateTimePrecision.Default {
		w.WriteElementValue(xmlNamespace.Types, "DateTimePrecision", w.sd.DateTimePrecision.String())
	}
}

func (w *RequestWriter) writeUserOrRoles() {
	if w.sd.ImpersonatedUserId != nil {
		w.sd.ImpersonatedUserId.WriteToXml(w)
	} else if w.sd.PrivilegedUserId != nil {
		w.sd.PrivilegedUserId.WriteToXml(w)
	} else if w.sd.ManagementRoles != nil {
		w.sd.ManagementRoles.WriteToXml(w)
	}
}

//nolint:unused
func (w *RequestWriter) writeTimeZoneHeader() {
	if w.sd.Exchange2007CompatibilityMode {
		return
	}

	if w.sd.Version == exchangeVersion.Exchange2007SP1 || w.isEmitTimeZoneHeader() {
		w.WriteStartElement(xmlNamespace.Types, "TimeZoneContext")

		w.WriteEndElement()
		w.SetIsTimeZoneHeaderEmitted(true)
	}
}

//nolint:unused
func (w *RequestWriter) isEmitTimeZoneHeader() bool {
	if val, ok := w.request.(interface{ EmitTimeZoneHeader() bool }); ok {
		return val.EmitTimeZoneHeader()
	}
	return false
}

func (w *RequestWriter) writeBody() {
	w.WriteStartElement(xmlNamespace.Soap, "Body")

	w.WriteEndElement()
}
