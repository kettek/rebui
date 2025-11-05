package widgets

import (
	"image/color"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kettek/rebui"
	"github.com/kettek/rebui/clipboard"
)

type TextInput struct {
	Label
	Border
	text            string
	canvas          *ebiten.Image
	cursor          int
	showCursor      bool
	cursorX         float64
	cursorY         float64
	cursorHeight    float64
	selectInitial   int
	selectStart     int
	selectEnd       int
	ScrollX         float64
	backgroundColor color.Color
	OnChange        func(string)
	OnSubmit        func(string)
	lastTime        time.Time
	cursorHidden    bool
	controlHeld     bool // TODO: Move this to be as part of KeyEvent system.
	obfuscated      bool
}

func (w *TextInput) AssignWidth(width float64) {
	w.Width = width
	w.refreshCanvas()
	w.refreshText()
}

func (w *TextInput) AssignHeight(height float64) {
	w.Height = height
	w.refreshCanvas()
	w.refreshText()
}

func (w *TextInput) AssignText(text string) {
	w.selectStart = 0
	w.selectEnd = 0
	w.text = text
	if w.obfuscated {
		w.Label.AssignText(strings.Repeat("*", len(text)))
	} else {
		w.Label.AssignText(text)
	}
	if w.cursor > len(text) {
		w.cursor = len(text)
	}
	if w.OnChange != nil {
		w.OnChange(text)
	}
	w.refreshText()
}

func (w *TextInput) AssignFontSize(size float64) {
	w.Label.AssignFontSize(size)
	w.refreshText()
}

func (w *TextInput) AssignForegroundColor(clr color.Color) {
	w.Label.AssignForegroundColor(clr)
	w.refreshText()
}

func (w *TextInput) AssignBackgroundColor(clr color.Color) {
	w.backgroundColor = clr
}

func (w *TextInput) AssignObfuscation(b bool) {
	w.obfuscated = b
	w.AssignText(w.text)
}

func (w *TextInput) GetObfuscation() bool {
	return w.obfuscated
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

	if w.selectStart != w.selectEnd {
		startX, _ := text.Measure(w.text[:w.selectStart], w.face, 0)
		endX, _ := text.Measure(w.text[:w.selectEnd], w.face, 0)
		vector.DrawFilledRect(screen, float32(x+startX), float32(y+w.cursorY)-1, float32(endX-startX), float32(w.cursorHeight)+2, color.RGBA{R: 128, G: 128, B: 128, A: 128}, true)
	}

	if w.showCursor && len(w.text) > 0 {
		if time.Since(w.lastTime) > time.Millisecond*500 {
			w.lastTime = time.Now()
			w.cursorHidden = !w.cursorHidden
		}
		if !w.cursorHidden {
			cursorY := y + w.cursorY
			cursorX := x + w.cursorX - w.ScrollX

			vector.StrokeLine(screen, float32(cursorX), float32(cursorY), float32(cursorX), float32(cursorY+w.cursorHeight), 1, w.foregroundColor, false)
		}
	}

	w.drawBorder(screen, float32(x), float32(y), float32(w.Width), float32(w.Height))
}

func (w *TextInput) HandleFocus(evt rebui.EventFocus) {
	w.showCursor = true
	w.refreshCursor()
}

func (w *TextInput) HandleUnfocus(evt rebui.EventUnfocus) {
	w.showCursor = false
}

func (w *TextInput) HandlePointerPress(evt rebui.EventPointerPress) {
	w.cursor = w.getTextIndex(evt.RelativeX)
	w.refreshCursor()
	w.selectInitial = w.cursor
	w.setSelect(w.cursor, w.cursor)
}

func (w *TextInput) getTextIndex(x float64) int {
	if len(w.text) == 0 {
		return 0
	}
	// This seems awful, but I can't think of a more reliable way to fetch such information.
	for i := range w.text {
		width, _ := text.Measure(w.text[:i], w.face, 0)
		if x > width {
			continue
		}
		return i
	}
	return len(w.text)
}

func (w *TextInput) setSelect(x1, x2 int) {
	w.selectStart = x1
	w.selectEnd = x2
}

