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

	bytes, _ := os.ReadFile("buttons.json")

	layout, err := rebui.NewLayout(string(bytes))
	if err != nil {
		log.Fatal(err)
	}

	g.layout = layout
	g.layout.ClampPointers = true
	g.layout.Generate()

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Layout (Ebiten Demo)")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

// DraggableButton is a button that can be dragged around.
type DraggableButton struct {
	widgets.Button
	pressCount int
}

// HandlePointerPress receives a press event if the pointer is over the element.
func (b *DraggableButton) HandlePointerPress(e rebui.EventPointerPress) {
	b.pressCount++
	b.SetBackgroundColor(color.NRGBA{255, 0, 0, 255})
}

// HandlePointerRelease receives a release event if the pointer is over the element.
func (b *DraggableButton) HandlePointerRelease(e rebui.EventPointerRelease) {
	b.pressCount--
	if b.pressCount <= 0 {
		b.SetBackgroundColor(rebui.CurrentTheme().BackgroundColor)
	}
}

// HandlePointerGlobalRelease receives a release event _if_ the element was initially pressed, but did not receive a release over it.
func (b *DraggableButton) HandlePointerGlobalRelease(e rebui.EventPointerRelease) {
	b.pressCount--
	if b.pressCount <= 0 {
		b.SetBackgroundColor(rebui.CurrentTheme().BackgroundColor)
	}
}

// HandlePointerGlobalMove receives a move event if the pointer received the a press event. Note that this will receive move events _per_ pointer, so 3 mouse buttons will move the element at 3 times delta.
func (b *DraggableButton) HandlePointerGlobalMove(e rebui.EventPointerMove) {
	// Ignore default move event.
	if e.PointerID == -1 {
		return
	}
	if e.PointerID == 0 {
		b.X += e.DX
		b.Y += e.DY
	}
	if e.PointerID == 2 {
		b.Width += e.DX
		b.Height += e.DY
	}
}

func init() {
	rebui.RegisterElement("DraggableButton", &DraggableButton{})
}
