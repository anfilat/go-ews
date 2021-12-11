package ewsType

import (
	"fmt"
	"time"

	"github.com/anfilat/go-ews/enumerations/freeBusyViewType"
	"github.com/anfilat/go-ews/enumerations/suggestionQuality"
	"github.com/anfilat/go-ews/internal/base"
	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
	"github.com/anfilat/go-ews/internal/errors"
	"github.com/anfilat/go-ews/internal/validate"
)

type AvailabilityOptions struct {
	mergedFreeBusyInterval               int
	requestedFreeBusyView                freeBusyViewType.Enum
	goodSuggestionThreshold              int
	maximumSuggestionsPerDay             int
	maximumNonWorkHoursSuggestionsPerDay int
	meetingDuration                      int
	minimumSuggestionQuality             suggestionQuality.Enum
	detailedSuggestionsWindow            *TimeWindow
	currentMeetingTime                   *time.Time
	globalObjectId                       *string
}

func NewAvailabilityOptions() *AvailabilityOptions {
	return &AvailabilityOptions{
		mergedFreeBusyInterval:               30,
		requestedFreeBusyView:                freeBusyViewType.Detailed,
		goodSuggestionThreshold:              25,
		maximumSuggestionsPerDay:             10,
		maximumNonWorkHoursSuggestionsPerDay: 0,
		meetingDuration:                      60,
		minimumSuggestionQuality:             suggestionQuality.Fair,
		detailedSuggestionsWindow:            nil,
		currentMeetingTime:                   nil,
		globalObjectId:                       nil,
	}
}

func (a *AvailabilityOptions) WithMergedFreeBusyInterval(value int) *AvailabilityOptions {
	a.mergedFreeBusyInterval = value
	return a
}

func (a *AvailabilityOptions) WithRequestedFreeBusyView(value freeBusyViewType.Enum) *AvailabilityOptions {
	a.requestedFreeBusyView = value
	return a
}

func (a *AvailabilityOptions) WithGoodSuggestionThreshold(value int) *AvailabilityOptions {
	a.goodSuggestionThreshold = value
	return a
}

func (a *AvailabilityOptions) WithMaximumSuggestionsPerDay(value int) *AvailabilityOptions {
	a.maximumSuggestionsPerDay = value
	return a
}

func (a *AvailabilityOptions) WithMaximumNonWorkHoursSuggestionsPerDay(value int) *AvailabilityOptions {
	a.maximumNonWorkHoursSuggestionsPerDay = value
	return a
}

func (a *AvailabilityOptions) WithMeetingDuration(value int) *AvailabilityOptions {
	a.meetingDuration = value
	return a
}

func (a *AvailabilityOptions) WithMinimumSuggestionQuality(value suggestionQuality.Enum) *AvailabilityOptions {
	a.minimumSuggestionQuality = value
	return a
}

func (a *AvailabilityOptions) WithDetailedSuggestionsWindow(value *TimeWindow) *AvailabilityOptions {
	a.detailedSuggestionsWindow = value
	return a
}

func (a *AvailabilityOptions) WithCurrentMeetingTime(value *time.Time) *AvailabilityOptions {
	a.currentMeetingTime = value
	return a
}

func (a *AvailabilityOptions) WithGlobalObjectId(value string) *AvailabilityOptions {
	a.globalObjectId = &value
	return a
}

func (a *AvailabilityOptions) Validate() error {
	if a.mergedFreeBusyInterval < 5 || a.mergedFreeBusyInterval > 1440 {
		return errors.NewValidateError(fmt.Sprintf("mergedFreeBusyInterval must be between %v and %v", 5, 1440))
	}
	if a.goodSuggestionThreshold < 1 || a.goodSuggestionThreshold > 49 {
		return errors.NewValidateError(fmt.Sprintf("goodSuggestionThreshold must be between %v and %v", 1, 49))
	}
	if a.maximumSuggestionsPerDay < 0 || a.maximumSuggestionsPerDay > 48 {
		return errors.NewValidateError(fmt.Sprintf("maximumSuggestionsPerDay must be between %v and %v", 0, 48))
	}
	if a.maximumNonWorkHoursSuggestionsPerDay < 0 || a.maximumNonWorkHoursSuggestionsPerDay > 48 {
		return errors.NewValidateError(fmt.Sprintf("maximumNonWorkHoursSuggestionsPerDay must be between %v and %v", 0, 48))
	}
	if a.meetingDuration < 30 || a.meetingDuration > 1440 {
		return errors.NewValidateError(fmt.Sprintf("meetingDuration must be between %v and %v", 30, 1440))
	}
	return validate.ParamAllowNull(a.detailedSuggestionsWindow, "DetailedSuggestionsWindow")
}

func (a *AvailabilityOptions) ValidateTimeWindow(timeWindow *TimeWindow) error {
	if time.Duration(a.mergedFreeBusyInterval)*time.Minute > timeWindow.Duration() {
		return fmt.Errorf("mergedFreeBusyInterval must be smaller than the specified time window")
	}
	return nil
}

func (a *AvailabilityOptions) WriteToXml(writer base.Writer, request interface{}) {
	timeWindow := request.(interface{ GetTimeWindow() *TimeWindow }).GetTimeWindow()

	if request.(interface{ IsFreeBusyViewRequested() bool }).IsFreeBusyViewRequested() {
		writer.WriteStartElement(xmlNamespace.Types, "FreeBusyViewOptions")

		timeWindow.WriteToXmlUnscopedDatesOnly(writer, "TimeWindow")

		writer.WriteElementValue(xmlNamespace.Types, "MergedFreeBusyIntervalInMinutes", a.mergedFreeBusyInterval)
		writer.WriteElementValue(xmlNamespace.Types, "RequestedView", a.requestedFreeBusyView)

		writer.WriteEndElement()
	}

	if request.(interface{ IsSuggestionsViewRequested() bool }).IsSuggestionsViewRequested() {
		writer.WriteStartElement(xmlNamespace.Types, "SuggestionsViewOptions")

		writer.WriteElementValue(xmlNamespace.Types, "GoodThreshold", a.goodSuggestionThreshold)
		writer.WriteElementValue(xmlNamespace.Types, "MaximumResultsByDay", a.maximumSuggestionsPerDay)
		writer.WriteElementValue(xmlNamespace.Types, "MaximumNonWorkHourResultsByDay", a.maximumNonWorkHoursSuggestionsPerDay)
		writer.WriteElementValue(xmlNamespace.Types, "MeetingDurationInMinutes", a.meetingDuration)
		writer.WriteElementValue(xmlNamespace.Types, "MinimumSuggestionQuality", a.minimumSuggestionQuality)

		timeWindowToSerialize := timeWindow
		if a.detailedSuggestionsWindow != nil {
			timeWindowToSerialize = a.detailedSuggestionsWindow
		}

		timeWindowToSerialize.WriteToXmlUnscopedDatesOnly(writer, "DetailedSuggestionsWindow")

		if a.currentMeetingTime != nil {
			writer.WriteElementValue(xmlNamespace.Types, "CurrentMeetingTime", a.currentMeetingTime)
		}

		writer.WriteElementValue(xmlNamespace.Types, "GlobalObjectId", a.globalObjectId)

		writer.WriteEndElement()
	}
}
