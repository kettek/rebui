package rebui

import (
	"encoding/json"
	"image/color"
	"log"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kettek/rebui/events"
	"github.com/kettek/rebui/style"
	"github.com/kettek/rebui/widgets/assigners"
	"github.com/kettek/rebui/widgets/getters"
	"github.com/kettek/rebui/widgets/receivers"
)

// Layout is used to control layout and manage evts.
type Layout struct {
	RenderTarget  *ebiten.Image
	ClampPointers bool
	generated     bool
	Nodes         []*Node
	currentState  currentState
	//
	relayout            bool
	pressedKeys         []key
	pressedMouseButtons []mouse
	activeTouches       []touch
	focusedNode         *Node
	lastMouseX          int
	lastMouseY          int
	lastWidth           float64
	lastHeight          float64
}

type key struct {
	key   ebiten.Key
	time  time.Time // When this key event was started.
	next  time.Time // Next time to repeat this key.
	count int
}

type mouse struct {
	id   ebiten.MouseButton
	time time.Time // When this mouse event was started.
}

type touch struct {
	id             ebiten.TouchID
	time           time.Time // When this touch was started.
	deltaX, deltaY int       // How much this touch has moved since last time.
	x, y           int       // The current position of the touch.
	movement       int       // How much this touch has moved -- used to determine longpress or move.
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

var handlers = make(map[string]Widget)

// RegisterWidget is used to register an Widget for parsing into the passed in type.
func RegisterWidget(name string, el Widget) {
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

// Generate creates proper Widgets from the list of Nodes.
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
func (l *Layout) AddNode(n Node) *Node {
	l.Nodes = append(l.Nodes, &n)
	l.generateNode(&n)
	l.relayout = true
	return l.Nodes[len(l.Nodes)-1]
}

// RemoveNode removes the given node from the layout.
func (l *Layout) RemoveNode(n *Node) {
	for i, node := range l.Nodes {
		if node == n {
			l.Nodes = append(l.Nodes[:i], l.Nodes[i+1:]...)
			l.relayout = true
			return
		}
	}
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
	evts = append(evts, l.getMouseEvents()...)
	evts = append(evts, l.getTouchEvents()...)
	evts = append(evts, l.getKeyEvents()...)
	return
}

func (l *Layout) getMouseEvents() (evts []Event) {
	x, y := l.getCursor()
	w, h := l.getSize()

	if l.ClampPointers && (x < 0 || y < 0 || x >= w || y >= h) {
		return
	}
	deltaX, deltaY := x-l.lastMouseX, y-l.lastMouseY
	l.lastMouseX, l.lastMouseY = x, y
	ts := time.Now()

	var pressedMouseButtons []mouse
	var newPressedMouseButtons []mouse
	var releasedMouseButtons []mouse
	var oldPressedMouseButtons []mouse
	checkMouseButtons := []ebiten.MouseButton{
		ebiten.MouseButtonLeft,
		ebiten.MouseButtonRight,
		ebiten.MouseButtonMiddle,
	}

	// Get current press state.
	for _, mb := range checkMouseButtons {
		if ebiten.IsMouseButtonPressed(mb) {
			pressedMouseButtons = append(pressedMouseButtons, mouse{id: mb, time: ts})
		}
	}

	// Check for any releases.
	for _, mb := range l.pressedMouseButtons {
		exists := false
		for _, mb2 := range pressedMouseButtons {
			if mb.id == mb2.id {
				exists = true
				break
			}
		}
		if !exists {
			releasedMouseButtons = append(releasedMouseButtons, mb)
		}
	}

	// Check for any new presses.
	for i, mb := range pressedMouseButtons {
		exists := false
		var prevMouse mouse
		for _, mb2 := range l.pressedMouseButtons {
			if mb.id == mb2.id {
				prevMouse = mb2
				exists = true
				break
			}
		}
		if !exists {
			newPressedMouseButtons = append(newPressedMouseButtons, mb)
		} else {
			oldPressedMouseButtons = append(oldPressedMouseButtons, prevMouse)
			pressedMouseButtons[i] = prevMouse // Ensure the mouse data is the same as from last frame, as pressedMouseButtons replaces l.pressedMouseButtons.
		}
	}

	// Alright, let's convert them to pointer evts.
	for _, mb := range newPressedMouseButtons {
		evts = append(evts, events.PointerPress{
			Timestamp: events.Timestamp{Timestamp: ts},
			Pointer: events.Pointer{
				X:        float64(x),
				Y:        float64(y),
				DX:       float64(deltaX),
				DY:       float64(deltaY),
				ButtonID: int(mb.id),
			},
		})
	}

	for _, mb := range releasedMouseButtons {
		evts = append(evts, events.PointerRelease{
			Timestamp: events.Timestamp{Timestamp: ts},
			Duration:  events.Duration{Duration: ts.Sub(mb.time)},
			Pointer: events.Pointer{
				X:        float64(x),
				Y:        float64(y),
				DX:       float64(deltaX),
				DY:       float64(deltaY),
				ButtonID: int(mb.id),
			},
		})
	}

	if deltaX != 0 || deltaY != 0 {
		for _, mb := range oldPressedMouseButtons {
			evts = append(evts, events.PointerMove{
				Timestamp: events.Timestamp{Timestamp: ts},
				Duration:  events.Duration{Duration: ts.Sub(mb.time)},
				Pointer: events.Pointer{
					X:        float64(x),
					Y:        float64(y),
					DX:       float64(deltaX),
					DY:       float64(deltaY),
					ButtonID: int(mb.id),
				},
			})
		}
		// And have an event for no pointer (-1)
		evts = append(evts, events.PointerMove{
			Timestamp: events.Timestamp{Timestamp: ts},
			Pointer: events.Pointer{
				X:        float64(x),
				Y:        float64(y),
				DX:       float64(deltaX),
				DY:       float64(deltaY),
				ButtonID: -1,
			},
		})
	}

	// Replace the old.
	l.pressedMouseButtons = pressedMouseButtons

	return
}

func (l *Layout) getTouchEvents() (evts []Event) {
	var activeTouches []touch
	var releasedTouches []touch
	var newTouches []touch
	var oldTouches []touch

	ts := time.Now()

	// Get current touch state.
	for _, id := range ebiten.AppendTouchIDs(nil) {
		tx, ty := ebiten.TouchPosition(id)
		activeTouches = append(activeTouches, touch{id: id, x: tx, y: ty, time: ts})
	}

	// Check for any touch releases.
	for _, t := range l.activeTouches {
		exists := false
		var newTouch touch
		for _, t2 := range activeTouches {
			if t.id == t2.id {
				exists = true
				newTouch = t2
				break
			}
		}
		if !exists {
			t.deltaX = newTouch.x - t.x
			t.deltaY = newTouch.y - t.y
			t.movement += int(math.Abs(float64(t.deltaX)) + math.Abs(float64(t.deltaY)))
			releasedTouches = append(releasedTouches, t)
		}
	}

	// Check for any new touches.
	for i, t := range activeTouches {
		exists := false
		var prevTouch touch
		for _, t2 := range l.activeTouches {
			if t.id == t2.id {
				prevTouch = t2
				exists = true
				break
			}
		}
		if !exists {
			newTouches = append(newTouches, t)
		} else {
			tx, ty := ebiten.TouchPosition(t.id)
			prevTouch.deltaX = tx - prevTouch.x
			prevTouch.deltaY = ty - prevTouch.y
			prevTouch.movement += int(math.Abs(float64(prevTouch.deltaX)) + math.Abs(float64(prevTouch.deltaY)))
			prevTouch.x = tx
			prevTouch.y = ty
			oldTouches = append(oldTouches, prevTouch)
			activeTouches[i] = prevTouch // Ensure the touch data is the same as from last frame, as activeTouches replaces l.activeTouches.
		}
	}

	// Convert to events.
	for _, t := range newTouches {
		evts = append(evts, events.PointerPress{
			Timestamp: events.Timestamp{Timestamp: ts},
			Pointer: events.Pointer{
				X:       float64(t.x),
				Y:       float64(t.y),
				DX:      float64(t.deltaX),
				DY:      float64(t.deltaY),
				TouchID: int(t.id),
			},
		})
	}

	for _, t := range releasedTouches {
		evts = append(evts, events.PointerRelease{
			Timestamp: events.Timestamp{Timestamp: ts},
			Duration:  events.Duration{Duration: ts.Sub(t.time)},
			Pointer: events.Pointer{
				X:       float64(t.x),
				Y:       float64(t.y),
				DX:      float64(t.deltaX),
				DY:      float64(t.deltaY),
				TouchID: int(t.id),
			},
		})
	}

	for _, t := range oldTouches {
		if t.deltaX != 0 || t.deltaY != 0 {
			evts = append(evts, events.PointerMove{
				Timestamp: events.Timestamp{Timestamp: ts},
				Duration:  events.Duration{Duration: ts.Sub(t.time)},
				Pointer: events.Pointer{
					X:       float64(t.x),
					Y:       float64(t.y),
					DX:      float64(t.deltaX),
					DY:      float64(t.deltaY),
					TouchID: int(t.id),
				},
			})
		}
	}

	l.activeTouches = activeTouches

	return
}

func (l *Layout) getKeyEvents() (evts []Event) {
	ts := time.Now()

	var pressedKeys []key
	var releasedKeys []key
	var newPressedKeys []key
	var repeatKeys []key

	for _, k := range inpututil.AppendPressedKeys(nil) {
		pressedKeys = append(pressedKeys, key{key: k, time: ts, next: ts.Add(500 * time.Millisecond)})
	}

	for _, k := range l.pressedKeys {
		exists := false
		for _, k2 := range pressedKeys {
			if k.key == k2.key {
				exists = true
				break
			}
		}
		if !exists {
			releasedKeys = append(releasedKeys, k)
		}
	}
	for i, k := range pressedKeys {
		exists := false
		var prevKey key
		for _, k2 := range l.pressedKeys {
			if k.key == k2.key {
				prevKey = k2
				exists = true
				break
			}
		}
		if !exists {
			newPressedKeys = append(newPressedKeys, k)
		} else {
			if prevKey.next.Before(ts) {
				repeatKeys = append(repeatKeys, key{
					key:   prevKey.key,
					time:  ts,
					count: prevKey.count + 1,
				})
				prevKey.next = ts.Add(50 * time.Millisecond)
				prevKey.count++
			}
			pressedKeys[i] = prevKey
		}
	}
	for _, k := range newPressedKeys {
		evts = append(evts, events.KeyPress{
			Timestamp: events.Timestamp{Timestamp: ts},
			Key:       k.key,
		})
	}
	for _, k := range releasedKeys {
		evts = append(evts, events.KeyRelease{
			Timestamp: events.Timestamp{Timestamp: ts},
			Key:       k.key,
			Duration:  events.Duration{Duration: ts.Sub(k.time)},
		})
	}

	for _, k := range repeatKeys {
		evts = append(evts, events.KeyPress{
			Timestamp: events.Timestamp{Timestamp: ts},
			Key:       k.key,
			Repeat:    k.count,
		})
	}

	// Also handle input chars. AFAIK we shouldn't handle the whole key press logic with input chars.
	for _, k := range ebiten.AppendInputChars(nil) {
		evts = append(evts, events.KeyInput{
			Timestamp: events.Timestamp{Timestamp: ts},
			Rune:      k,
		})
	}

	l.pressedKeys = pressedKeys

	return
}

// Update collects evts and propagates them to the contained Widgets.
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
			l.processEvent(e)
		}
	}
}

