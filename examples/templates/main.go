package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"

	_ "github.com/kettek/rebui/widgets"

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

	rebui.SetTemplateLoader(func(name string) (rebui.Nodes, error) {
		bytes, err := os.ReadFile("templates/" + name + ".json")
		if err != nil {
			return nil, err
		}
		var nodes rebui.Nodes
		if err := json.Unmarshal(bytes, &nodes); err != nil {
			return nil, err
		}
		return nodes, nil
	})

	bytes, _ := os.ReadFile("layout.json")

	layout, err := rebui.NewLayout(string(bytes))
	if err != nil {
		log.Fatal(err)
	}

	g.layout = layout
	g.layout.Generate()

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Templates")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
