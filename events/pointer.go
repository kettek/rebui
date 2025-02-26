package events

// Pointer is an event that has pointer information.
type Pointer struct {
	X, Y   float64
	DX, DY float64
	// These values are X and Y values relative to the interior of the event receiver, if applicable.
	RelativeX, RelativeY float64
	ButtonID             int // The mouse button this represents if applicable.
	TouchID              int // The touch this represents if applicable.
}

// Cancelable is an event that can be canceled. This is the case for all events.
type Cancelable struct {
	canceled bool
}

// Cancel cancels the event.
func (c *Cancelable) Cancel() {
	c.canceled = true
}

// Canceled returns true if the event has been canceled.
func (c *Cancelable) Canceled() bool {
	return c.canceled
}

// PointerMove is an event that is triggered when a pointer moves with an element or if an element was pressed and the pointer moves outside of its hit box.
type PointerMove struct {
	Cancelable
	TargetWidget
	Timestamp
	Duration // How long this move event has been happening.
	Pointer
}

// PointerIn is an event that is triggered when a pointer enters an element.
type PointerIn struct {
	Cancelable
	TargetWidget
	Timestamp
	Pointer
}

// PointerOut is an event that is triggered when a pointer leaves an element.
type PointerOut struct {
	Cancelable
	TargetWidget
	Timestamp
	Pointer
}

// PointerPress is an event that is triggered when a pointer has depressed an element.
type PointerPress struct {
	Cancelable
	TargetWidget
	Timestamp
	Pointer
}

// PointerRelease is an event that is triggered when a pointer has released an element.
type PointerRelease struct {
	Cancelable
	TargetWidget
	Timestamp
	Duration // How long elapsed from press until release.
	Pointer
}

// PointerPressed is an event that is triggered when an element has received both a press and a release event.
type PointerPressed struct {
	Cancelable
	TargetWidget
	Timestamp
	Duration // How long elapsed from press until release.
	Pointer
}
