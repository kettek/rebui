package main

import (
	"bytes"
	"log"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
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

//go:embed x10y12pxDonguriDuel.ttf
var font1Bytes []byte

//go:embed x12y16pxSolidLinker.ttf
var font2Bytes []byte

func main() {
	g := &Game{}

	rebui.SetFontLoader(func(name string) (text.Face, error) {
		var b []byte
		if name == "x10y12pxDonguriDuel" {
			b = font1Bytes
		} else if name == "x12y16pxSolidLinker" {
			b = font2Bytes
		}
		s, err := text.NewGoTextFaceSource(bytes.NewReader(b))
		if err != nil {
			return nil, err
		}
		return &text.GoTextFace{
			Source: s,
			Size:   12,
		}, nil
	})

	font := "x10y12pxDonguriDuel"
	g.layout.AddNode(rebui.Node{
		Type:            "Text",
		ID:              "text",
		Width:           "50%",
		Height:          "20%",
		X:               "50%",
		Y:               "50%",
		OriginX:         "-50%",
		OriginY:         "-50%",
		BorderColor:     "white",
		BackgroundColor: "red",
		Font:            "x10y12pxDonguriDuel",
		FontSize:        "12",
		Text:            "Click me to change font :)",
		TextWrap:        rebui.WrapWord,
		HorizontalAlign: rebui.AlignCenter,
		VerticalAlign:   rebui.AlignMiddle,
	}).OnPointerPressed = func(evt rebui.EventPointerPressed) {
		size := 12.0
		if font == "x10y12pxDonguriDuel" {
			font = "x12y16pxSolidLinker"
			size = 16.0
		} else {
			font = "x10y12pxDonguriDuel"
		}
		font, _ := rebui.LoadFont(font)
		evt.Widget.(*widgets.Text).AssignFontFace(font)
		evt.Widget.(*widgets.Text).AssignFontSize(size)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Layout (Ebiten Demo)")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
