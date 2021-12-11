package requests

import (
	"github.com/anfilat/go-ews/enumerations/availabilityData"
	"github.com/anfilat/go-ews/enumerations/exchangeVersion"
	"github.com/anfilat/go-ews/ewsType"
	"github.com/anfilat/go-ews/internal/base"
)

type getUserAvailabilityRequest struct {
	attendees     []ewsType.AttendeeInfo
	timeWindow    *ewsType.TimeWindow
	requestedData availabilityData.Enum
	options       *ewsType.AvailabilityOptions
}

func NewGetUserAvailabilityRequest(
	attendees []ewsType.AttendeeInfo,
	timeWindow ewsType.TimeWindow,
	requestedData availabilityData.Enum,
	options *ewsType.AvailabilityOptions,
) base.Request {
	return &getUserAvailabilityRequest{
		attendees:     attendees,
		timeWindow:    &timeWindow,
		requestedData: requestedData,
		options:       options,
	}
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
	if err := r.timeWindow.Validate(); err != nil {
		return err
	}
	if err := r.options.ValidateTimeWindow(r.timeWindow); err != nil {
		return err
	}
	return nil
}
