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

func (b *Icon) SetBorderColor(clr color.Color) {
	b.Button.borderColor = clr
	b.Image.borderColor = clr
}

func (b *Icon) SetX(x float64) {
	b.Button.SetX(x)
	b.Image.SetX(x)
}

func (b *Icon) SetY(y float64) {
	b.Button.SetY(y)
	b.Image.SetY(y)
}

func (b *Icon) SetWidth(w float64) {
	b.Button.SetWidth(w)
	b.Image.SetWidth(w)
}

func (b *Icon) SetHeight(h float64) {
	b.Button.SetHeight(h)
	b.Image.SetHeight(h)
}

func init() {
	rebui.RegisterWidget("Icon", &Icon{})
}
