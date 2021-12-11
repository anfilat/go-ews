package ewsType

import (
	"time"

	"github.com/anfilat/go-ews/internal/base"
	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
	"github.com/anfilat/go-ews/internal/errors"
)

type TimeWindow struct {
	startTime time.Time
	endTime   time.Time
}

func NewTimeWindow(startTime, endTime time.Time) TimeWindow {
	return TimeWindow{
		startTime: startTime,
		endTime:   endTime,
	}
}

func (t TimeWindow) Validate() error {
	if t.startTime.After(t.endTime) || t.startTime.Equal(t.endTime) {
		return errors.NewValidateError("the time window's end time must be greater than its start time")
	}
	return nil
}

func (t TimeWindow) Duration() time.Duration {
	return t.endTime.Sub(t.startTime)
}

func (t TimeWindow) WriteToXmlUnscopedDatesOnly(writer base.Writer, xmlElementName string) {
	startTime := t.startTime.Format("2006-01-02T00:00:00")
	endTime := t.endTime.Format("2006-01-02T00:00:00")
	t.WriteToXml(writer, xmlElementName, startTime, endTime)
}

func (t TimeWindow) WriteToXml(writer base.Writer, xmlElementName string, startTime, endTime interface{}) {
	writer.WriteStartElement(xmlNamespace.Types, xmlElementName)

	writer.WriteElementValue(xmlNamespace.Types, "StartTime", startTime)
	writer.WriteElementValue(xmlNamespace.Types, "EndTime", endTime)

	writer.WriteEndElement()
}
