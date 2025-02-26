package events

import "time"

// Timestamp is an event that has a timestamp.
type Timestamp struct {
	Timestamp time.Time
}

// Duration is an event that has a duration.
type Duration struct {
	Duration time.Duration
}
