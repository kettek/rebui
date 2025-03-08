package widgets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kettek/rebui"
)

type Button struct {
	Label
	backgroundColor color.Color
	borderColor     color.Color
}

func (b *Button) AssignBackgroundColor(clr color.Color) {
	b.backgroundColor = clr
}

func (b *Button) AssignBorderColor(clr color.Color) {
	b.borderColor = clr
}

func (b *Button) Draw(screen *ebiten.Image, sop *ebiten.DrawImageOptions) {
	x := sop.GeoM.Element(0, 2)
	y := sop.GeoM.Element(1, 2)

	vector.DrawFilledRect(screen, float32(x), float32(y), float32(b.Width), float32(b.Height), b.backgroundColor, true)
	vector.StrokeRect(screen, float32(x), float32(y), float32(b.Width), float32(b.Height), 1, b.borderColor, false)

	b.Label.Draw(screen, sop)
}

func init() {
	rebui.RegisterWidget("Button", &Button{})
}
