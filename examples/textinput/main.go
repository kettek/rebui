package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"
	_ "github.com/kettek/rebui/defaults/font"
	"github.com/kettek/rebui/widgets"
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
	return 320, 240
}

func main() {
	g := &Game{}

	node := g.layout.AddNode(rebui.Node{
		Type:            "TextInput",
		Width:           "50%",
		Height:          "30",
		X:               "50%",
		Y:               "50",
		OriginX:         "-50%",
		OriginY:         "-50%",
		ForegroundColor: "white",
		BackgroundColor: "red",
		BorderColor:     "white",
		VerticalAlign:   rebui.AlignMiddle,
		FocusIndex:      1,
	})

	node.Widget.(*widgets.TextInput).OnChange = func(text string) {
		log.Println("Text changed:", text)
	}
	node.Widget.(*widgets.TextInput).OnSubmit = func(text string) {
		node.Widget.(*widgets.TextInput).AssignText("")
		log.Println("Text submitted:", text)
	}

	hideNode := g.layout.AddNode(rebui.Node{
		Type:            "Button",
		Text:            "hide",
		Width:           "40",
		Height:          "30",
		X:               "50%",
		Y:               "90",
		OriginX:         "-50%",
		OriginY:         "-50%",
		ForegroundColor: "white",
		BackgroundColor: "red",
		BorderColor:     "white",
		VerticalAlign:   rebui.AlignMiddle,
		FocusIndex:      2,
	})
	hideNode.OnPointerPressed = func(epp rebui.EventPointerPressed) {
		node.Widget.(*widgets.TextInput).AssignObfuscation(!node.Widget.(*widgets.TextInput).GetObfuscation())
	}

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Layout (Ebiten Demo)")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
