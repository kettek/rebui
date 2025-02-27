package events

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Key is the base struct for KeyPress and KeyRelease
type Key struct {
	Key  ebiten.Key
	Rune rune // ???
}

// KeyPress is an event that is triggered when a key is pressed.
type KeyPress struct {
	Cancelable
	TargetWidget // TargetWidget will be set to the current focused widget
	Key
	Timestamp
	Repeat int
}

// KeyRelease is an event that is triggered when a key is released.
type KeyRelease struct {
	Cancelable
	TargetWidget // TargetWidget will be set to the current focused widget
	Key
	Timestamp
}
