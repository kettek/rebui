package events

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Key struct {
	Key  ebiten.Key
	Rune rune // ???
}

type KeyPress struct {
	Cancelable
	TargetWidget // TargetWidget will be set to the current focused widget
	Key
	Timestamp
	Repeat int
}

type KeyRelease struct {
	Cancelable
	TargetWidget // TargetWidget will be set to the current focused widget
	Key
	Timestamp
}
