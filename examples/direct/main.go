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
		Type:            "MyButton",
		Width:           "50%",
		Height:          "50",
		X:               "50%",
		Y:               "50",
		OriginX:         "-50%",
		OriginY:         "-50%",
		BackgroundColor: "red",
	})

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Layout (Ebiten Demo)")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

type MyButton struct {
	widgets.Button
}

func (b *MyButton) HandlePointerIn(e rebui.EventPointerIn) {
	b.SetBackgroundColor(color.NRGBA{0, 255, 0, 255})
}

func (b *MyButton) HandlePointerOut(e rebui.EventPointerOut) {
	b.SetBackgroundColor(color.NRGBA{255, 0, 0, 255})
}

func (b *MyButton) HandlePointerPress(e rebui.EventPointerPress) {
	b.SetBackgroundColor(color.NRGBA{0, 0, 255, 255})
}

func (b *MyButton) HandlePointerRelease(e rebui.EventPointerRelease) {
	b.SetBackgroundColor(color.NRGBA{0, 255, 0, 255})
}

func init() {
	rebui.RegisterWidget("MyButton", &MyButton{})
}
