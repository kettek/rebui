package widgets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kettek/rebui"
)

type Button struct {
	Label
	Border
	backgroundColor color.Color
}

func (b *Button) AssignBackgroundColor(clr color.Color) {
	b.backgroundColor = clr
}

func (b *Button) HandlePointerIn(evt rebui.EventPointerIn) {
	b.backgroundColor = rebui.CurrentTheme().HoverBackgroundColor
	b.borderColor = rebui.CurrentTheme().HoverBorderColor
	b.Label.AssignForegroundColor(rebui.CurrentTheme().HoverForegroundColor)
}

func (b *Button) HandlePointerOut(evt rebui.EventPointerOut) {
	b.backgroundColor = rebui.CurrentTheme().BackgroundColor
	b.borderColor = rebui.CurrentTheme().BorderColor
	b.Label.AssignForegroundColor(rebui.CurrentTheme().ForegroundColor)
}

func (b *Button) HandlePointerPress(evt rebui.EventPointerPress) {
	b.backgroundColor = rebui.CurrentTheme().ActiveBackgroundColor
	b.borderColor = rebui.CurrentTheme().ActiveBorderColor
	b.Label.AssignForegroundColor(rebui.CurrentTheme().ActiveForegroundColor)
}

func (b *Button) HandlePointerRelease(evt rebui.EventPointerRelease) {
	b.backgroundColor = rebui.CurrentTheme().BackgroundColor
	b.borderColor = rebui.CurrentTheme().BorderColor
	b.Label.AssignForegroundColor(rebui.CurrentTheme().ForegroundColor)
}

func (b *Button) HandlePointerPressed(evt rebui.EventPointerPressed) {
	b.backgroundColor = rebui.CurrentTheme().HoverBackgroundColor
	b.borderColor = rebui.CurrentTheme().HoverBorderColor
	b.Label.AssignForegroundColor(rebui.CurrentTheme().HoverForegroundColor)
}

func (b *Button) Draw(screen *ebiten.Image, sop *ebiten.DrawImageOptions) {
	x := sop.GeoM.Element(0, 2)
	y := sop.GeoM.Element(1, 2)

	vector.DrawFilledRect(screen, float32(x), float32(y), float32(b.Width), float32(b.Height), b.backgroundColor, true)
	b.drawBorder(screen, float32(x), float32(y), float32(b.Width), float32(b.Height))

	b.Label.Draw(screen, sop)
}

func init() {
	rebui.RegisterWidget("Button", &Button{})
}
