package rebui

import (
	"github.com/kettek/rebui/events"
)

// Event is our most basic event interface.
type Event interface {
}

// EventCancelable is an event that is cancelable. This applies to all events.
type EventCancelable interface {
	Cancel()
	Canceled() bool
}

// EventPointerMove is an event that is triggered when a pointer within an element or when an element is pressed and the pointer moves anywhere.
type EventPointerMove = *events.PointerMove

// EventPointerIn is an event that is triggered when a pointer enters an element.
type EventPointerIn = *events.PointerIn

// EventPointerOut is an event that is triggered when a pointer leaves an element.
type EventPointerOut = *events.PointerOut

// EventPointerPress is an event that is triggered when a pointer has depressed an element.
type EventPointerPress = *events.PointerPress

// EventPointerRelease is an event that is triggered when a pointer has released an element.
type EventPointerRelease = *events.PointerRelease

// EventPointerPressed is an event that is triggered when a pointer has depressed an element.
type EventPointerPressed = *events.PointerPressed

// EventFocus is triggered when a widget gains focus.
type EventFocus = *events.Focus

// EventUnfocus is triggered when a widget has lost focus.
type EventUnfocus = *events.Unfocus

// EventKeyPress is used to receive key press events. Only the focused element will receive this event.
type EventKeyPress = *events.KeyPress

// EventKeyRelease is used to receive key release events. Only the focused element will receive this event.
type EventKeyRelease = *events.KeyRelease

// EventKeyInput is used to receive key input events. Only the focused element will receive this event.
type EventKeyInput = *events.KeyInput
