package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"
	"github.com/kettek/rebui/widgets"

	// This import sets the default ui font
	_ "github.com/kettek/rebui/defaults/font"
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
	g.layout.Generate()

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
	b.Button.HandlePointerIn(e) // For default hover behavior
	fmt.Println("PointerIn", e)
}

func (b *MyButton) HandlePointerOut(e rebui.EventPointerOut) {
	b.Button.HandlePointerOut(e) // For default hover behavior
	fmt.Println("PointerOut", e)
}

func (b *MyButton) HandlePointerMove(e rebui.EventPointerMove) {
	fmt.Println("PointerMove", e)
}

func (b *MyButton) HandlePointerPress(e rebui.EventPointerPress) {
	b.Button.HandlePointerPress(e)
	fmt.Println("PointerPress", e)
}

func (b *MyButton) HandlePointerRelease(e rebui.EventPointerRelease) {
	b.Button.HandlePointerRelease(e)
	fmt.Println("Release", e)
}

func (b *MyButton) HandlePointerPressed(e rebui.EventPointerPressed) {
	b.Button.HandlePointerPressed(e)
	fmt.Println("Pressed", e)
}

func init() {
	rebui.RegisterWidget("MyButton", &MyButton{})
}
