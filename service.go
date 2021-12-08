package ews

import (
	"context"

	"github.com/anfilat/go-ews/enumerations/availabilityData"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/ewsCredentials"
	"github.com/anfilat/go-ews/ewsType"
	"github.com/anfilat/go-ews/internal/errors"
	"github.com/anfilat/go-ews/internal/requests"
	"github.com/anfilat/go-ews/internal/validate"
)

type ExchangeService struct {
	ImpersonatedUserId            *ewsType.ImpersonatedUserId
	PrivilegedUserId              *ewsType.PrivilegedUserId
	Exchange2007CompatibilityMode bool

	version     exchangeVersion.Enum
	credentials ewsCredentials.ExchangeCredentials
	url         string
	client      *client
}

func New(version exchangeVersion.Enum) *ExchangeService {
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
		return errors.NewValidateError("the Url property on the ExchangeService object must be set")
	}

	if e.PrivilegedUserId != nil && e.ImpersonatedUserId != nil {
		return errors.NewValidateError("can't set both impersonated user and privileged user in the ExchangeService object")
	}

	return nil
}

func (e *ExchangeService) getRequestedServiceVersionString() string {
	if e.Exchange2007CompatibilityMode && e.version == exchangeVersion.Exchange2007SP1 {
		return "Exchange2007"
	}
	return e.version.String()
}

func (e *ExchangeService) GetUserAvailability(
	ctx context.Context,
	attendees []ewsType.AttendeeInfo,
	timeWindow ewsType.TimeWindow,
	requestedData availabilityData.Enum,
	options *ewsType.AvailabilityOptions,
) (*ewsType.GetUserAvailabilityResults, error) {
	if err := validate.ParamSlice(attendees, "attendees"); err != nil {
		return nil, err
	}
	if err := validate.Param(timeWindow, "timeWindow"); err != nil {
		return nil, err
	}
	if err := validate.Param(options, "options"); err != nil {
		return nil, err
	}

	request := requests.NewGetUserAvailabilityRequest(attendees, timeWindow, requestedData, options)
	_, err := execute(ctx, e, request)
	// return result.(*ewsType.GetUserAvailabilityResults), err
	return nil, err
}
