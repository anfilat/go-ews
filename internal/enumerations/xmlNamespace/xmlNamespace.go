package xmlNamespace

// Enum the namespaces as used by the EwsXmlReader, EwsServiceXmlReader, and EwsServiceXmlWriter classes.
type Enum int

const (
	// NotSpecified - the namespace is not specified
	NotSpecified = iota
	// Messages - the EWS Messages namespace.
	Messages
	// Types - the EWS Types namespace.
	Types
	// Errors - the EWS Errors namespace.
	Errors
	// Soap - the SOAP 1.1 namespace.
	Soap
	// Soap12 - the SOAP 1.2 namespace.
	Soap12
	// XmlSchemaInstance - XmlSchema-Instance namespace.
	XmlSchemaInstance
	// PassportSoapFault - the Passport SOAP services SOAP fault namespace.
	PassportSoapFault
	// WSTrustFebruary2005 - the WS-Trust February 2005 namespace.
	WSTrustFebruary2005
	// WSAddressing - the WS Addressing 1.0 namespace.
	WSAddressing
	// Autodiscover - the Autodiscover SOAP service namespace.
	Autodiscover
)
