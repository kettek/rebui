package rebui

import (
	"encoding/json"
	"image/color"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kettek/rebui/elements/receivers"
	"github.com/kettek/rebui/elements/setters"
	"github.com/kettek/rebui/events"
	"github.com/kettek/rebui/style"
)

// Layout is used to control layout and manage evts.
type Layout struct {
	RenderTarget  *ebiten.Image
	ClampPointers bool
	generated     bool
	Nodes         []*Node
	currentState  currentState
	//
	imageLoader         func(string) (*ebiten.Image, error)
	relayout            bool
	pressedMouseButtons []ebiten.MouseButton
	lastMouseX          int
	lastMouseY          int
	lastWidth           float64
	lastHeight          float64
}

// GetByID returns the given node by its ID.
func (l *Layout) GetByID(id string) *Node {
	for _, n := range l.Nodes {
		if n.ID == id {
			return n
		}
	}
	return nil
}

var handlers = make(map[string]Element)

// RegisterElement is used to register an Element for parsing into the passed in type.
func RegisterElement(name string, el Element) {
	handlers[name] = el
}

// NewLayout creates a new layout from the given JSON string.
func NewLayout(src string) (*Layout, error) {
	l := &Layout{}
	err := l.Parse(src)
	if err != nil {
		return nil, err
	}
	return l, nil
}

// Parse parses the given JSON source string.
func (l *Layout) Parse(src string) error {
	l.Nodes = make([]*Node, 0)
	reader := json.NewDecoder(strings.NewReader(src))
	return reader.Decode(&l.Nodes)
}

// Generate creates proper Elements from the list of Nodes.
func (l *Layout) Generate() {
	for _, n := range l.Nodes {
		l.generateNode(n)
	}
	l.relayout = true
}

// Layout repositions all nodes.
func (l *Layout) Layout(ow, oh float64) {
	if l.lastWidth != ow || l.lastHeight != oh {
		for _, n := range l.Nodes {
			l.layoutNode(n, ow, oh)
		}
		l.lastWidth = ow
		l.lastHeight = oh
	}
}

// AddNode adds the given node and generates it.
func (l *Layout) AddNode(n Node) {
	l.Nodes = append(l.Nodes, &n)
	l.generateNode(&n)
	l.relayout = true
}

func (l *Layout) getCursor() (x, y int) {
	if l.RenderTarget != nil {
		w, h := l.RenderTarget.Bounds().Dx(), l.RenderTarget.Bounds().Dy()
		x, y = ebiten.CursorPosition()
		x = int((float64(x) / float64(w)) * float64(w))
		y = int((float64(y) / float64(h)) * float64(h))
	} else {
		x, y = ebiten.CursorPosition()
	}
	return
}

func (l *Layout) getSize() (w, h int) {
	w, h = ebiten.WindowSize()
	if l.RenderTarget != nil {
		w, h = l.RenderTarget.Bounds().Dx(), l.RenderTarget.Bounds().Dy()
	}
	return
}

