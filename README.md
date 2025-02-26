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

func (b *DraggableButton) HandlePointerGlobalMove(evt rebui.EventPointerMove) {
	// left mouse button to drag
	if evt.PointerID == 0 {
		b.X += e.DX
		b.Y += e.DY
	}
	// right mouse button to resize
	if evt.PointerID == 2 {
		b.Width += e.DX
		b.Height += e.DY
	}
}

func init() {
	rebui.RegisterWidget("DraggableButton", &DraggableButton{})
}

```

For most use-cases, Nodes will not be defined manually, but through an external layout file. The default format is through JSON, but any type of unmarshalling type/process can be used.

## Coordinates, Dimensions, and IDs

rebui Nodes can have their coordinates and dimensions specify pixel values, percentage values of their container, or percentage values of another node.

For, example, a JSON layout with 2 buttons next to each other on the x axis, one of which is half the size of the other, would be as follows:

```json
[
  {
    "Type": "Button",
    "ID": "button1",
    "X": "50",
    "Y": "50",
    "Width": "100",
    "Height": "100"
  },
  {
    "Type": "Button",
    "X": "after button1",
    "Y": "at button1",
    "Width": "50% of button1",
    "Height": "50% of button1"
  }
]
```

## Origin

An additional step when determining a Node's position is the OriginX and OriginY values. These values are relative to the dimensions of the node, so to have a node that spawns in the middle of the screen centered about its own middle-point, would be:

```json
{
  "Type": "Button",
  "X": "50%",
  "Y": "50%",
  "OriginX": "-50%",
  "OriginY": "-50%",
  "Width": "100",
  "Height": "100"
}
```

## GetByID

rebui Nodes can be accessed by calling the `Layout.GetByID(string)` method. This allows one to retrieve the Node which also contains the `node.Widget` field that can be used to directly access the underlying widget.

```golang
node := layout.GetByID("button1")
node.Widget.(*widgets.Button).SetBackgroundColor(color.Black)
```

## Events

rebui Events are sent to Widget handlers or to Node event callbacks.

A custom Widget that receives a pointer pressed event:

```golang
type MyButton struct {
	widgets.Button
}

func (b *MyButton) HandlePointerPressed(evt rebui.EventPointerPressed) {
	b.SetBackgroundColor(color.NRGBA{255, 0, 255, 255})
}
```

A locally scoped Node variable:

```golang
node := layout.GetByID("button1")

node.OnPointerPressed = func(evt rebui.EventPointerPressed) {
	node.Widget.(*widgets.Button).SetBackgroundColor(color.NRGBA{255, 0, 255, 255})
}

```

Or a Widget acquired from the pointer event:

```golang
layout.GetByID("button1").OnPointerPressed = func(evt rebui.EventPointerPressed) {
	evt.Widget.(*widgets.Button).SetBackgroundColor(color.NRGBA{255, 0, 255, 255})
}
```
