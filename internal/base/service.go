package base

import "github.com/anfilat/go-ews/enumerations/exchangeVersion"

type Service interface {
	GetVersion() exchangeVersion.Enum
}
