package ewsType

import (
	"github.com/anfilat/go-ews/enumerations/connectingIdType"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/enumerations/privilegedLogonType"
	"github.com/anfilat/go-ews/enumerations/privilegedUserIdBudgetType"
	"github.com/anfilat/go-ews/internal/base"
	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
	"github.com/anfilat/go-ews/internal/errors"
)

type PrivilegedUserId struct {
	LogonType  privilegedLogonType.Enum
	IdType     connectingIdType.Enum
	Id         string
	BudgetType *privilegedUserIdBudgetType.Enum
}

func NewPrivilegedUserId(logonType privilegedLogonType.Enum, idType connectingIdType.Enum, id string) PrivilegedUserId {
	return PrivilegedUserId{
		LogonType: logonType,
		IdType:    idType,
		Id:        id,
	}
}

func (u *PrivilegedUserId) Validate() error {
	if u.Id == "" {
		return errors.NewValidateError("the Id property of PrivilegedUserId must be set")
	}
	return nil
}

func (u *PrivilegedUserId) WriteToXml(writer base.Writer) {
	writer.WriteStartElement(xmlNamespace.Types, "OpenAsAdminOrSystemService")
	writer.WriteAttributeString("", "LogonType", u.LogonType.String())
	if writer.GetService().GetVersion() >= exchangeVersion.Exchange2013 && u.BudgetType != nil {
		writer.WriteAttributeString("", "BudgetType", u.BudgetType.String())
	}

	writer.WriteStartElement(xmlNamespace.Types, "ConnectingSID")
	writer.WriteElementValue(xmlNamespace.Types, u.IdType.String(), u.Id)
	writer.WriteEndElement()

	writer.WriteEndElement()
}
