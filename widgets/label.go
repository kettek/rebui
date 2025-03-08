package widgets

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kettek/rebui"
)

type Label struct {
	Basic
	text            string
	face            text.Face
	foregroundColor color.Color
	valign          rebui.Alignment
	halign          rebui.Alignment
}

func (w *Label) AssignText(text string) {
	w.text = text
}

func (w *Label) AssignForegroundColor(clr color.Color) {
	w.foregroundColor = clr
}

func (w *Label) AssignVerticalAlignment(align rebui.Alignment) {
	w.valign = align
}

func (w *Label) AssignHorizontalAlignment(align rebui.Alignment) {
	w.halign = align
}

func (w *Label) AssignFontFace(face text.Face) {
	w.face = face
}

func (w *Label) AssignFontSize(size float64) {
	// Re-use FontFace.
	if textFace, ok := w.face.(*text.GoTextFace); ok {
		txt := *textFace
		txt.Size = size
		w.face = &txt
	}
}

func (w *Label) Draw(screen *ebiten.Image, sop *ebiten.DrawImageOptions) {
	if w.text != "" && w.face != nil {
		txtOptions := &text.DrawOptions{}
		txtOptions.GeoM.Concat(sop.GeoM)
		txtOptions.LineSpacing = w.face.Metrics().HAscent + w.face.Metrics().HDescent // TODO: Check if this is actually correct.

		switch w.halign {
		case rebui.AlignCenter:
			txtOptions.GeoM.Translate(w.Width/2, 0)
			txtOptions.LayoutOptions.PrimaryAlign = text.AlignCenter
		case rebui.AlignRight:
			txtOptions.LayoutOptions.PrimaryAlign = text.AlignEnd
			txtOptions.GeoM.Translate(w.Width, 0)
		}

		switch w.valign {
		case rebui.AlignMiddle:
			txtOptions.GeoM.Translate(0, w.Height/2)
			txtOptions.LayoutOptions.SecondaryAlign = text.AlignCenter
		case rebui.AlignBottom:
			txtOptions.GeoM.Translate(0, w.Height)
			txtOptions.LayoutOptions.SecondaryAlign = text.AlignEnd
		}

		txtOptions.ColorScale.ScaleWithColor(w.foregroundColor)

		text.Draw(screen, w.text, w.face, txtOptions)
	}
}

func init() {
	rebui.RegisterWidget("Label", &Label{})
}
