package ews

import (
	"fmt"
	"time"
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
		return fmt.Errorf("the time window's end time must be greater than its start time")
	}
	return nil
}

func (t TimeWindow) Duration() time.Duration {
	return t.endTime.Sub(t.startTime)
}
