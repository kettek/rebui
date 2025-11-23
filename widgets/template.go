package widgets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"
)

type Template struct {
	Basic
}

func (t *Template) IsTemplate() {
	// function to provide interface
}

func (t *Template) Draw(screen *ebiten.Image, sop *ebiten.DrawImageOptions) {
	// noopies
}

func init() {
	rebui.RegisterWidget("Template", &Template{})
}