// Draw draws the Nodes to the screen
func (l *Layout) Draw(screen *ebiten.Image) {
	l.RenderTarget = screen
	if l.lastWidth != float64(screen.Bounds().Dx()) || l.lastHeight != float64(screen.Bounds().Dy()) {
		l.lastWidth = float64(screen.Bounds().Dx())
		l.lastHeight = float64(screen.Bounds().Dy())
		l.relayout = true
	}

	for _, n := range l.Nodes {
		op := &ebiten.DrawImageOptions{}
		if n.Widget != nil {
			if xg, ok := n.Widget.(getters.X); ok {
				op.GeoM.Translate(xg.GetX(), 0)
			} else {
				op.GeoM.Translate(n.x, 0)
			}
			if yg, ok := n.Widget.(getters.Y); ok {
				op.GeoM.Translate(0, yg.GetY())
			} else {
				op.GeoM.Translate(0, n.y)
			}
			n.Widget.Draw(l.RenderTarget, op)
		}
	}
}

// HasEvents returns if there are any active events like a mouse press,
func (l *Layout) HasEvents() bool {
	if len(l.currentState.hoveredNodes) > 0 || len(l.currentState.pressedNodes) > 0 || len(l.pressedKeys) > 0 || len(l.activeTouches) > 0 || len(l.pressedMouseButtons) > 0 {
		return true
	}
	return false
}

