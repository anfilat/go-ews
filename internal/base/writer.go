package base

import "github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"

type Writer interface {
	GetService() Service
	WriteStartElement(namespace xmlNamespace.Enum, localName string)
	WriteEndElement()
	WriteElementValue(namespace xmlNamespace.Enum, localName string, value interface{})
	WriteAttributeValueBool(localName string, alwaysWriteEmptyString bool, value interface{})
	WriteAttributeValue(namespacePrefix, localName string, value interface{})
	WriteAttributeString(namespacePrefix, localName, value string)
	GetIsTimeZoneHeaderEmitted() bool
	SetIsTimeZoneHeaderEmitted(val bool)
}

type WriterToXml interface {
	WriteToXml(writer Writer)
}

type WriterAttributesToXml interface {
	WriteAttributesToXml(writer Writer)
}

type WriterElementsToXml interface {
	WriteElementsToXml(writer Writer)
}
