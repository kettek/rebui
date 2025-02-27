package events

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Key is the base struct for KeyPress and KeyRelease
type Key struct {
	Key  ebiten.Key
	Name string
}

// KeyPress is an event that is triggered when a key is pressed.
type KeyPress struct {
	Cancelable
	TargetWidget // TargetWidget will be set to the current focused widget
	Timestamp
	Key    ebiten.Key
	Repeat int
}

// KeyRelease is an event that is triggered when a key is released.
type KeyRelease struct {
	Cancelable
	TargetWidget // TargetWidget will be set to the current focused widget
	Timestamp
	Duration // How long this move event has been happening.
	Key      ebiten.Key
}

// KeyInput is an event that is triggered when input runes are read from the keyboard.
type KeyInput struct {
	Cancelable
	TargetWidget // TargetWidget will be set to the current focused widget
	Timestamp
	Rune rune
}
