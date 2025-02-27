package receivers

import "github.com/kettek/rebui/events"

// PointerMove is used to receive pointer move events.
type PointerMove interface {
	HandlePointerMove(*events.PointerMove)
}

// PointerGlobalMove is used to receive global pointer move events. Note that these move events will only be received if the element received a press event (there does not have to be a handler).
type PointerGlobalMove interface {
	HandlePointerGlobalMove(*events.PointerMove)
}

// PointerIn is used to receive pointer in events.
type PointerIn interface {
	HandlePointerIn(*events.PointerIn)
}

// PointerOut is used to receive pointer out events.
type PointerOut interface {
	HandlePointerOut(*events.PointerOut)
}

// PointerPress is used to receive pointer press events. This occurs when a mouse button is pressed on the element.
type PointerPress interface {
	HandlePointerPress(*events.PointerPress)
}

// PointerRelease is used to receive pointer release events. This occurs when a mouse button is released on the element.
type PointerRelease interface {
	HandlePointerRelease(*events.PointerRelease)
}

// PointerGlobalRelease is used to receive global pointer release events. Note that these release events will only be received if the element received a press event (there does not have to be a handler).
type PointerGlobalRelease interface {
	HandlePointerGlobalRelease(*events.PointerRelease)
}

// PointerPressed is used to receive a press and subsequent release on the same element. This is akin to a "Click" event.
type PointerPressed interface {
	HandlePointerPressed(*events.PointerPressed)
}

// Focus is used to receive focus events. This occurs when a widget gains focus.
type Focus interface {
	HandleFocus(*events.Focus)
}

// Unfocus is used to receive unfocus events. This occurs when a widget loses focus.
type Unfocus interface {
	HandleUnfocus(*events.Unfocus)
}

// KeyPress is used to receive key press events. Only the focused element will receive this event.
type KeyPress interface {
	HandleKeyPress(*events.KeyPress)
}

// KeyRelease is used to receive key release events. Only the focused element will receive this event.
type KeyRelease interface {
	HandleKeyRelease(*events.KeyRelease)
}

// KeyInput is used to receive key input events. Only the focused element will receive this event.
type KeyInput interface {
	HandleKeyInput(*events.KeyInput)
}
