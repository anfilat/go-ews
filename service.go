package ews

import (
	"errors"

	"github.com/anfilat/go-ews/enumerations/availabilityData"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/ewsCredentials"
	"github.com/anfilat/go-ews/ewsType"
	"github.com/anfilat/go-ews/internal/validate"
)

var (
	ServiceUrlMustBeSet                        = errors.New("the Url property on the ExchangeService object must be set")
	CannotSetBothImpersonatedAndPrivilegedUser = errors.New("can't set both impersonated user and privileged user in the ExchangeService object")
)

type ExchangeService struct {
	version            exchangeVersion.Enum
	credentials        ewsCredentials.ExchangeCredentials
	url                string
	ImpersonatedUserId *ewsType.ImpersonatedUserId
	PrivilegedUserId   *ewsType.PrivilegedUserId
	client             *client
}

func NewExchangeService(version exchangeVersion.Enum) *ExchangeService {
	return &ExchangeService{
		version: version,
	}
}

func (e *ExchangeService) SetCredentials(credentials ewsCredentials.ExchangeCredentials) {
	e.client = nil
	e.credentials = credentials
}

func (e *ExchangeService) SetUrl(url string) {
	e.client = nil
	e.url = url
}

func (e *ExchangeService) ensureClient() {
	if e.client != nil {
		return
	}

	var opts []option
	if e.credentials != nil {
		opts = append(opts, withCredentials(e.credentials))
	}

	e.client = newClient(e.url, opts...)
}

func (e *ExchangeService) validate() error {
	if e.url == "" {
		return ServiceUrlMustBeSet
	}

	if e.PrivilegedUserId != nil && e.ImpersonatedUserId != nil {
		return CannotSetBothImpersonatedAndPrivilegedUser
	}

	return nil
}

func (e *ExchangeService) GetUserAvailability(
	attendees []ewsType.AttendeeInfo,
	timeWindow ewsType.TimeWindow,
	requestedData availabilityData.Enum,
	options *ewsType.AvailabilityOptions,
) (*GetUserAvailabilityResults, error) {
	if err := validate.ParamSlice(attendees, "attendees"); err != nil {
		return nil, err
	}
	if err := validate.Param(timeWindow, "timeWindow"); err != nil {
		return nil, err
	}
	if err := validate.Param(options, "options"); err != nil {
		return nil, err
	}

	e.ensureClient()
	return NewGetUserAvailabilityRequest(attendees, timeWindow, requestedData, options).Execute(e)
}
