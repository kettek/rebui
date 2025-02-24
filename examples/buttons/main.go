package main

import (
	"fmt"
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

type MyButton struct {
	widgets.Button
}

func (b *MyButton) HandlePointerIn(e rebui.PointerInEvent) {
	fmt.Println("PointerIn", e)
}

func (b *MyButton) HandlePointerOut(e rebui.PointerOutEvent) {
	fmt.Println("PointerOut", e)
}

func (b *MyButton) HandlePointerMove(e rebui.PointerMoveEvent) {
	fmt.Println("PointerMove", e)
}

func (b *MyButton) HandlePointerPress(e rebui.PointerPressEvent) {
	fmt.Println("PointerPress", e)
}

func (b *MyButton) HandlePointerRelease(e rebui.PointerReleaseEvent) {
	fmt.Println("Release", e)
}

func (b *MyButton) HandlePointerPressed(e rebui.PointerPressedEvent) {
	fmt.Println("Pressed", e)
}

func init() {
	rebui.RegisterElement("MyButton", &MyButton{})
}
