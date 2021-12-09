package ewsType

import (
	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
	"github.com/anfilat/go-ews/internal/errors"
	"github.com/anfilat/go-ews/internal/validate"
	"github.com/anfilat/go-ews/internal/xmlWriter"
)

type ManagementRoles struct {
	userRoles        []string
	applicationRoles []string
}

func NewManagementRoles(userRoles, applicationRoles []string) *ManagementRoles {
	return &ManagementRoles{
		userRoles:        userRoles,
		applicationRoles: applicationRoles,
	}
}

func (m *ManagementRoles) Validate() error {
	if len(m.userRoles) == 0 && len(m.applicationRoles) == 0 {
		return errors.NewValidateError("the roles of ManagementRoles must be set")
	}
	if len(m.userRoles) > 0 {
		return validate.ParamSlice(m.userRoles, "userRoles")
	}
	return validate.ParamSlice(m.applicationRoles, "applicationRoles")
}

func (m *ManagementRoles) WriteToXml(writer *xmlWriter.Writer) {
	writer.WriteStartElement(xmlNamespace.Types, "ManagementRole")
	m.writeRolesToXml(writer, m.userRoles, "UserRoles")
	m.writeRolesToXml(writer, m.applicationRoles, "ApplicationRoles")
	writer.WriteEndElement()
}

func (m *ManagementRoles) writeRolesToXml(writer *xmlWriter.Writer, roles []string, elementName string) {
	if len(roles) == 0 {
		return
	}

	writer.WriteStartElement(xmlNamespace.Types, elementName)

	for _, role := range roles {
		writer.WriteElementValue(xmlNamespace.Types, "Role", role)
	}

	writer.WriteEndElement()
}