func (l *Layout) getEvents() (evts []Event) {
	x, y := l.getCursor()
	w, h := l.getSize()

	if l.ClampPointers && (x < 0 || y < 0 || x >= w || y >= h) {
		return
	}
	deltaX, deltaY := x-l.lastMouseX, y-l.lastMouseY
	l.lastMouseX, l.lastMouseY = x, y
	ts := time.Now()

	var pressedMouseButtons []ebiten.MouseButton
	var newPressedMouseButtons []ebiten.MouseButton
	var releasedMouseButtons []ebiten.MouseButton
	var oldPressedMouseButtons []ebiten.MouseButton
	checkMouseButtons := []ebiten.MouseButton{
		ebiten.MouseButtonLeft,
		ebiten.MouseButtonRight,
		ebiten.MouseButtonMiddle,
	}

	// Get current press state.
	for _, mb := range checkMouseButtons {
		if ebiten.IsMouseButtonPressed(mb) {
			pressedMouseButtons = append(pressedMouseButtons, mb)
		}
	}

	// Check for any releases.
	for i := 0; i < len(l.pressedMouseButtons); i++ {
		btn := l.pressedMouseButtons[i]
		exists := false
		for _, mb := range pressedMouseButtons {
			if btn == mb {
				exists = true
				break
			}
		}
		if !exists {
			releasedMouseButtons = append(releasedMouseButtons, btn)
		}
	}

	// Check for any new presses.
	for _, mb := range pressedMouseButtons {
		exists := false
		for i := 0; i < len(l.pressedMouseButtons); i++ {
			btn := l.pressedMouseButtons[i]
			if btn == mb {
				exists = true
				break
			}
		}
		if !exists {
			newPressedMouseButtons = append(newPressedMouseButtons, mb)
		} else {
			oldPressedMouseButtons = append(oldPressedMouseButtons, mb)
		}
	}

	// Alright, let's convert them to pointer evts.
	for _, mb := range newPressedMouseButtons {
		evts = append(evts, events.PointerPress{
			Timestamp: events.Timestamp{Timestamp: ts},
			Pointer: events.Pointer{
				X:         float64(x),
				Y:         float64(y),
				DX:        float64(deltaX),
				DY:        float64(deltaY),
				PointerID: int(mb),
			},
		})
	}

	for _, mb := range releasedMouseButtons {
		evts = append(evts, events.PointerRelease{
			Timestamp: events.Timestamp{Timestamp: ts},
			Pointer: events.Pointer{
				X:         float64(x),
				Y:         float64(y),
				DX:        float64(deltaX),
				DY:        float64(deltaY),
				PointerID: int(mb),
			},
		})
	}

	if deltaX != 0 || deltaY != 0 {
		for _, mb := range oldPressedMouseButtons {
			evts = append(evts, events.PointerMove{
				Timestamp: events.Timestamp{Timestamp: ts},
				Pointer: events.Pointer{
					X:         float64(x),
					Y:         float64(y),
					DX:        float64(deltaX),
					DY:        float64(deltaY),
					PointerID: int(mb),
				},
			})
		}
		// And have an event for no pointer (-1)
		evts = append(evts, events.PointerMove{
			Timestamp: events.Timestamp{Timestamp: ts},
			Pointer: events.Pointer{
				X:         float64(x),
				Y:         float64(y),
				DX:        float64(deltaX),
				DY:        float64(deltaY),
				PointerID: -1,
			},
		})
	}

	// Replace the old.
	l.pressedMouseButtons = pressedMouseButtons

	return
}

// Update collects evts and propagates them to the contained Elements.
func (l *Layout) Update() {
	if l.relayout {
		w, h := l.getSize()
		l.Layout(float64(w), float64(h))
		l.relayout = false
	}

	// TODO: Allow passing in a block evts list, where various event types can be prevented from occurring -- this might come in use.
	if evts := l.getEvents(); len(evts) > 0 {
		for _, e := range evts {
			// Iterate our nodes...
			for _, n := range l.Nodes {
				l.processNodeEvent(n, e)
				if ec, ok := e.(EventCancelable); ok && ec.Canceled() {
					break
				}
			}
			switch evt := e.(type) {
			case events.PointerRelease:
				// Clear out any held releases.
				for _, n := range l.Nodes {
					if l.currentState.isPressed(n, evt.PointerID) {
						evt.Target = n.Element
						if n.OnPointerGlobalRelease != nil {
							n.OnPointerGlobalRelease(&evt)
						}
						if hrelease, ok := n.Element.(receivers.PointerGlobalRelease); ok {
							hrelease.HandlePointerGlobalRelease(&evt)
						}
					}
				}
				l.currentState.removePressedID(evt.PointerID)
			case events.PointerMove:
				// Handle any global move handlers that were pressed.
				for _, n := range l.Nodes {
					evt.Target = n.Element
					if l.currentState.isPressed(n, evt.PointerID) {
						if n.OnPointerGlobalMove != nil {
							n.OnPointerGlobalMove(&evt)
						}
						if hmove, ok := n.Element.(receivers.PointerGlobalMove); ok {
							hmove.HandlePointerGlobalMove(&evt)
						}
					}
				}
			}
		}
	}
}

// Draw draws the Nodes to the screen
func (l *Layout) Draw(screen *ebiten.Image) {
	l.RenderTarget = screen
	for _, n := range l.Nodes {
		if n.Element != nil {
			n.Element.Draw(screen)
		}
	}
}

