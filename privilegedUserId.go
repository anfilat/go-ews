package ews

import (
	"github.com/anfilat/go-ews/enumerations/connectingIdType"
	"github.com/anfilat/go-ews/enumerations/privilegedLogonType"
	"github.com/anfilat/go-ews/enumerations/privilegedUserIdBudgetType"
)

type PrivilegedUserId struct {
	LogonType privilegedLogonType.Enum
	IdType    connectingIdType.Enum
	Id        string
	//nolint:structcheck,unused
	budgetType privilegedUserIdBudgetType.Enum
}

func NewPrivilegedUserId(logonType privilegedLogonType.Enum, idType connectingIdType.Enum, id string) PrivilegedUserId {
	return PrivilegedUserId{
		LogonType: logonType,
		IdType:    idType,
		Id:        id,
	}
}
