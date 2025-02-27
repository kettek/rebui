package main

import (
	"fmt"
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
		Type:        "MyButton",
		ID:          "button1",
		FocusIndex:  1,
		Width:       "25%",
		Height:      "25%",
		X:           "50%",
		Y:           "25%",
		OriginX:     "-50%",
		OriginY:     "-50%",
		BorderColor: "white",
	})

	g.layout.AddNode(rebui.Node{
		Type:        "MyButton",
		ID:          "button2",
		FocusIndex:  0,
		Width:       "25%",
		Height:      "25%",
		X:           "at button1",
		Y:           "after button1",
		BorderColor: "white",
	})

	g.layout.AddNode(rebui.Node{
		Type:        "MyButton",
		FocusIndex:  1,
		Width:       "25%",
		Height:      "25%",
		X:           "at button2",
		Y:           "after button2",
		BorderColor: "white",
	})

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Focus")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

type MyButton struct {
	widgets.Button
}

func (b *MyButton) HandleFocus(e rebui.EventFocus) {
	b.SetBorderColor(color.NRGBA{0, 255, 0, 255})
}

func (b *MyButton) HandleUnfocus(e rebui.EventUnfocus) {
	b.SetBorderColor(color.NRGBA{255, 255, 255, 255})
}

func (b *MyButton) HandleKeyPress(e rebui.EventKeyPress) {
	fmt.Println("got key press", e)
}

func (b *MyButton) HandleKeyRelease(e rebui.EventKeyRelease) {
	fmt.Println("got key release", e)
}

func init() {
	rebui.RegisterWidget("MyButton", &MyButton{})
}
