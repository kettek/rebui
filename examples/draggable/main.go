package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"
	"github.com/kettek/rebui/widgets"
)

type Game struct {
	layout *rebui.Layout
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

	bytes, _ := os.ReadFile("buttons.json")

	layout, err := rebui.NewLayout(string(bytes))
	if err != nil {
		log.Fatal(err)
	}

	g.layout = layout
	g.layout.Generate(320, 240)

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Layout (Ebiten Demo)")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

type DraggableButton struct {
	widgets.Button
}

func (b *DraggableButton) HandlePointerGlobalMove(e rebui.PointerMoveEvent) {
	// Ignore default move event.
	if e.PointerID == -1 {
		return
	}
	b.X += e.DX
	b.Y += e.DY
}

func init() {
	rebui.RegisterElement("DraggableButton", &DraggableButton{})
}