// ClearEvents clears all events that have been processed, such as pointer presses, key presses, etc.
func (l *Layout) ClearEvents() {
	l.currentState.hoveredNodes = nil
	l.currentState.pressedNodes = nil
	l.pressedKeys = nil
	l.activeTouches = nil
	l.pressedMouseButtons = nil
	l.focusedNode = nil
}

func (l *Layout) generateNode(n *Node) {
	// Might as well prevent re-generation.
	if n.Widget != nil {
		return
	}
	for k, h := range handlers {
		if k == n.Type {
			n.Widget = reflect.New(reflect.TypeOf(h).Elem()).Interface().(Widget)
			// Call our setter interfaces if desired.
			if bcs, ok := n.Widget.(assigners.BackgroundColor); ok {
				bcs.AssignBackgroundColor(stringToColor(n.BackgroundColor, style.CurrentTheme().BackgroundColor))
			}
			if fcs, ok := n.Widget.(assigners.ForegroundColor); ok {
				fcs.AssignForegroundColor(stringToColor(n.ForegroundColor, style.CurrentTheme().ForegroundColor))
			}
			if bcs, ok := n.Widget.(assigners.BorderColor); ok {
				bcs.AssignBorderColor(stringToColor(n.BorderColor, style.CurrentTheme().BorderColor))
			}
			if vas, ok := n.Widget.(assigners.VerticalAlignment); ok {
				vas.AssignVerticalAlignment(n.VerticalAlign)
			}
			if has, ok := n.Widget.(assigners.HorizontalAlignment); ok {
				has.AssignHorizontalAlignment(n.HorizontalAlign)
			}
			if ts, ok := n.Widget.(assigners.Text); ok {
				ts.AssignText(n.Text)
			}
			if tws, ok := n.Widget.(assigners.TextWrap); ok {
				tws.AssignTextWrap(n.TextWrap)
			}
			if fs, ok := n.Widget.(assigners.FontFace); ok {
				fs.AssignFontFace(style.CurrentTheme().FontFace)
			}
			if n.Font != "" {
				if ff, ok := n.Widget.(assigners.FontFace); ok {
					face, err := LoadFont(n.Font)
					if err == nil {
						ff.AssignFontFace(face)
					} else {
						log.Println(err)
					}
				}
			}
			if n.FontSize != "" {
				if fs, ok := n.Widget.(assigners.FontSize); ok {
					if ff, ok := style.CurrentTheme().FontFace.(*text.GoTextFace); ok {
						size := stringToPosition(l, n.FontSize, ff.Size, true) // FIXME: This re-use is goofy, as it allows unintended at/after usage.
						fs.AssignFontSize(size)
					}
				}
			}
			if is, ok := n.Widget.(assigners.ImageStretch); ok {
				is.AssignImageStretch(n.ImageStretch)
			}
			if is, ok := n.Widget.(assigners.Image); ok {
				img, err := LoadImage(n.Image)
				if err == nil {
					is.AssignImage(img)
				} else {
					log.Println(err)
				}
			}
		}
	}
}

