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
)

// Layout is used to control layout and manage events.
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
	for _, n := range l.Nodes {
		l.layoutNode(n, ow, oh)
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

func (l *Layout) getEvents() (events []Event) {
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

	// Alright, let's convert them to pointer events.
	for _, mb := range newPressedMouseButtons {
		events = append(events, &pointerPressEvent{
			TimestampEvent: TimestampEvent{Timestamp: ts},
			PointerEvent: PointerEvent{
				X:         float64(x),
				Y:         float64(y),
				DX:        float64(deltaX),
				DY:        float64(deltaY),
				PointerID: int(mb),
			},
		})
	}

	for _, mb := range releasedMouseButtons {
		events = append(events, &pointerReleaseEvent{
			TimestampEvent: TimestampEvent{Timestamp: ts},
			PointerEvent: PointerEvent{
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
			events = append(events, &pointerMoveEvent{
				TimestampEvent: TimestampEvent{Timestamp: ts},
				PointerEvent: PointerEvent{
					X:         float64(x),
					Y:         float64(y),
					DX:        float64(deltaX),
					DY:        float64(deltaY),
					PointerID: int(mb),
				},
			})
		}
		// And have an event for no pointer (-1)
		events = append(events, &pointerMoveEvent{
			TimestampEvent: TimestampEvent{Timestamp: ts},
			PointerEvent: PointerEvent{
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

// Update collects events and propagates them to the contained Elements.
func (l *Layout) Update() {
	if l.relayout {
		w, h := l.getSize()
		l.Layout(float64(w), float64(h))
		l.relayout = false
	}

	// TODO: Allow passing in a block events list, where various event types can be prevented from occurring -- this might come in use.
	if events := l.getEvents(); len(events) > 0 {
		for _, e := range events {
			// Iterate our nodes...
			for _, n := range l.Nodes {
				l.processNodeEvent(n, e)
				if e.Canceled() {
					break
				}
			}
			switch event := e.(type) {
			case *pointerReleaseEvent:
				// Clear out any held releases.
				for _, n := range l.Nodes {
					if l.currentState.isPressed(n, event.PointerID) {
						if hrelease, ok := n.Element.(GlobalReleaseReceiver); ok {
							hrelease.HandlePointerGlobalRelease(event)
						}
					}
				}
				l.currentState.removePressedID(event.PointerID)
			case *pointerMoveEvent:
				// Handle any global move handlers that were pressed.
				for _, n := range l.Nodes {
					if l.currentState.isPressed(n, event.PointerID) {
						if hmove, ok := n.Element.(GlobalMoveReceiver); ok {
							hmove.HandlePointerGlobalMove(event)
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
			if bcs, ok := n.Element.(BackgroundColorSetter); ok {
				bcs.SetBackgroundColor(stringToColor(n.BackgroundColor, DefaultTheme.BackgroundColor))
			}
			if fcs, ok := n.Element.(ForegroundColorSetter); ok {
				fcs.SetForegroundColor(stringToColor(n.ForegroundColor, DefaultTheme.ForegroundColor))
			}
			if bcs, ok := n.Element.(BorderColorSetter); ok {
				bcs.SetBorderColor(stringToColor(n.BorderColor, DefaultTheme.BorderColor))
			}
			if vas, ok := n.Element.(VerticalAlignmentSetter); ok {
				vas.SetVerticalAlignment(n.VerticalAlign)
			}
			if has, ok := n.Element.(HorizontalAlignmentSetter); ok {
				has.SetHorizontalAlignment(n.HorizontalAlign)
			}
			if ts, ok := n.Element.(TextSetter); ok {
				ts.SetText(n.Text)
			}
			if is, ok := n.Element.(ImageScaleSetter); ok {
				is.SetImageScale(n.ImageScale)
			}
			if is, ok := n.Element.(ImageSetter); ok {
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
		if ws, ok := n.Element.(WidthSetter); ok {
			ws.SetWidth(nodeWidth)
		}
		n.width = nodeWidth
	}
	if n.Height != "" {
		nodeHeight = stringToPosition(l, n.Height, outerHeight, true)
		if hs, ok := n.Element.(HeightSetter); ok {
			hs.SetHeight(nodeHeight)
		}
		n.height = nodeHeight
	}
	// Origin uses the node's own widht and height to determine offsets.
	originX := stringToPosition(l, n.OriginX, nodeWidth, false)
	if oxs, ok := n.Element.(OriginXSetter); ok {
		oxs.SetOriginX(originX)
	}
	originY := stringToPosition(l, n.OriginY, nodeHeight, true)
	if oys, ok := n.Element.(OriginYSetter); ok {
		oys.SetOriginY(originY)
	}
	if n.X != "" {
		nodeX = stringToPosition(l, n.X, outerWidth, false)
		if xs, ok := n.Element.(XSetter); ok {
			xs.SetX(nodeX + originX)
		}
		n.x = nodeX + originX
	}
	if n.Y != "" {
		nodeY = stringToPosition(l, n.Y, outerHeight, true)
		if ys, ok := n.Element.(YSetter); ok {
			ys.SetY(nodeY + originY)
		}
		n.y = nodeY + originY
	}
}

func (l *Layout) processNodeEvent(n *Node, e Event) {
	if hit, ok := n.Element.(HitChecker); ok {
		switch event := e.(type) {
		case PointerMoveEvent:
			if hit.Hit(event.X, event.Y) {
				if hmove, ok := n.Element.(PointerMoveReceiver); ok {
					hmove.HandlePointerMove(event)
				}
				if !l.currentState.isHovered(n) {
					if hin, ok := n.Element.(PointerInReceiver); ok {
						hin.HandlePointerIn(&pointerInEvent{
							TimestampEvent: event.TimestampEvent,
							PointerEvent:   event.PointerEvent,
						})
					}
					l.currentState.addHovered(n)
				}
			} else {
				if l.currentState.isHovered(n) {
					if hout, ok := n.Element.(PointerOutReceiver); ok {
						hout.HandlePointerOut(&pointerOutEvent{
							TimestampEvent: event.TimestampEvent,
							PointerEvent:   event.PointerEvent,
						})
					}
					l.currentState.removeHovered(n)
				}
			}
		case PointerPressEvent:
			if hit.Hit(event.X, event.Y) {
				if hpress, ok := n.Element.(PressReceiver); ok {
					hpress.HandlePointerPress(event)
				}
				if !l.currentState.isPressed(n, e.(*pointerPressEvent).PointerID) {
					l.currentState.addPressed(n, e.(*pointerPressEvent).PointerID)
				}
			}
		case PointerReleaseEvent:
			if hit.Hit(event.X, event.Y) {
				if hrelease, ok := n.Element.(ReleaseReceiver); ok {
					hrelease.HandlePointerRelease(event)
				}
				if l.currentState.isPressed(n, event.PointerID) {
					l.currentState.removePressed(n, event.PointerID)
					if hpress, ok := n.Element.(PressedReceiver); ok {
						hpress.HandlePointerPressed(&pointerPressedEvent{
							TimestampEvent: event.TimestampEvent,
							PointerEvent:   event.PointerEvent,
						})
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
