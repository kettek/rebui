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
	Obfuscated      bool
	Font            string
	FontSize        string
	Widget          Widget `json:"-"`
	BackgroundColor string
	ForegroundColor string
	BorderColor     string
	BorderWidth     string
	VerticalAlign   Alignment
	HorizontalAlign Alignment
	ImageStretch    ImageStretch
	Image           string // ???
	FocusIndex      int
	Children        Nodes
	Hidden          bool
	Parent          *Node // Hmm... uncertain if this paradigm is wise.
	nodeHooks
}

// Nodes are a slice of nodes, wow.
type Nodes []*Node

// GetByID returns nil or a Node pointer, iterating through children.
func (ns *Nodes) GetByID(id string) *Node {
	for _, n := range *ns {
		if n2 := n.getNodeByID(id); n2 != nil {
			return n2
		}
	}
	return nil
}

// ForEach iterates through all nodes and traverses their children. If cb returns true, the for each processing is canceled.
func (ns *Nodes) ForEach(cb func(*Node) bool) bool {
	var childrenByOccurance Nodes
	for _, n := range *ns {
		if n.Hidden {
			continue
		}
		if cb(n) {
			return true
		}
		childrenByOccurance = append(childrenByOccurance, n.Children...)
	}
	// Now process the children in the order we discovered them.
	if len(childrenByOccurance) != 0 && childrenByOccurance.ForEach(cb) {
		return true
	}
	return false
}

// ForEachDeepest iterates through nodes from the first to last with the deepest node in each occurring first.
func (ns *Nodes) ForEachDeepest(cb func(*Node) bool) bool {
	var pending Nodes
	for _, n := range *ns {
		if n.Hidden {
			continue
		}
		pending = append(pending, n.Children...)
	}
	if len(pending) != 0 && pending.ForEachDeepest(cb) {
		return true
	}
	pending = append(pending, *ns...)
	for _, n := range pending {
		if cb(n) {
			return true
		}
	}
	return false
}

// getNodeByID returns any node that has the passed ID, including any nested children.
func (n *Node) getNodeByID(id string) *Node {
	if n.ID == id {
		return n
	}
	for _, ch := range n.Children {
		if n2 := ch.getNodeByID(id); n2 != nil {
			return n2
		}
	}
	return nil
}

func copyNode(n Node) Node {
	var n2 Node
	n2 = n
	n2.Widget = nil // Ensure widget is nil, as we use that to determine if we should create the underlying widget.
	// Also clone the children.
	n2.Children = make(Nodes, len(n.Children))
	for i, child := range n.Children {
		copiedNode := copyNode(*child) // Dereferencing feels sus.
		n2.Children[i] = &copiedNode
		n2.Children[i].Parent = &n2
	}
	return n2
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
