package ewsProperty

import (
	"github.com/anfilat/go-ews/internal/base"
	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
)

func WriteToXml(property interface{}, writer base.Writer, xmlElementName string) {
	WriteToXmlFull(property, writer, xmlNamespace.Types, xmlElementName)
}

func WriteToXmlFull(property interface{}, writer base.Writer, namespace xmlNamespace.Enum, xmlElementName string) {
	writer.WriteStartElement(namespace, xmlElementName)

	if val, ok := property.(base.WriterAttributesToXml); ok {
		val.WriteAttributesToXml(writer)
	}
	if val, ok := property.(base.WriterElementsToXml); ok {
		val.WriteElementsToXml(writer)
	}

	writer.WriteEndElement()
}
