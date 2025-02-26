package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"
	"github.com/kettek/rebui/widgets"
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
	return 320, 240
}

func main() {
	g := &Game{}

	g.layout.AddNode(rebui.Node{
		Type:    "Button",
		ID:      "button1",
		Width:   "50%",
		Height:  "50",
		X:       "50%",
		Y:       "50",
		OriginX: "-50%",
		OriginY: "-50%",
	})

	node := g.layout.GetByID("button1")

	node.OnPointerIn = func(e rebui.EventPointerIn) {
		e.Widget.(*widgets.Button).SetBackgroundColor(color.NRGBA{0, 255, 0, 255})
	}

	node.OnPointerOut = func(e rebui.EventPointerOut) {
		e.Widget.(*widgets.Button).SetBackgroundColor(color.NRGBA{255, 0, 0, 255})
	}
	node.OnPointerPress = func(e rebui.EventPointerPress) {
		e.Widget.(*widgets.Button).SetBackgroundColor(color.NRGBA{0, 0, 255, 255})
	}
	node.OnPointerRelease = func(e rebui.EventPointerRelease) {
		e.Widget.(*widgets.Button).SetBackgroundColor(color.NRGBA{0, 255, 0, 255})
	}

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Layout (Ebiten Demo)")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
