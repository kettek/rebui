package blocks

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kettek/rebui"
)

// Config represents a configuration for blocks generation.
type Config struct {
	Face   text.Face
	Width  float64
	Height float64
	Wrap   rebui.Wrap
	VAlign rebui.Alignment
	HAlign rebui.Alignment
}