// layoutNode sets the node's various positions and sizings based upon the containing outer width and height.
func (l *Layout) layoutNode(n *Node, outerWidth, outerHeight float64) {
	nodeWidth := outerWidth
	nodeHeight := outerHeight
	nodeX := 0.0
	nodeY := 0.0

	var skipWidth bool
	var skipHeight bool
	if wg, ok := n.Widget.(getters.Width); ok {
		if wg.GetWidth() != n.width {
			skipWidth = true
		}
	}
	if hg, ok := n.Widget.(getters.Height); ok {
		if hg.GetHeight() != n.height {
			skipHeight = true
		}
	}

	if !skipWidth && n.Width != "" {
		nodeWidth = stringToPosition(l, n.Width, outerWidth, false)
	}
	if !skipHeight && n.Height != "" {
		nodeHeight = stringToPosition(l, n.Height, outerHeight, true)
	}

	// Allow the widget to layout its final size.
	if lw, ok := n.Widget.(LayoutWidget); ok {
		nodeWidth, nodeHeight = lw.Layout(nodeWidth, nodeHeight)
	}
	// And then assign it.
	if wa, ok := n.Widget.(assigners.Width); ok {
		wa.AssignWidth(nodeWidth)
	}
	if ha, ok := n.Widget.(assigners.Height); ok {
		ha.AssignHeight(nodeHeight)
	}

	n.width = nodeWidth
	n.height = nodeHeight

	// Check if X has changed by comparing any user-set value to our stored node value.
	var skipX bool
	var skipY bool
	if xg, ok := n.Widget.(getters.X); ok {
		if xg.GetX() != n.x {
			skipX = true
		}
	}
	if yg, ok := n.Widget.(getters.Y); ok {
		if yg.GetY() != n.y {
			skipY = true
		}
	}

	if !skipX {
		// Origin uses the node's own width and height to determine offsets.
		originX := stringToPosition(l, n.OriginX, nodeWidth, false)
		if oxs, ok := n.Widget.(assigners.OriginX); ok {
			oxs.AssignOriginX(originX)
		}
		if n.X != "" {
			nodeX = stringToPosition(l, n.X, outerWidth, false)
			if xs, ok := n.Widget.(assigners.X); ok {
				xs.AssignX(nodeX + originX)
			}
			n.x = nodeX + originX
		}
	}
	if !skipY {
		originY := stringToPosition(l, n.OriginY, nodeHeight, true)
		if oys, ok := n.Widget.(assigners.OriginY); ok {
			oys.AssignOriginY(originY)
		}
		if n.Y != "" {
			nodeY = stringToPosition(l, n.Y, outerHeight, true)
			if ys, ok := n.Widget.(assigners.Y); ok {
				ys.AssignY(nodeY + originY)
			}
			n.y = nodeY + originY
		}
	}
}

