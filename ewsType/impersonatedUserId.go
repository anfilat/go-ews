package ewsType

import (
	"github.com/anfilat/go-ews/enumerations/connectingIdType"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
	"github.com/anfilat/go-ews/internal/errors"
	"github.com/anfilat/go-ews/internal/xmlWriter"
)

type ImpersonatedUserId struct {
	IdType connectingIdType.Enum
	Id     string
}

func NewImpersonatedUserId(idType connectingIdType.Enum, id string) *ImpersonatedUserId {
	return &ImpersonatedUserId{
		IdType: idType,
		Id:     id,
	}
}

func (u *ImpersonatedUserId) Validate() error {
	if u.Id == "" {
		return errors.NewValidateError("the Id property of ImpersonatedUserId must be set")
	}
	return nil
}

func (u *ImpersonatedUserId) WriteToXml(writer *xmlWriter.Writer, version exchangeVersion.Enum) {
	writer.WriteStartElement(xmlNamespace.Types, "ExchangeImpersonation")
	writer.WriteStartElement(xmlNamespace.Types, "ConnectingSID")

	connectingIdTypeLocalName := u.IdType.String()
	if u.IdType == connectingIdType.SmtpAddress && version == exchangeVersion.Exchange2007SP1 {
		connectingIdTypeLocalName = "PrimarySmtpAddress"
	}

	writer.WriteElementValue(xmlNamespace.Types, connectingIdTypeLocalName, u.Id)

	writer.WriteEndElement()
	writer.WriteEndElement()
}
