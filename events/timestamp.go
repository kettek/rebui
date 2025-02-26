package events

import "time"

// Timestamp is an event that has a timestamp.
type Timestamp struct {
	Timestamp time.Time
}
