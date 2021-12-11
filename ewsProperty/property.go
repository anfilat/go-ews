package ewsProperty

import (
	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
	"github.com/anfilat/go-ews/internal/xmlWriter"
)

func WriteToXml(property interface{}, writer *xmlWriter.Writer, xmlElementName string) {
	WriteToXmlFull(property, writer, xmlNamespace.Types, xmlElementName)
}

func WriteToXmlFull(property interface{}, writer *xmlWriter.Writer, namespace xmlNamespace.Enum, xmlElementName string) {
	writer.WriteStartElement(namespace, xmlElementName)
	if val, ok := property.(interface {
		WriteAttributesToXml(writer *xmlWriter.Writer)
	}); ok {
		val.WriteAttributesToXml(writer)
	}
	if val, ok := property.(interface {
		WriteElementsToXml(writer *xmlWriter.Writer)
	}); ok {
		val.WriteElementsToXml(writer)
	}
	writer.WriteEndElement()
}