func (w *TextInput) HandlePointerGlobalMove(evt rebui.EventPointerMove) {
	index := w.getTextIndex(evt.RelativeX)
	if index < w.selectInitial {
		w.setSelect(index, w.selectInitial)
	} else {
		w.setSelect(w.selectInitial, index)
	}
}

func (w *TextInput) HandleKeyInput(evt rebui.EventKeyInput) {
	if w.controlHeld && (evt.Rune == 'v' || evt.Rune == 'c' || evt.Rune == 'a') {
		return
	}
	if w.selectStart != w.selectEnd {
		var text string
		if w.selectStart == 0 {
			text = w.text[w.selectEnd:]
		} else {
			text = w.text[:w.selectStart] + string(evt.Rune) + w.text[w.selectEnd:]
		}
		w.cursor = w.selectStart
		w.AssignText(text)
	} else {
		w.AssignText(w.text[:w.cursor] + string(evt.Rune) + w.text[w.cursor:])
	}
	w.cursor++
	w.refreshCursor()
}

func (w *TextInput) HandleKeyPress(evt rebui.EventKeyPress) {
	if evt.Key == ebiten.KeyBackspace {
		if len(w.text) > 0 {
			if w.selectStart != w.selectEnd {
				var text string
				if w.selectStart == 0 {
					text = w.text[w.selectEnd:]
				} else {
					text = w.text[:w.selectStart] + w.text[w.selectEnd:]
				}
				w.cursor = w.selectStart
				w.AssignText(text)
				w.refreshCursor()
			} else if w.cursor > 0 {
				var text string
				if w.cursor == len(w.text) {
					text = w.text[:len(w.text)-1]
				} else {
					text = w.text[:w.cursor-1] + w.text[w.cursor:]
				}
				w.cursor--
				w.AssignText(text)
				w.refreshCursor()
			}
		}
	} else if evt.Key == ebiten.KeyDelete {
		if len(w.text) > 0 {
			if w.selectStart != w.selectEnd {
				var text string
				if w.selectStart == 0 {
					text = w.text[w.selectEnd:]
				} else {
					text = w.text[:w.selectStart] + w.text[w.selectEnd:]
				}
				w.cursor = w.selectStart
				w.AssignText(text)
				w.refreshCursor()
			} else if w.cursor < len(w.text) {
				var text string
				if w.cursor == len(w.text)-1 {
					text = w.text[:len(w.text)-1]
				} else {
					text = w.text[:w.cursor] + w.text[w.cursor+1:]
				}
				w.AssignText(text)
			}
		}
	} else if evt.Key == ebiten.KeyLeft {
		if w.cursor > 0 {
			w.cursor--
		}
		w.setSelect(0, 0)
		w.refreshCursor()
	} else if evt.Key == ebiten.KeyRight {
		if w.cursor < len(w.text) {
			w.cursor++
		}
		w.setSelect(0, 0)
		w.refreshCursor()
	} else if evt.Key == ebiten.KeyEnter {
		if w.OnSubmit != nil {
			w.OnSubmit(w.text)
		}
	} else if evt.Key == ebiten.KeyControl {
		w.controlHeld = true
	} else if evt.Key == ebiten.KeyC && w.controlHeld {
		if w.selectStart != w.selectEnd {
			clipboard.SetText(w.text[w.selectStart:w.selectEnd])
		}
	} else if evt.Key == ebiten.KeyV && w.controlHeld {
		text := w.text[:w.cursor] + clipboard.GetText() + w.text[w.cursor:]
		w.cursor += len(clipboard.GetText())
		w.AssignText(text)
		w.refreshCursor()
	} else if evt.Key == ebiten.KeyA && w.controlHeld {
		w.setSelect(0, len(w.text))
		w.refreshCursor()
	}
}

func (w *TextInput) HandleKeyRelease(evt rebui.EventKeyRelease) {
	if evt.Key == ebiten.KeyControl {
		w.controlHeld = false
	}
}

func init() {
	rebui.RegisterWidget("TextInput", &TextInput{})
}
