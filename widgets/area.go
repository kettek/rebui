package widgets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"
)

type Area struct {
	Basic
}

func (a *Area) Draw(screen *ebiten.Image, sop *ebiten.DrawImageOptions) {
	// NOP for now. We may want to allow areas to draw background/border...
}

func init() {
	rebui.RegisterWidget("Area", &Area{})
}
