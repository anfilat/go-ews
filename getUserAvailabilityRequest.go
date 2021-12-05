package ews

import (
	"github.com/anfilat/go-ews/enumerations/availabilityData"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
)

type getUserAvailabilityRequest struct {
	attendees     []AttendeeInfo
	timeWindow    TimeWindow
	requestedData availabilityData.Enum
	options       *AvailabilityOptions
}

func NewGetUserAvailabilityRequest(
	attendees []AttendeeInfo,
	timeWindow TimeWindow,
	requestedData availabilityData.Enum,
	options *AvailabilityOptions,
) *getUserAvailabilityRequest {
	return &getUserAvailabilityRequest{
		attendees:     attendees,
		timeWindow:    timeWindow,
		requestedData: requestedData,
		options:       options,
	}
}

func (r *getUserAvailabilityRequest) Execute(service *ExchangeService) (*GetUserAvailabilityResults, error) {
	if err := service.validate(); err != nil {
		return nil, err
	}
	if err := r.Validate(); err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *getUserAvailabilityRequest) IsFreeBusyViewRequested() bool {
	return r.requestedData == availabilityData.FreeBusy || r.requestedData == availabilityData.FreeBusyAndSuggestions
}

func (r *getUserAvailabilityRequest) IsSuggestionsViewRequested() bool {
	return r.requestedData == availabilityData.Suggestions || r.requestedData == availabilityData.FreeBusyAndSuggestions
}

func (r *getUserAvailabilityRequest) EmitTimeZoneHeader() bool {
	return true
}

func (r *getUserAvailabilityRequest) GetMinimumRequiredServerVersion() exchangeVersion.Enum {
	return exchangeVersion.Exchange2007SP1
}

func (r *getUserAvailabilityRequest) GetResponseXmlElementName() string {
	return "GetUserAvailabilityResponse"
}

func (r *getUserAvailabilityRequest) GetXmlElementName() string {
	return "GetUserAvailabilityRequest"
}

func (r *getUserAvailabilityRequest) Validate() error {
	if err := r.options.Validate(r.timeWindow.Duration()); err != nil {
		return err
	}
	return nil
}
