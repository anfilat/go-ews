package ews

import (
	"github.com/anfilat/go-ews/internal"
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
	return internal.ValidateParam(a.smtpAddress, "SmtpAddress")
}