func (l *Layout) processNodeEvent(n *Node, e Event) {
	if hit, ok := n.Widget.(HitChecker); ok {
		switch evt := e.(type) {
		case events.PointerMove:
			if hit.Hit(evt.X, evt.Y) {
				evt.Widget = n.Widget
				if gx, ok := n.Widget.(getters.X); ok {
					evt.RelativeX = evt.X - gx.GetX()
				}
				if gy, ok := n.Widget.(getters.Y); ok {
					evt.RelativeY = evt.Y - gy.GetY()
				}
				if n.OnPointerMove != nil {
					n.OnPointerMove(&evt)
				}
				if hmove, ok := n.Widget.(receivers.PointerMove); ok {
					hmove.HandlePointerMove(&evt)
				}
				if !l.currentState.isHovered(n) {
					pointerInEvent := events.PointerIn{
						TargetWidget: events.TargetWidget{Widget: n.Widget},
						Timestamp:    evt.Timestamp,
						Pointer:      evt.Pointer,
					}
					if n.OnPointerIn != nil {
						n.OnPointerIn(&pointerInEvent)
					}
					if hin, ok := n.Widget.(receivers.PointerIn); ok {
						hin.HandlePointerIn(&pointerInEvent)
					}
					l.currentState.addHovered(n)
				}
			} else {
				if l.currentState.isHovered(n) {
					pointerOutEvent := events.PointerOut{
						TargetWidget: events.TargetWidget{Widget: n.Widget},
						Timestamp:    evt.Timestamp,
						Pointer:      evt.Pointer,
					}
					if n.OnPointerOut != nil {
						n.OnPointerOut(&pointerOutEvent)
					}
					if hout, ok := n.Widget.(receivers.PointerOut); ok {
						hout.HandlePointerOut(&pointerOutEvent)
					}
					l.currentState.removeHovered(n)
				}
			}
		case events.PointerPress:
			if hit.Hit(evt.X, evt.Y) {
				pid := -1
				if evt.TouchID > 0 { // I hope touches can't be 0...
					pid = evt.TouchID
				} else {
					pid = evt.ButtonID
				}
				evt.Widget = n.Widget
				if gx, ok := n.Widget.(getters.X); ok {
					evt.RelativeX = evt.X - gx.GetX()
				}
				if gy, ok := n.Widget.(getters.Y); ok {
					evt.RelativeY = evt.Y - gy.GetY()
				}
				if n.OnPointerPress != nil {
					n.OnPointerPress(&evt)
				}
				if hpress, ok := n.Widget.(receivers.PointerPress); ok {
					hpress.HandlePointerPress(&evt)
				}
				if !l.currentState.isPressed(n, pid) {
					l.currentState.addPressed(n, pid)
				}
				// Focus node on pointer press -- do we want to limit focused to only first or last node receiving a pointer press?
				if l.focusedNode != nil && l.focusedNode != n {
					unfocusEvent := events.Unfocus{
						TargetWidget: events.TargetWidget{Widget: l.focusedNode.Widget},
						Timestamp:    evt.Timestamp,
					}
					if l.focusedNode.OnUnfocus != nil {
						l.focusedNode.OnUnfocus(&unfocusEvent)
					}
					if hunfocus, ok := l.focusedNode.Widget.(receivers.Unfocus); ok {
						hunfocus.HandleUnfocus(&unfocusEvent)
					}
				}
				if n.FocusIndex > 0 {
					if l.focusedNode != n {
						focusEvent := events.Focus{
							TargetWidget: events.TargetWidget{Widget: evt.Widget},
							Timestamp:    evt.Timestamp,
							Pointer:      evt.Pointer,
						}
						if n.OnFocus != nil {
							n.OnFocus(&focusEvent)
						}
						if hfocus, ok := n.Widget.(receivers.Focus); ok {
							hfocus.HandleFocus(&focusEvent)
						}
						l.focusedNode = n
					}
				} else {
					l.focusedNode = nil
				}
			}
		case events.PointerRelease:
			if hit.Hit(evt.X, evt.Y) {
				pid := -1
				if evt.TouchID > 0 { // I hope touches can't be 0...
					pid = evt.TouchID
				} else {
					pid = evt.ButtonID
				}
				evt.Widget = n.Widget
				if gx, ok := n.Widget.(getters.X); ok {
					evt.RelativeX = evt.X - gx.GetX()
				}
				if gy, ok := n.Widget.(getters.Y); ok {
					evt.RelativeY = evt.Y - gy.GetY()
				}
				if n.OnPointerRelease != nil {
					n.OnPointerRelease(&evt)
				}
				if hrelease, ok := n.Widget.(receivers.PointerRelease); ok {
					hrelease.HandlePointerRelease(&evt)
				}
				if l.currentState.isPressed(n, pid) {
					l.currentState.removePressed(n, pid)
					pointerPressedEvent := events.PointerPressed{
						TargetWidget: events.TargetWidget{Widget: n.Widget},
						Duration:     evt.Duration,
						Timestamp:    evt.Timestamp,
						Pointer:      evt.Pointer,
					}
					if n.OnPointerPressed != nil {
						n.OnPointerPressed(&pointerPressedEvent)
					}
					if hpress, ok := n.Widget.(receivers.PointerPressed); ok {
						hpress.HandlePointerPressed(&pointerPressedEvent)
					}
				}
			}
		}
	}
}