func (l *Layout) generateNode(n *Node) {
	// Might as well prevent re-generation.
	if n.Element != nil {
		return
	}
	for k, h := range handlers {
		if k == n.Type {
			n.Element = reflect.New(reflect.TypeOf(h).Elem()).Interface().(Element)
			// Call our setter interfaces if desired.
			if bcs, ok := n.Element.(setters.BackgroundColor); ok {
				bcs.SetBackgroundColor(stringToColor(n.BackgroundColor, style.CurrentTheme().BackgroundColor))
			}
			if fcs, ok := n.Element.(setters.ForegroundColor); ok {
				fcs.SetForegroundColor(stringToColor(n.ForegroundColor, style.CurrentTheme().ForegroundColor))
			}
			if bcs, ok := n.Element.(setters.BorderColor); ok {
				bcs.SetBorderColor(stringToColor(n.BorderColor, style.CurrentTheme().BorderColor))
			}
			if vas, ok := n.Element.(setters.VerticalAlignment); ok {
				vas.SetVerticalAlignment(n.VerticalAlign)
			}
			if has, ok := n.Element.(setters.HorizontalAlignment); ok {
				has.SetHorizontalAlignment(n.HorizontalAlign)
			}
			if ts, ok := n.Element.(setters.Text); ok {
				ts.SetText(n.Text)
			}
			if fs, ok := n.Element.(setters.FontFace); ok {
				fs.SetFontFace(style.CurrentTheme().FontFace)
			}
			if n.FontSize != "" {
				if fs, ok := n.Element.(setters.FontSize); ok {
					if ff, ok := style.CurrentTheme().FontFace.(*text.GoTextFace); ok {
						size := stringToPosition(l, n.FontSize, ff.Size, true) // FIXME: This re-use is goofy, as it allows unintended at/after usage.
						fs.SetFontSize(size)
					}
				}
			}
			if is, ok := n.Element.(setters.ImageScale); ok {
				is.SetImageScale(n.ImageScale)
			}
			if is, ok := n.Element.(setters.Image); ok {
				if l.imageLoader != nil {
					img, err := l.imageLoader(n.Image)
					if err == nil {
						is.SetImage(img)
					} else {
						log.Println(err)
					}
				}
			}
		}
	}
}

// SetImageLoader can be used to set a loader that controls loading images by string.
func (l *Layout) SetImageLoader(cb func(string) (*ebiten.Image, error)) {
	l.imageLoader = cb
}

// layoutNode sets the node's various positions and sizings based upon the containing outer width and height.
func (l *Layout) layoutNode(n *Node, outerWidth, outerHeight float64) {
	nodeWidth := outerWidth
	nodeHeight := outerHeight
	nodeX := 0.0
	nodeY := 0.0
	if n.Width != "" {
		nodeWidth = stringToPosition(l, n.Width, outerWidth, false)
		if ws, ok := n.Element.(setters.Width); ok {
			ws.SetWidth(nodeWidth)
		}
		n.width = nodeWidth
	}
	if n.Height != "" {
		nodeHeight = stringToPosition(l, n.Height, outerHeight, true)
		if hs, ok := n.Element.(setters.Height); ok {
			hs.SetHeight(nodeHeight)
		}
		n.height = nodeHeight
	}
	// Origin uses the node's own widht and height to determine offsets.
	originX := stringToPosition(l, n.OriginX, nodeWidth, false)
	if oxs, ok := n.Element.(setters.OriginX); ok {
		oxs.SetOriginX(originX)
	}
	originY := stringToPosition(l, n.OriginY, nodeHeight, true)
	if oys, ok := n.Element.(setters.OriginY); ok {
		oys.SetOriginY(originY)
	}
	if n.X != "" {
		nodeX = stringToPosition(l, n.X, outerWidth, false)
		if xs, ok := n.Element.(setters.X); ok {
			xs.SetX(nodeX + originX)
		}
		n.x = nodeX + originX
	}
	if n.Y != "" {
		nodeY = stringToPosition(l, n.Y, outerHeight, true)
		if ys, ok := n.Element.(setters.Y); ok {
			ys.SetY(nodeY + originY)
		}
		n.y = nodeY + originY
	}
}

