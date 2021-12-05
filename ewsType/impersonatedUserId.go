package ewsType

import (
	"github.com/anfilat/go-ews/enumerations/connectingIdType"
)

type ImpersonatedUserId struct {
	IdType connectingIdType.Enum
	Id     string
}

func NewImpersonatedUserId(idType connectingIdType.Enum, id string) ImpersonatedUserId {
	return ImpersonatedUserId{
		IdType: idType,
		Id:     id,
	}
}
