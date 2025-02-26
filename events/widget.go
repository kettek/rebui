package events

import "github.com/hajimehoshi/ebiten/v2"

// Widget is the interface that all widgets must implement.
type Widget interface {
	Draw(*ebiten.Image)
}

// TargetWidget is an event that has a target element.
type TargetWidget struct {
	Widget Widget
}
