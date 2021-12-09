package ews

import (
	"context"

	"github.com/anfilat/go-ews/enumerations/availabilityData"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/ewsCredentials"
	"github.com/anfilat/go-ews/ewsType"
	"github.com/anfilat/go-ews/internal/ews"
	"github.com/anfilat/go-ews/internal/requests"
	"github.com/anfilat/go-ews/internal/validate"
)

type ExchangeService struct {
	data *ews.Data
}

func New(version exchangeVersion.Enum) *ExchangeService {
	return &ExchangeService{
		data: &ews.Data{
			Version: version,
		},
	}
}

func (e *ExchangeService) SetCredentials(value ewsCredentials.ExchangeCredentials) {
	e.data.Client = nil
	e.data.Credentials = value
}

func (e *ExchangeService) SetUrl(value string) {
	e.data.Client = nil
	e.data.Url = value
}

func (e *ExchangeService) SetImpersonatedUserId(value *ewsType.ImpersonatedUserId) {
	e.data.ImpersonatedUserId = value
}

func (e *ExchangeService) SetPrivilegedUserId(value *ewsType.PrivilegedUserId) {
	e.data.PrivilegedUserId = value
}

func (e *ExchangeService) SetExchange2007CompatibilityMode(value bool) {
	e.data.Exchange2007CompatibilityMode = value
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
	_, err := execute(ctx, e.data, request)
	// return result.(*ewsType.GetUserAvailabilityResults), err
	return nil, err
}
