package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"
	_ "github.com/kettek/rebui/defaults/font"
	_ "github.com/kettek/rebui/widgets"
)

type Game struct {
	layout rebui.Layout
}

func (g *Game) Update() error {
	g.layout.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.layout.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{}

	g.layout.AddNode(rebui.Node{
		Type:            "Text",
		Width:           "50%",
		Height:          "50%",
		X:               "50%",
		Y:               "50%",
		OriginX:         "-50%",
		OriginY:         "-50%",
		BorderColor:     "white",
		BackgroundColor: "red",
		Text:            "This is some text! Wowwwwwwwwwwwwwwwwwwww, and it should have word wrap too, I think!, MAYBE!!!\nor maybe not?\nit does!!!",
		TextWrap:        rebui.WrapWord,
		HorizontalAlign: rebui.AlignCenter,
		VerticalAlign:   rebui.AlignMiddle,
	})

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Layout (Ebiten Demo)")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
