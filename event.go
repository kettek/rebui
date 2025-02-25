package rebui

import "time"

// Event is our most basic event interface.
type Event interface {
	Cancel()
	Canceled() bool
}

// TimestampEvent is an event that has a timestamp.
type TimestampEvent struct {
	Timestamp time.Time
}

type TargetElementEvent struct {
	Target Element
}

// PointerEvent is an event that has pointer information.
type PointerEvent struct {
	X, Y      float64
	DX, DY    float64
	PointerID int
}

// CancelableEvent is an event that can be canceled. This is the case for all events.
type CancelableEvent struct {
	canceled bool
}

// Cancel cancels the event.
func (c *CancelableEvent) Cancel() {
	c.canceled = true
}

// Canceled returns true if the event has been canceled.
func (c *CancelableEvent) Canceled() bool {
	return c.canceled
}

type pointerMoveEvent struct {
	TargetElementEvent
	TimestampEvent
	CancelableEvent
	PointerEvent
}

// PointerMoveEvent is an event that is triggered when a pointer moves with an element or if an element was pressed and the pointer moves outside of its hit box.
type PointerMoveEvent = *pointerMoveEvent

type pointerInEvent struct {
	TargetElementEvent
	TimestampEvent
	CancelableEvent
	PointerEvent
}

// PointerInEvent is an event that is triggered when a pointer enters an element.
type PointerInEvent = *pointerInEvent

type pointerOutEvent struct {
	TargetElementEvent
	TimestampEvent
	CancelableEvent
	PointerEvent
}

// PointerOutEvent is an event that is triggered when a pointer leaves an element.
type PointerOutEvent = *pointerOutEvent

type pointerPressEvent struct {
	TargetElementEvent
	TimestampEvent
	CancelableEvent
	PointerEvent
}

// PointerPressEvent is an event that is triggered when a pointer has depressed an element.
type PointerPressEvent = *pointerPressEvent

type pointerReleaseEvent struct {
	TargetElementEvent
	TimestampEvent
	CancelableEvent
	PointerEvent
}

// PointerReleaseEvent is an event that is triggered when a pointer has released an element.
type PointerReleaseEvent = *pointerReleaseEvent

type pointerPressedEvent struct {
	TargetElementEvent
	TimestampEvent
	CancelableEvent
	PointerEvent
}

// PointerPressedEvent is an event that is triggered when an elemenbt has received both a press and a release event.
type PointerPressedEvent = *pointerPressedEvent
