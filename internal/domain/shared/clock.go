package shared

import "time"

type Clock interface {
	Now() time.Time
}

// System clock

type systemClock struct{}

func (systemClock) Now() time.Time {
	return time.Now()
}

// Global clock

var clock Clock = systemClock{}

func Now() time.Time {
	return clock.Now()
}
