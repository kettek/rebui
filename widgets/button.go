package widgets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kettek/rebui"
)

type Button struct {
	Text
	backgroundColor color.Color
	borderColor     color.Color
}

func (b *Button) SetBackgroundColor(clr color.Color) {
	b.backgroundColor = clr
}

func (b *Button) SetBorderColor(clr color.Color) {
	b.borderColor = clr
}

func (b *Button) Draw(screen *ebiten.Image) {
	x := b.X
	y := b.Y

	vector.DrawFilledRect(screen, float32(x), float32(y), float32(b.Width), float32(b.Height), b.backgroundColor, true)
	vector.StrokeRect(screen, float32(x), float32(y), float32(b.Width), float32(b.Height), 1, b.borderColor, false)

	b.Text.Draw(screen)
}

func init() {
	rebui.RegisterElement("Button", &Button{})
}
