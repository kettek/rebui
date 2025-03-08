package widgets

import (
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

func (w *Text) AssignBorderColor(clr color.Color) {
	w.borderColor = clr
}

func (w *Text) AssignTextWrap(wrap rebui.Wrap) {
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

func (w *Text) AssignText(text string) {
	w.text = text
}

func (w *Text) AssignForegroundColor(clr color.Color) {
	w.foregroundColor = clr
}

func (w *Text) AssignVerticalAlignment(align rebui.Alignment) {
	w.valign = align
}

func (w *Text) AssignHorizontalAlignment(align rebui.Alignment) {
	w.halign = align
}

func (w *Text) AssignFontFace(face text.Face) {
	w.face = face
}

func (w *Text) AssignFontSize(size float64) {
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

	oy := 0.0
	if w.valign == rebui.AlignMiddle {
		_, oy = blocks.GetSize(w.blocks, blocks.Config{
			Face: w.face,
		})
		oy = (w.Height - oy) / 2
	} else if w.valign == rebui.AlignBottom {
		_, oy = blocks.GetSize(w.blocks, blocks.Config{
			Face: w.face,
		})
		oy = w.Height - oy
	}

	for i := 0; i < len(w.blocks); i++ {
		block := w.blocks[i]

		if _, ok := block.(blocks.Break); ok {
			ty += lineH
			tx = 0
			continue
		}
		if w.halign == rebui.AlignCenter {
			// Collect blocks until break.
			var subBlocks []blocks.Block
			for j := i; j < len(w.blocks); j++ {
				if _, ok := w.blocks[j].(blocks.Break); ok {
					break
				}
				subBlocks = append(subBlocks, w.blocks[j])
			}
			i += len(subBlocks)
			// Get their collective widths.
			var totalWidth float64
			for _, b := range subBlocks {
				if b, ok := b.(blocks.Text); ok {
					totalWidth += b.Width
				}
			}
			// Center them.
			tx = (w.Width - totalWidth) / 2
			// And draw.
			for _, b := range subBlocks {
				if block, ok := b.(blocks.Text); ok {
					txtOptions := &text.DrawOptions{}
					txtOptions.GeoM.Concat(sop.GeoM)
					txtOptions.GeoM.Translate(tx, ty+oy)

					text.Draw(screen, block.Text, w.face, txtOptions)

					tx += block.Width
				}
			}
			ty += lineH
		} else if w.halign == rebui.AlignRight {
			tx = w.Width
			// Find position of next break.
			breakPos := len(w.blocks)
			for j := i; j < len(w.blocks); j++ {
				if _, ok := w.blocks[j].(blocks.Break); ok {
					breakPos = j
					break
				}
			}
			// Draw blocks in reverse from breakPos to i.
			for j := breakPos - 1; j >= i; j-- {
				block := w.blocks[j]
				if textBlock, ok := block.(blocks.Text); ok {
					tx -= textBlock.Width
					txtOptions := &text.DrawOptions{}
					txtOptions.GeoM.Concat(sop.GeoM)
					txtOptions.GeoM.Translate(tx, ty+oy)

					text.Draw(screen, textBlock.Text, w.face, txtOptions)
				}
			}
		} else {
			if textBlock, ok := block.(blocks.Text); ok {
				txtOptions := &text.DrawOptions{}
				txtOptions.GeoM.Concat(sop.GeoM)
				txtOptions.GeoM.Translate(tx, ty+oy)

				text.Draw(screen, textBlock.Text, w.face, txtOptions)

				tx += textBlock.Width
			}
		}
	}

	if w.borderColor != nil {
		vector.StrokeRect(screen, float32(x), float32(y), float32(w.Width), float32(w.Height), 1, w.borderColor, false)
	}
}

func init() {
	rebui.RegisterWidget("Text", &Text{})
}
