package ews

import (
	"github.com/anfilat/go-ews/enumerations/dateTimePrecision"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/ewsType"
	"github.com/anfilat/go-ews/internal/base"
	"github.com/anfilat/go-ews/internal/errors"
)

type ServiceData struct {
	Url                           string
	Credentials                   base.ExchangeCredentials
	Version                       exchangeVersion.Enum
	ImpersonatedUserId            *ewsType.ImpersonatedUserId
	PrivilegedUserId              *ewsType.PrivilegedUserId
	ManagementRoles               *ewsType.ManagementRoles
	Exchange2007CompatibilityMode bool
	DateTimePrecision             dateTimePrecision.Enum

	Client *Client
}

func (d *ServiceData) EnsureClient() {
	if d.Client != nil {
		return
	}

	var opts []Option
	if d.Credentials != nil {
		opts = append(opts, WithCredentials(d.Credentials))
	}

	d.Client = NewClient(d.Url, opts...)
}

func (d *ServiceData) Validate() error {
	if d.Url == "" {
		return errors.NewValidateError("the Url property on the ExchangeService object must be set")
	}

	if d.PrivilegedUserId != nil {
		if err := d.PrivilegedUserId.Validate(); err != nil {
			return err
		}
	}
	if d.ImpersonatedUserId != nil {
		if err := d.ImpersonatedUserId.Validate(); err != nil {
			return err
		}
	}
	if d.ManagementRoles != nil {
		if err := d.ManagementRoles.Validate(); err != nil {
			return err
		}
	}
	if d.PrivilegedUserId != nil && d.ImpersonatedUserId != nil {
		return errors.NewValidateError("can't set both impersonated user and privileged user in the ExchangeService object")
	}

	return nil
}

func (d *ServiceData) GetRequestedServiceVersionString() string {
	if d.Exchange2007CompatibilityMode && d.Version == exchangeVersion.Exchange2007SP1 {
		return "Exchange2007"
	}
	return d.Version.String()
}

func (d *ServiceData) GetVersion() exchangeVersion.Enum {
	return d.Version
}
