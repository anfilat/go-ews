package ewsType

import (
	"fmt"
	"time"
)

type TimeSpan time.Duration

func NewTimeSpan(d time.Duration) TimeSpan {
	return TimeSpan(d)
}

func (d TimeSpan) ToXSDuration() string {
	offsetStr := ""
	if d < 0 {
		offsetStr = "-"
		d = -d
	}
	return fmt.Sprintf("%sP%dDT%dH%dM%d.%dS", offsetStr, d.Days(), d.Hours(), d.Minutes(), d.Seconds(), d.Milliseconds())
}

func (d TimeSpan) TotalSeconds() int64 {
	return time.Duration(d).Nanoseconds() / int64(time.Second)
}

func (d TimeSpan) Milliseconds() int {
	return int(time.Duration(d).Nanoseconds() % int64(time.Second) / 1e6)
}

func (d TimeSpan) Seconds() int {
	return int(time.Duration(d).Nanoseconds() % int64(time.Minute) / 1e9)
}

func (d TimeSpan) Minutes() int {
	return int(time.Duration(d).Nanoseconds() % int64(time.Hour) / 60e9)
}

func (d TimeSpan) Hours() int {
	return int(time.Duration(d).Nanoseconds() % int64(24*time.Hour) / 3600e9)
}

func (d TimeSpan) Days() int {
	return int(time.Duration(d).Nanoseconds() / int64(24*time.Hour))
}
