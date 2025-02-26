# rebui

[![Go Reference](https://pkg.go.dev/badge/github.com/kettek/rebui.svg)](https://pkg.go.dev/github.com/kettek/rebui)

**rebui** is a relatively simple UI library for Ebitengine. It manages UI layout and event handling in a clear and dynamic fashion. It _does not_ seek to directly implement flexbox, grids, or other automatic container-based layout systems, but rather provides a simple-to-use framework to create and use widgets within a shared screen space. However, flexbox or grid-like functionality can be implemented, as the layout engine allows placing widgets relative to one another and/or using percentage-based size and position calculations.

## How Does It Look?

A full example of a large red button in the middle of the screen would be:

```golang
package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"
	// Blank import ensures all the default widgets are available.
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

	g.layout.AddNode(rebui.Node{
		Type:            "Button",
		Width:           "50%",
		Height:          "50%",
		X:               "50%",
		Y:               "50%",
		OriginX:         "-50%",
		OriginY:         "-50%",
		BackgroundColor: "red",
	})

	ebiten.SetWindowSize(320, 240)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
```

Another example with a draggable and resizable custom widget would be:

```golang
package main

import (
	"image/color"
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
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{}

	g.layout.AddNode(rebui.Node{
		Type:            "DraggableButton",
		Width:           "25%",
		Height:          "25%",
		X:               "50%",
		Y:               "50%",
		OriginX:         "-50%",
		OriginY:         "-50%",
	})

	ebiten.SetWindowSize(320, 240)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

type DraggableButton struct {
	widgets.Button
}

func (b *DraggableButton) HandlePointerGlobalMove(e rebui.EventPointerMove) {
	// left mouse button to drag
	if e.PointerID == 0 {
		b.X += e.DX
		b.Y += e.DY
	}
	// right mouse button to resize
	if e.PointerID == 2 {
		b.Width += e.DX
		b.Height += e.DY
	}
}

func init() {
	rebui.RegisterWidget("DraggableButton", &DraggableButton{})
}

```
