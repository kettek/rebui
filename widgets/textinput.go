package widgets

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kettek/rebui"
)

type TextInput struct {
	Label
	canvas          *ebiten.Image
	cursor          int
	showCursor      bool
	cursorX         float64
	cursorY         float64
	cursorHeight    float64
	ScrollX         float64
	borderColor     color.Color
	backgroundColor color.Color
	OnChange        func(string)
	OnSubmit        func(string)
	lastTime        time.Time
	cursorHidden    bool
}

func (w *TextInput) SetWidth(width float64) {
	w.Width = width
	w.refreshCanvas()
	w.refreshText()
}

func (w *TextInput) SetHeight(height float64) {
	w.Height = height
	w.refreshCanvas()
	w.refreshText()
}

func (w *TextInput) SetText(text string) {
	w.Label.SetText(text)
	if w.cursor > len(w.text) {
		w.cursor = len(w.text)
	}
	if w.OnChange != nil {
		w.OnChange(text)
	}
	w.refreshText()
}

func (w *TextInput) SetFontSize(size float64) {
	w.Label.SetFontSize(size)
	w.refreshText()
}

func (w *TextInput) SetForegroundColor(clr color.Color) {
	w.Label.SetForegroundColor(clr)
	w.refreshText()
}

func (w *TextInput) SetBackgroundColor(clr color.Color) {
	w.backgroundColor = clr
}

func (w *TextInput) SetBorderColor(clr color.Color) {
	w.borderColor = clr
}

func (w *TextInput) refreshCanvas() {
	if w.Width == 0 || w.Height == 0 {
		return
	}
	if w.canvas != nil {
		w.canvas.Deallocate()
	}
	w.canvas = ebiten.NewImage(int(w.Width), int(w.Height))
}

func (w *TextInput) refreshText() {
	if w.canvas == nil {
		return
	}
	sop := &ebiten.DrawImageOptions{}
	sop.GeoM.Translate(-w.ScrollX, 0)
	w.canvas.Clear()
	w.Label.Draw(w.canvas, sop)
}

func (w *TextInput) refreshCursor() {
	w.cursorHeight = w.face.Metrics().HAscent + w.face.Metrics().HDescent
	if w.cursor == len(w.text) {
		w.cursorX, _ = text.Measure(w.text, w.face, 0)
	} else {
		w.cursorX, _ = text.Measure(w.text[:w.cursor], w.face, 0)
	}
	// TODO: Implement halign logic for cursor.
	/*switch w.halign {
	case rebui.AlignCenter:
		w.cursorX -= w.Width / 2
	case rebui.AlignRight:
		w.cursorX = w.Width - w.cursorX
	}*/
	switch w.valign {
	case rebui.AlignMiddle:
		w.cursorY = w.Height/2 - w.cursorHeight/2
	case rebui.AlignBottom:
		w.cursorY = w.Height - w.cursorHeight
	}

	w.cursorHidden = false
	w.lastTime = time.Now()
}

func (w *TextInput) Draw(screen *ebiten.Image, sop *ebiten.DrawImageOptions) {
	if w.canvas == nil {
		w.refreshCanvas()
	}

	x := sop.GeoM.Element(0, 2)
	y := sop.GeoM.Element(1, 2)

	vector.DrawFilledRect(screen, float32(x), float32(y), float32(w.Width), float32(w.Height), w.backgroundColor, true)

	screen.DrawImage(w.canvas, sop)

	if w.showCursor && len(w.text) > 0 {
		if time.Since(w.lastTime) > time.Millisecond*500 {
			w.lastTime = time.Now()
			w.cursorHidden = !w.cursorHidden
		}
		if !w.cursorHidden {
			cursorY := y + w.cursorY
			cursorX := x + w.cursorX - w.ScrollX

			vector.StrokeLine(screen, float32(cursorX), float32(cursorY), float32(cursorX), float32(cursorY+w.cursorHeight)-1, 1, w.foregroundColor, false)
		}
	}

	vector.StrokeRect(screen, float32(x), float32(y), float32(w.Width), float32(w.Height), 1, w.borderColor, false)
}

func (w *TextInput) HandleFocus(evt rebui.EventFocus) {
	w.showCursor = true
	w.refreshCursor()
}

func (w *TextInput) HandleUnfocus(evt rebui.EventUnfocus) {
	w.showCursor = false
}

func (w *TextInput) HandlePointerPress(evt rebui.EventPointerPress) {
	found := false
	// This seems awful, but I can't think of a more reliable way to fetch such information.
	for i := range w.text {
		width, _ := text.Measure(w.text[:i], w.face, 0)
		if evt.RelativeX > width {
			continue
		}
		w.cursor = i
		found = true
		break
	}
	if !found {
		w.cursor = len(w.text)
	}
	w.refreshCursor()
}

func (w *TextInput) HandleKeyInput(evt rebui.EventKeyInput) {
	w.SetText(w.text[:w.cursor] + string(evt.Rune) + w.text[w.cursor:])
	w.cursor++
	w.refreshCursor()
}

func (w *TextInput) HandleKeyPress(evt rebui.EventKeyPress) {
	if evt.Key == ebiten.KeyBackspace {
		if len(w.text) > 0 {
			if w.cursor > 0 {
				var text string
				if w.cursor == len(w.text) {
					text = w.text[:len(w.text)-1]
				} else {
					text = w.text[:w.cursor-1] + w.text[w.cursor:]
				}
				w.cursor--
				w.SetText(text)
				w.refreshCursor()
			}
		}
	} else if evt.Key == ebiten.KeyDelete {
		if len(w.text) > 0 {
			if w.cursor < len(w.text) {
				var text string
				if w.cursor == len(w.text)-1 {
					text = w.text[:len(w.text)-1]
				} else {
					text = w.text[:w.cursor] + w.text[w.cursor+1:]
				}
				w.SetText(text)
			}
		}
	} else if evt.Key == ebiten.KeyLeft {
		if w.cursor > 0 {
			w.cursor--
		}
		w.refreshCursor()
	} else if evt.Key == ebiten.KeyRight {
		if w.cursor < len(w.text) {
			w.cursor++
		}
		w.refreshCursor()
	} else if evt.Key == ebiten.KeyEnter {
		if w.OnSubmit != nil {
			w.OnSubmit(w.text)
		}
	}
}

func init() {
	rebui.RegisterWidget("TextInput", &TextInput{})
}
