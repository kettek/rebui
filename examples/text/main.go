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
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{}
	g.setupButtons()

	g.layout.AddNode(rebui.Node{
		Type:            "Text",
		ID:              "text",
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

// These are just the buttons for controlling the alignment of the text.
func (g *Game) setupButtons() {
	g.layout.AddNode(rebui.Node{
		Type:   "Button",
		ID:     "button1",
		Width:  "25%",
		Height: "30",
		Text:   "Left",
	}).OnPointerPressed = func(e rebui.EventPointerPressed) {
		txt := g.layout.GetByID("text").Widget.(*widgets.Text)
		txt.SetHorizontalAlignment(rebui.AlignLeft)
	}
	g.layout.AddNode(rebui.Node{
		Type:            "Button",
		ID:              "button2",
		Width:           "25%",
		Height:          "30",
		X:               "after button1",
		Text:            "Center",
		HorizontalAlign: rebui.AlignCenter,
	}).OnPointerPressed = func(e rebui.EventPointerPressed) {
		txt := g.layout.GetByID("text").Widget.(*widgets.Text)
		txt.SetHorizontalAlignment(rebui.AlignCenter)
	}
	g.layout.AddNode(rebui.Node{
		Type:            "Button",
		ID:              "button3",
		Width:           "25%",
		Height:          "30",
		X:               "after button2",
		Text:            "Right",
		HorizontalAlign: rebui.AlignRight,
	}).OnPointerPressed = func(e rebui.EventPointerPressed) {
		txt := g.layout.GetByID("text").Widget.(*widgets.Text)
		txt.SetHorizontalAlignment(rebui.AlignRight)
	}
	g.layout.AddNode(rebui.Node{
		Type:   "Button",
		ID:     "button4",
		Width:  "25%",
		Height: "30",
		Y:      "after button1",
		Text:   "Top",
	}).OnPointerPressed = func(e rebui.EventPointerPressed) {
		txt := g.layout.GetByID("text").Widget.(*widgets.Text)
		txt.SetVerticalAlignment(rebui.AlignTop)
	}
	g.layout.AddNode(rebui.Node{
		Type:          "Button",
		ID:            "button5",
		Width:         "25%",
		Height:        "30",
		X:             "after button4",
		Y:             "after button2",
		Text:          "Middle",
		VerticalAlign: rebui.AlignMiddle,
	}).OnPointerPressed = func(e rebui.EventPointerPressed) {
		txt := g.layout.GetByID("text").Widget.(*widgets.Text)
		txt.SetVerticalAlignment(rebui.AlignMiddle)
	}
	g.layout.AddNode(rebui.Node{
		Type:          "Button",
		ID:            "button6",
		Width:         "25%",
		Height:        "30",
		X:             "after button5",
		Y:             "after button3",
		Text:          "Bottom",
		VerticalAlign: rebui.AlignBottom,
	}).OnPointerPressed = func(e rebui.EventPointerPressed) {
		txt := g.layout.GetByID("text").Widget.(*widgets.Text)
		txt.SetVerticalAlignment(rebui.AlignBottom)
	}
}
