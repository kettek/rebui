package rebui

// Node is a parseable structure used for determining element position, style, and beyond.
type Node struct {
	ID              string
	Type            string
	X               string
	x               float64
	Y               string
	y               float64
	Width           string
	width           float64
	Height          string
	height          float64
	OriginX         string
	OriginY         string
	Text            string
	TextWrap        Wrap
	FontSize        string
	Widget          Widget `json:"-"`
	BackgroundColor string
	ForegroundColor string
	BorderColor     string
	VerticalAlign   Alignment
	HorizontalAlign Alignment
	ImageScale      ImageScale
	Image           string // ???
	FocusIndex      int
	nodeHooks
}

type nodeHooks struct {
	OnPointerIn            func(EventPointerIn)
	OnPointerOut           func(EventPointerOut)
	OnPointerMove          func(EventPointerMove)
	OnPointerPress         func(EventPointerPress)
	OnPointerRelease       func(EventPointerRelease)
	OnPointerPressed       func(EventPointerPressed)
	OnPointerGlobalRelease func(EventPointerRelease)
	OnPointerGlobalMove    func(EventPointerMove)
	OnFocus                func(EventFocus)
	OnUnfocus              func(EventUnfocus)
	OnKeyPress             func(EventKeyPress)
	OnKeyRelease           func(EventKeyRelease)
	OnKeyInput             func(EventKeyInput)
}

// pressedNode is a convenience struct that corresponds a given node with a pointer ID.
type pressedNode struct {
	node *Node
	id   int
}
