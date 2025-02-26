package events

import "github.com/hajimehoshi/ebiten/v2"

// Element is the interface that all widgets must implement.
type Element interface {
	Draw(*ebiten.Image)
}

// TargetElement is an event that has a target element.
type TargetElement struct {
	Target Element
}
