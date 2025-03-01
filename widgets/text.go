package widgets

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kettek/rebui"
)

type Text struct {
	Basic
	blocks          []textBlock
	wrap            rebui.Wrap
	text            string
	face            text.Face
	foregroundColor color.Color
	borderColor     color.Color
	valign          rebui.Alignment
	halign          rebui.Alignment
}

type textBlock struct {
	text   string
	breaks bool
	width  float64
}

func (w *Text) SetBorderColor(clr color.Color) {
	w.borderColor = clr
}

func (w *Text) SetTextWrap(wrap rebui.Wrap) {
	w.wrap = wrap
}

func (w *Text) Layout() {
	w.blocks = nil

	var blocks []textBlock

	var currentBlock textBlock
	var currentWidth float64
	for _, r := range w.text {
		if r == '\n' {
			currentBlock.width, _ = text.Measure(currentBlock.text, w.face, 0)
			blocks = append(blocks, currentBlock)
			blocks = append(blocks, textBlock{breaks: true})
			currentBlock = textBlock{}
			currentWidth = 0
			continue
		}
		currentBlock.text += string(r)
		if w.wrap == rebui.WrapNone {
			continue
		}
		width, _ := text.Measure(currentBlock.text, w.face, 0)
		if currentWidth+width >= w.Width {
			if w.wrap == rebui.WrapWord {
				// Find our last previous space, if possible.
				didit := false
				for i := len(currentBlock.text) - 1; i >= 0; i-- {
					if currentBlock.text[i] == ' ' {
						txt := currentBlock.text
						currentBlock.width, _ = text.Measure(currentBlock.text[:i], w.face, 0)
						currentBlock.text = txt[:i]
						blocks = append(blocks, currentBlock)
						blocks = append(blocks, textBlock{breaks: true})
						currentBlock = textBlock{text: txt[i+1:]}
						didit = true
						break
					}
				}
				// WrapRune when if we fail.
				if !didit {
					blocks, currentBlock = w.genToPreviousRune(blocks, currentBlock)
				}
			} else if w.wrap == rebui.WrapRune {
				blocks, currentBlock = w.genToPreviousRune(blocks, currentBlock)
			}
		}
	}
	// Add last blockie.
	currentBlock.width, _ = text.Measure(currentBlock.text, w.face, 0)
	blocks = append(blocks, currentBlock)

	w.blocks = blocks
}

func (w *Text) genToPreviousRune(blocks []textBlock, currentBlock textBlock) ([]textBlock, textBlock) {
	for i := len(currentBlock.text) - 1; i >= 0; i-- {
		width, _ := text.Measure(currentBlock.text[:i], w.face, 0)
		if width < w.Width {
			txt := currentBlock.text
			currentBlock.width = width
			currentBlock.text = txt[:i]
			blocks = append(blocks, currentBlock)
			blocks = append(blocks, textBlock{breaks: true})
			currentBlock = textBlock{text: txt[i:]}
			break
		}
	}
	return blocks, currentBlock
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
		if block.breaks {
			ty += lineH
			tx = 0
			continue
		}
		txtOptions := &text.DrawOptions{}
		txtOptions.GeoM.Concat(sop.GeoM)
		txtOptions.GeoM.Translate(tx, ty)

		text.Draw(screen, block.text, w.face, txtOptions)

		tx += block.width
	}

	if w.borderColor != nil {
		vector.StrokeRect(screen, float32(x), float32(y), float32(w.Width), float32(w.Height), 1, w.borderColor, false)
	}
}

func init() {
	rebui.RegisterWidget("Text", &Text{})
}
