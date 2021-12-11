package ewsType

import (
	"github.com/anfilat/go-ews/enumerations/meetingAttendeeType"
	"github.com/anfilat/go-ews/internal/base"
	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
	"github.com/anfilat/go-ews/internal/validate"
)

type AttendeeInfo struct {
	SmtpAddress      string
	AttendeeType     meetingAttendeeType.Enum
	ExcludeConflicts bool
}

func NewAttendeeInfo(smtpAddress string) AttendeeInfo {
	return AttendeeInfo{
		SmtpAddress:      smtpAddress,
		AttendeeType:     meetingAttendeeType.Required,
		ExcludeConflicts: false,
	}
}

func (a AttendeeInfo) Validate() error {
	return validate.Param(a.SmtpAddress, "SmtpAddress")
}

func (a *AttendeeInfo) WriteXml(writer base.Writer) {
	writer.WriteStartElement(xmlNamespace.Types, "MailboxData")

	writer.WriteStartElement(xmlNamespace.Types, "Email")
	writer.WriteElementValue(xmlNamespace.Types, "Address", a.SmtpAddress)
	writer.WriteEndElement()

	writer.WriteElementValue(xmlNamespace.Types, "AttendeeType", a.AttendeeType)

	writer.WriteElementValue(xmlNamespace.Types, "ExcludeConflicts", a.ExcludeConflicts)

	writer.WriteEndElement()
}