func (l *Layout) processNodeEvent(n *Node, e Event) {
	if hit, ok := n.Element.(HitChecker); ok {
		switch evt := e.(type) {
		case events.PointerMove:
			if hit.Hit(evt.X, evt.Y) {
				evt.Target = n.Element
				if n.OnPointerMove != nil {
					n.OnPointerMove(&evt)
				}
				if hmove, ok := n.Element.(receivers.PointerMove); ok {
					hmove.HandlePointerMove(&evt)
				}
				if !l.currentState.isHovered(n) {
					pointerInEvent := events.PointerIn{
						TargetElement: events.TargetElement{Target: n.Element},
						Timestamp:     evt.Timestamp,
						Pointer:       evt.Pointer,
					}
					if n.OnPointerIn != nil {
						n.OnPointerIn(&pointerInEvent)
					}
					if hin, ok := n.Element.(receivers.PointerIn); ok {
						hin.HandlePointerIn(&pointerInEvent)
					}
					l.currentState.addHovered(n)
				}
			} else {
				if l.currentState.isHovered(n) {
					pointerOutEvent := events.PointerOut{
						TargetElement: events.TargetElement{Target: n.Element},
						Timestamp:     evt.Timestamp,
						Pointer:       evt.Pointer,
					}
					if n.OnPointerOut != nil {
						n.OnPointerOut(&pointerOutEvent)
					}
					if hout, ok := n.Element.(receivers.PointerOut); ok {
						hout.HandlePointerOut(&pointerOutEvent)
					}
					l.currentState.removeHovered(n)
				}
			}
		case events.PointerPress:
			if hit.Hit(evt.X, evt.Y) {
				evt.Target = n.Element
				if n.OnPointerPress != nil {
					n.OnPointerPress(&evt)
				}
				if hpress, ok := n.Element.(receivers.PointerPress); ok {
					hpress.HandlePointerPress(&evt)
				}
				if !l.currentState.isPressed(n, e.(events.PointerPress).PointerID) {
					l.currentState.addPressed(n, e.(events.PointerPress).PointerID)
				}
			}
		case events.PointerRelease:
			if hit.Hit(evt.X, evt.Y) {
				evt.Target = n.Element
				if n.OnPointerRelease != nil {
					n.OnPointerRelease(&evt)
				}
				if hrelease, ok := n.Element.(receivers.PointerRelease); ok {
					hrelease.HandlePointerRelease(&evt)
				}
				if l.currentState.isPressed(n, evt.PointerID) {
					l.currentState.removePressed(n, evt.PointerID)
					pointerPressedEvent := events.PointerPressed{
						TargetElement: events.TargetElement{Target: n.Element},
						Timestamp:     evt.Timestamp,
						Pointer:       evt.Pointer,
					}
					if n.OnPointerPressed != nil {
						n.OnPointerPressed(&pointerPressedEvent)
					}
					if hpress, ok := n.Element.(receivers.PointerPressed); ok {
						hpress.HandlePointerPressed(&pointerPressedEvent)
					}
				}
			}
		}
	}
}

func stringToColor(s string, fallback color.Color) color.Color {
	if s == "" {
		return fallback
	}
	if s[0] == '#' {
		if len(s) == 4 { // Allow lazy RGB->RRGGBB
			rr := string(s[1] + s[1])
			gg := string(s[2] + s[2])
			bb := string(s[3] + s[3])
			s = "#" + rr + gg + bb
		}
		if len(s) == 7 {
			r, _ := strconv.ParseInt(s[1:3], 16, 0)
			g, _ := strconv.ParseInt(s[3:5], 16, 0)
			b, _ := strconv.ParseInt(s[5:7], 16, 0)
			return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		} else if len(s) == 9 {
			r, _ := strconv.ParseInt(s[1:3], 16, 0)
			g, _ := strconv.ParseInt(s[3:5], 16, 0)
			b, _ := strconv.ParseInt(s[5:7], 16, 0)
			a, _ := strconv.ParseInt(s[7:9], 16, 0)
			return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		}
	}
	// Simple crummy name parsing.
	switch s {
	case "black":
		return color.Black
	case "white":
		return color.White
	case "red":
		return color.RGBA{255, 0, 0, 255}
	case "green":
		return color.RGBA{0, 255, 0, 255}
	case "blue":
		return color.RGBA{0, 0, 255, 255}
	}
	return color.Black
}

func stringToPosition(l *Layout, s string, outer float64, vertical bool) float64 {
	if s == "" {
		return 0
	}
	if strings.HasPrefix(s, "after ") {
		after := l.GetByID(s[6:])
		if after != nil {
			if vertical {
				return after.y + after.height
			}
			return after.x + after.width
		}
	} else if strings.HasPrefix(s, "at ") {
		at := l.GetByID(s[3:])
		if at != nil {
			if vertical {
				return at.y
			}
			return at.x
		}
	} else if s[len(s)-1] == '%' {
		percent := s[:len(s)-1]
		p, _ := strconv.ParseFloat(percent, 64)
		return (p / 100) * outer
	} else {
		reg := regexp.MustCompile(`(\d+)%\sof\s(.*)$`)

		matches := reg.FindStringSubmatch(s)
		if len(matches) == 3 {
			percent := matches[1]
			target := l.GetByID(matches[2])
			p, _ := strconv.ParseFloat(percent, 64)
			p = (p / 100)
			if target != nil {
				if vertical {
					return target.height * p
				}
				return target.width * p
			}
		}
	}
	// Finally, let's just try to get non %
	p, _ := strconv.ParseFloat(s, 64)
	return p
}
