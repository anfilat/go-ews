package ews

import (
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/ewsCredentials"
	"github.com/anfilat/go-ews/ewsType"
	"github.com/anfilat/go-ews/internal/errors"
)

type Data struct {
	ImpersonatedUserId            *ewsType.ImpersonatedUserId
	PrivilegedUserId              *ewsType.PrivilegedUserId
	Exchange2007CompatibilityMode bool

	Version     exchangeVersion.Enum
	Credentials ewsCredentials.ExchangeCredentials
	Url         string
	Client      *Client
}

func (d *Data) EnsureClient() {
	if d.Client != nil {
		return
	}

	var opts []Option
	if d.Credentials != nil {
		opts = append(opts, WithCredentials(d.Credentials))
	}

	d.Client = NewClient(d.Url, opts...)
}

func (d *Data) Validate() error {
	if d.Url == "" {
		return errors.NewValidateError("the Url property on the ExchangeService object must be set")
	}

	if d.ImpersonatedUserId != nil {
		if err := d.ImpersonatedUserId.Validate(); err != nil {
			return err
		}
	}
	if d.PrivilegedUserId != nil && d.ImpersonatedUserId != nil {
		return errors.NewValidateError("can't set both impersonated user and privileged user in the ExchangeService object")
	}

	return nil
}

func (d *Data) GetRequestedServiceVersionString() string {
	if d.Exchange2007CompatibilityMode && d.Version == exchangeVersion.Exchange2007SP1 {
		return "Exchange2007"
	}
	return d.Version.String()
}
