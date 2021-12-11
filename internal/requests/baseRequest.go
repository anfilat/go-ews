package requests

import "github.com/anfilat/go-ews/enumerations/exchangeVersion"

type Request interface {
	Validate() error
	GetXmlElementName() string
	GetMinimumRequiredServerVersion() exchangeVersion.Enum
}
