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
	Element         Element `json:"-"`
	BackgroundColor string
	ForegroundColor string
	BorderColor     string
	VerticalAlign   Alignment
	HorizontalAlign Alignment
	// Internal state.
}

// pressedNode is a convenience struct that corresponds a given node with a pointer ID.
type pressedNode struct {
	node *Node
	id   int
}
