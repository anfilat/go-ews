package ews

import (
	"errors"

	"github.com/anfilat/go-ews/enumerations/availabilityData"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/internal"
)

var (
	ServiceUrlMustBeSet                        = errors.New("the Url property on the ExchangeService object must be set")
	CannotSetBothImpersonatedAndPrivilegedUser = errors.New("can't set both impersonated user and privileged user in the ExchangeService object")
)

type ExchangeService struct {
	version            exchangeVersion.Enum
	credentials        ExchangeCredentials
	url                string
	ImpersonatedUserId *ImpersonatedUserId
	PrivilegedUserId   *PrivilegedUserId
	client             *client
}

func NewExchangeService(version exchangeVersion.Enum) *ExchangeService {
	return &ExchangeService{
		version: version,
	}
}

func (e *ExchangeService) SetCredentials(credentials ExchangeCredentials) {
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
	attendees []AttendeeInfo,
	timeWindow TimeWindow,
	requestedData availabilityData.Enum,
	options *AvailabilityOptions,
) (*GetUserAvailabilityResults, error) {
	if err := internal.ValidateParamSlice(attendees, "attendees"); err != nil {
		return nil, err
	}
	if err := internal.ValidateParam(timeWindow, "timeWindow"); err != nil {
		return nil, err
	}
	if err := internal.ValidateParam(options, "options"); err != nil {
		return nil, err
	}

	e.ensureClient()
	return NewGetUserAvailabilityRequest(attendees, timeWindow, requestedData, options).Execute(e)
}
