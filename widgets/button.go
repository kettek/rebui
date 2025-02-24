package widgets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kettek/rebui"
)

type Button struct {
	rebui.BasicElement
	text            string
	backgroundColor color.Color
	foregroundColor color.Color
	borderColor     color.Color
	valign          rebui.Alignment
	halign          rebui.Alignment
}

func (b *Button) SetText(text string) {
	b.text = text
}

func (b *Button) SetBackgroundColor(clr color.Color) {
	b.backgroundColor = clr
}

func (b *Button) SetForegroundColor(clr color.Color) {
	b.foregroundColor = clr
}

func (b *Button) SetBorderColor(clr color.Color) {
	b.borderColor = clr
}

func (b *Button) SetVerticalAlignment(align rebui.Alignment) {
	b.valign = align
}

func (b *Button) SetHorizontalAlignment(align rebui.Alignment) {
	b.halign = align
}

func (b *Button) Draw(screen *ebiten.Image) {
	x := b.X
	y := b.Y

	vector.DrawFilledRect(screen, float32(x), float32(y), float32(b.Width), float32(b.Height), b.backgroundColor, true)
	vector.StrokeRect(screen, float32(x), float32(y), float32(b.Width), float32(b.Height), 1, b.borderColor, true)

	if b.text != "" {
		txtOptions := &text.DrawOptions{}
		txtOptions.GeoM.Translate(x, y)

		w, h := text.Measure(b.text, nil, 0)

		switch b.halign {
		case rebui.Center:
			txtOptions.GeoM.Translate(b.Width-w/2, 0)
		case rebui.Right:
			txtOptions.GeoM.Translate(b.Width-w, 0)
		}

		switch b.valign {
		case rebui.Middle:
			txtOptions.GeoM.Translate(0, b.Height-h/2)
		case rebui.Bottom:
			txtOptions.GeoM.Translate(0, b.Height-h)
		}

		text.Draw(screen, b.text, nil, txtOptions)
	}
}

func init() {
	rebui.RegisterElement("Button", &Button{})
}
