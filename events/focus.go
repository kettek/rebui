package events

// Focus is an event that is triggered when a widget has been focused.
type Focus struct {
	Cancelable
	TargetWidget
	Timestamp
}

// Unfocus is triggered when a widget has lost focus.
type Unfocus struct {
	Cancelable
	TargetWidget
	Timestamp
}
