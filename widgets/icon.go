package widgets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"
)

type Icon struct {
	Button
	Image
}

func (b *Icon) Draw(screen *ebiten.Image, sop *ebiten.DrawImageOptions) {
	b.Button.Draw(screen, sop)
	b.Image.Draw(screen, sop)
}

func (b *Icon) AssignBorderColor(clr color.Color) {
	b.Button.borderColor = clr
	b.Image.borderColor = clr
}

func (b *Icon) AssignX(x float64) {
	b.Button.AssignX(x)
	b.Image.AssignX(x)
}

func (b *Icon) AssignY(y float64) {
	b.Button.AssignY(y)
	b.Image.AssignY(y)
}

func (b *Icon) AssignWidth(w float64) {
	b.Button.AssignWidth(w)
	b.Image.AssignWidth(w)
}

func (b *Icon) AssignHeight(h float64) {
	b.Button.AssignHeight(h)
	b.Image.AssignHeight(h)
}

func init() {
	rebui.RegisterWidget("Icon", &Icon{})
}
