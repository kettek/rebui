package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"
	"gopkg.in/yaml.v3"

	// This import sets the default ui font
	_ "github.com/kettek/rebui/defaults/font"
	// This import ensures we have our required widgets.
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

	yml, _ := os.ReadFile("layout.yaml")

	if err := yaml.Unmarshal(yml, &g.layout.Nodes); err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("we got", g.layout.Nodes)
	for i, node := range g.layout.Nodes {
		fmt.Println("Node", i, "is", node)
	}
	g.layout.Generate()

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Layout (Ebiten Demo)")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
