package main

import (
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
	ebiten.SetWindowTitle("Lifetime")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

type MyButton struct {
	widgets.Button
}

func (b *MyButton) HandleGenerate() {
	log.Println("generated")
}

func init() {
	rebui.RegisterWidget("MyButton", &MyButton{})
}
