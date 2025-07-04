package widgets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Border struct {
	borderWidth float64
	borderColor color.Color
}

// AssignBorderWidth sets the width of the element.
func (b *Border) AssignBorderWidth(w float64) {
	b.borderWidth = w
}

// GetBorderWidth returns the width of the element.
func (b *Border) GetBorderWidth() float64 {
	return b.borderWidth
}

// AssignBorderColor sets the color of the element.
func (b *Border) AssignBorderColor(clr color.Color) {
	b.borderColor = clr
}

// GetBorderColor returns the color of the element.
func (b *Border) GetBorderColor() color.Color {
	return b.borderColor
}

func (b *Border) drawBorder(screen *ebiten.Image, x, y, width, height float32) {
	if b.borderColor != nil {
		_, _, _, a := b.borderColor.RGBA()
		if a > 0 {
			vector.StrokeRect(screen, x, y, width, height, float32(b.borderWidth), b.borderColor, false)
		}
	}
}
