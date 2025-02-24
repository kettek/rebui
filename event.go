package rebui

import "time"

type Event interface {
	Cancel()
	Canceled() bool
}

type TimestampEvent struct {
	Timestamp time.Time
}

type PointerEvent struct {
	X, Y      float64
	DX, DY    float64
	PointerID int
}

type CancelableEvent struct {
	canceled bool
}

func (c *CancelableEvent) Cancel() {
	c.canceled = true
}

func (c *CancelableEvent) Canceled() bool {
	return c.canceled
}

type pointerMoveEvent struct {
	TimestampEvent
	CancelableEvent
	PointerEvent
}

type PointerMoveEvent = *pointerMoveEvent

type pointerInEvent struct {
	TimestampEvent
	CancelableEvent
	PointerEvent
}

type PointerInEvent = *pointerInEvent

type pointerOutEvent struct {
	TimestampEvent
	CancelableEvent
	PointerEvent
}

type PointerOutEvent = *pointerOutEvent

type ClickEvent struct {
	TimestampEvent
	CancelableEvent
	PointerEvent
}

type pointerPressEvent struct {
	TimestampEvent
	CancelableEvent
	PointerEvent
}

type PointerPressEvent = *pointerPressEvent

type pointerReleaseEvent struct {
	TimestampEvent
	CancelableEvent
	PointerEvent
}

type PointerReleaseEvent = *pointerReleaseEvent

type pointerPressedEvent struct {
	TimestampEvent
	CancelableEvent
	PointerEvent
}

type PointerPressedEvent = *pointerPressedEvent
