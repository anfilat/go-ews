package base

import "github.com/anfilat/go-ews/enumerations/exchangeVersion"

type Request interface {
	Validator
	GetXmlElementName() string
	GetMinimumRequiredServerVersion() exchangeVersion.Enum
}