// processEvent is called after processNodeEvent and does any further handling beyond what the nodes can handle.
func (l *Layout) processEvent(e Event) {
	switch evt := e.(type) {
	case events.PointerPress:
		// Unfocus the current focused node if we have a press that does not hit it.
		if l.focusedNode != nil {
			if hit, ok := l.focusedNode.Widget.(HitChecker); ok {
				if hit.Hit(evt.X, evt.Y) {
					// We hit the focused node, so we don't need to do anything.
					break
				}
				unfocusEvent := events.Unfocus{
					TargetWidget: events.TargetWidget{Widget: l.focusedNode.Widget},
					Timestamp:    events.Timestamp{Timestamp: time.Now()},
				}
				if l.focusedNode.OnUnfocus != nil {
					l.focusedNode.OnUnfocus(&unfocusEvent)
				}
				if hunfocus, ok := l.focusedNode.Widget.(receivers.Unfocus); ok {
					hunfocus.HandleUnfocus(&unfocusEvent)
				}
				l.focusedNode = nil
			}
		}
	case events.PointerRelease:
		pid := -1
		if evt.TouchID > 0 { // I hope touches can't be 0...
			pid = evt.TouchID
		} else {
			pid = evt.ButtonID
		}
		// Clear out any held releases.
		for _, n := range l.Nodes {
			if l.currentState.isPressed(n, pid) {
				evt.Widget = n.Widget
				if gx, ok := n.Widget.(getters.X); ok {
					evt.RelativeX = evt.X - gx.GetX()
				}
				if gy, ok := n.Widget.(getters.Y); ok {
					evt.RelativeY = evt.Y - gy.GetY()
				}
				if n.OnPointerGlobalRelease != nil {
					n.OnPointerGlobalRelease(&evt)
				}
				if hrelease, ok := n.Widget.(receivers.PointerGlobalRelease); ok {
					hrelease.HandlePointerGlobalRelease(&evt)
				}
			}
		}
		l.currentState.removePressedID(pid)
	case events.PointerMove:
		pid := -1
		if evt.TouchID > 0 { // I hope touches can't be 0...
			pid = evt.TouchID
		} else {
			pid = evt.ButtonID
		}
		// Handle any global move handlers that were pressed.
		for _, n := range l.Nodes {
			evt.Widget = n.Widget
			if gx, ok := n.Widget.(getters.X); ok {
				evt.RelativeX = evt.X - gx.GetX()
			}
			if gy, ok := n.Widget.(getters.Y); ok {
				evt.RelativeY = evt.Y - gy.GetY()
			}
			if l.currentState.isPressed(n, pid) {
				if n.OnPointerGlobalMove != nil {
					n.OnPointerGlobalMove(&evt)
				}
				if hmove, ok := n.Widget.(receivers.PointerGlobalMove); ok {
					hmove.HandlePointerGlobalMove(&evt)
				}
			}
		}
	case events.KeyPress:
		if l.focusedNode != nil {
			evt.Widget = l.focusedNode.Widget
			if l.focusedNode.OnKeyPress != nil {
				l.focusedNode.OnKeyPress(&evt)
			}
			if evt.Canceled() {
				break
			}
			if n, ok := l.focusedNode.Widget.(receivers.KeyPress); ok {
				n.HandleKeyPress(&evt)
			}
		}
	case events.KeyRelease:
		if l.focusedNode != nil {
			evt.Widget = l.focusedNode.Widget
			if l.focusedNode.OnKeyRelease != nil {
				l.focusedNode.OnKeyRelease(&evt)
			}
			if evt.Canceled() {
				break
			}
			if n, ok := l.focusedNode.Widget.(receivers.KeyRelease); ok {
				n.HandleKeyRelease(&evt)
			}
		}
	case events.KeyInput:
		if l.focusedNode != nil {
			evt.Widget = l.focusedNode.Widget
			if l.focusedNode.OnKeyInput != nil {
				l.focusedNode.OnKeyInput(&evt)
			}
			if evt.Canceled() {
				break
			}
			if n, ok := l.focusedNode.Widget.(receivers.KeyInput); ok {
				n.HandleKeyInput(&evt)
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
