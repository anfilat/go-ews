package ewsType

import (
	"github.com/anfilat/go-ews/internal/validate"
)

type AttendeeInfo struct {
	smtpAddress string
}

func NewAttendeeInfo(smtpAddress string) AttendeeInfo {
	return AttendeeInfo{
		smtpAddress: smtpAddress,
	}
}

func (a AttendeeInfo) Validate() error {
	return validate.Param(a.smtpAddress, "SmtpAddress")
}
