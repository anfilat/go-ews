package utils

import (
	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
)

func GetNamespacePrefix(namespace xmlNamespace.Enum) string {
	switch namespace {
	case xmlNamespace.Types:
		return "t"
	case xmlNamespace.Messages:
		return "m"
	case xmlNamespace.Errors:
		return "e"
	case xmlNamespace.Soap, xmlNamespace.Soap12:
		return "soap"
	case xmlNamespace.XmlSchemaInstance:
		return "xsi"
	case xmlNamespace.PassportSoapFault:
		return "psf"
	case xmlNamespace.WSTrustFebruary2005:
		return "wst"
	case xmlNamespace.WSAddressing:
		return "wsa"
	case xmlNamespace.Autodiscover:
		return "a"
	case xmlNamespace.NotSpecified:
		return ""
	}
	return ""
}

func GetNamespaceUri(namespace xmlNamespace.Enum) string {
	switch namespace {
	case xmlNamespace.Types:
		return "http://schemas.microsoft.com/exchange/services/2006/types"
	case xmlNamespace.Messages:
		return "http://schemas.microsoft.com/exchange/services/2006/messages"
	case xmlNamespace.Errors:
		return "http://schemas.microsoft.com/exchange/services/2006/errors"
	case xmlNamespace.Soap:
		return "http://schemas.xmlsoap.org/soap/envelope/"
	case xmlNamespace.Soap12:
		return "http://www.w3.org/2003/05/soap-envelope"
	case xmlNamespace.XmlSchemaInstance:
		return "http://www.w3.org/2001/XMLSchema-instance"
	case xmlNamespace.PassportSoapFault:
		return "http://schemas.microsoft.com/Passport/SoapServices/SOAPFault"
	case xmlNamespace.WSTrustFebruary2005:
		return "http://schemas.xmlsoap.org/ws/2005/02/trust"
	case xmlNamespace.WSAddressing:
		return "http://www.w3.org/2005/08/addressing"
	case xmlNamespace.Autodiscover:
		return "http://schemas.microsoft.com/exchange/2010/Autodiscover"
	case xmlNamespace.NotSpecified:
		return ""
	}
	return ""
}
