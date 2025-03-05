package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"os"

	_ "image/png"

	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"
	_ "github.com/kettek/rebui/widgets"
)

//go:embed icon.png
var imageBytes []byte

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
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{}

	rebui.SetImageLoader(func(path string) (*ebiten.Image, error) {
		if path == "icon.png" {
			reader := bytes.NewReader(imageBytes)
			img, _, err := image.Decode(reader)
			if err != nil {
				return nil, err
			}
			eimg := ebiten.NewImageFromImage(img)
			return eimg, nil
		}
		return nil, fmt.Errorf("image %s not found", path)
	})

	json, _ := os.ReadFile("layout.json")

	layout, err := rebui.NewLayout(string(json))
	if err != nil {
		log.Fatal(err)
	}

	g.layout = layout
	g.layout.Generate()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Layout (Ebiten Demo)")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
