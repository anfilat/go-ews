package timeZones

import (
	"strings"
	"time"

	"github.com/anfilat/go-ews/ewsProperty"
	"github.com/anfilat/go-ews/ewsType"
	"github.com/anfilat/go-ews/internal/base"
)

const (
	StandardPeriodId   = "Std"
	StandardPeriodName = "Standard"
	DaylightPeriodId   = "Dlt"
	DaylightPeriodName = "Daylight"
)

type TimeZonePeriod struct {
	name string
	id   string
	bias ewsType.TimeSpan
}

func NewTimeZonePeriod(name, id string, bias time.Duration) TimeZonePeriod {
	return TimeZonePeriod{name, id, ewsType.TimeSpan(bias)}
}

func (t TimeZonePeriod) WriteToXml(writer base.Writer) {
	ewsProperty.WriteToXml(t, writer, "Period")
}

func (t TimeZonePeriod) WriteAttributesToXml(writer base.Writer) {
	writer.WriteAttributeValue("", "Bias", t.bias.ToXSDuration())
	writer.WriteAttributeValue("", "Name", t.name)
	writer.WriteAttributeValue("", "Id", t.id)
}

func (t TimeZonePeriod) IsStandardPeriod() bool {
	return strings.EqualFold(t.name, StandardPeriodName)
}
