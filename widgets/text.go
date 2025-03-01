package widgets

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kettek/rebui"
	"github.com/kettek/rebui/blocks"
)

type Text struct {
	Basic
	blocks          []blocks.Block
	wrap            rebui.Wrap
	text            string
	face            text.Face
	foregroundColor color.Color
	borderColor     color.Color
	valign          rebui.Alignment
	halign          rebui.Alignment
}

func (w *Text) SetBorderColor(clr color.Color) {
	w.borderColor = clr
}

func (w *Text) SetTextWrap(wrap rebui.Wrap) {
	w.wrap = wrap
}

func (w *Text) Layout() {
	w.blocks = blocks.FromText(w.text, blocks.Config{
		Face:   w.face,
		Width:  w.Width,
		Height: w.Height,
		Wrap:   w.wrap,
		VAlign: w.valign,
		HAlign: w.halign,
	})
}

func (w *Text) SetText(text string) {
	w.text = text
}

func (w *Text) SetForegroundColor(clr color.Color) {
	w.foregroundColor = clr
}

func (w *Text) SetVerticalAlignment(align rebui.Alignment) {
	fmt.Println("Text: Vertical Alignment not implemented yet.")
	w.valign = align
}

func (w *Text) SetHorizontalAlignment(align rebui.Alignment) {
	fmt.Println("Text: Horizontal Alignment not implemented yet.")
	w.halign = align
}

func (w *Text) SetFontFace(face text.Face) {
	w.face = face
}

func (w *Text) SetFontSize(size float64) {
	// Re-use FontFace.
	if textFace, ok := w.face.(*text.GoTextFace); ok {
		txt := *textFace
		txt.Size = size
		w.face = &txt
	}
}

func (w *Text) Draw(screen *ebiten.Image, sop *ebiten.DrawImageOptions) {
	x := sop.GeoM.Element(0, 2)
	y := sop.GeoM.Element(1, 2)

	w.Layout() // for now.
	tx, ty := 0.0, 0.0
	lineH := w.face.Metrics().HAscent + w.face.Metrics().HDescent
	for _, block := range w.blocks {
		if _, ok := block.(blocks.Break); ok {
			ty += lineH
			tx = 0
			continue
		} else if block, ok := block.(blocks.Text); ok {
			txtOptions := &text.DrawOptions{}
			txtOptions.GeoM.Concat(sop.GeoM)
			txtOptions.GeoM.Translate(tx, ty)

			text.Draw(screen, block.Text, w.face, txtOptions)

			tx += block.Width
		}
	}

	if w.borderColor != nil {
		vector.StrokeRect(screen, float32(x), float32(y), float32(w.Width), float32(w.Height), 1, w.borderColor, false)
	}
}

func init() {
	rebui.RegisterWidget("Text", &Text{})
}
