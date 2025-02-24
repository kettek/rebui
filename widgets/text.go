package widgets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kettek/rebui"
)

type Text struct {
	rebui.BasicElement
	text            string
	face            text.Face
	foregroundColor color.Color
	valign          rebui.Alignment
	halign          rebui.Alignment
}

func (w *Text) SetText(text string) {
	w.text = text
}

func (w *Text) SetForegroundColor(clr color.Color) {
	w.foregroundColor = clr
}

func (w *Text) SetVerticalAlignment(align rebui.Alignment) {
	w.valign = align
}

func (w *Text) SetHorizontalAlignment(align rebui.Alignment) {
	w.halign = align
}

func (w *Text) SetFontFace(face text.Face) {
	w.face = face
}

func (w *Text) Draw(screen *ebiten.Image) {
	if w.text != "" && w.face != nil {
		txtOptions := &text.DrawOptions{}
		txtOptions.GeoM.Translate(w.X, w.Y)

		switch w.halign {
		case rebui.Center:
			txtOptions.GeoM.Translate(w.Width/2, 0)
			txtOptions.LayoutOptions.PrimaryAlign = text.AlignCenter
		case rebui.Right:
			txtOptions.LayoutOptions.PrimaryAlign = text.AlignEnd
			txtOptions.GeoM.Translate(w.Width, 0)
		}

		switch w.valign {
		case rebui.Middle:
			txtOptions.GeoM.Translate(0, w.Height/2)
			txtOptions.LayoutOptions.SecondaryAlign = text.AlignCenter
		case rebui.Bottom:
			txtOptions.GeoM.Translate(0, w.Height)
			txtOptions.LayoutOptions.SecondaryAlign = text.AlignEnd
		}

		txtOptions.ColorScale.ScaleWithColor(w.foregroundColor)

		text.Draw(screen, w.text, w.face, txtOptions)
	}
}

func init() {
	rebui.RegisterElement("Text", &Text{})
}
