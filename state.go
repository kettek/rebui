package rebui

type currentState struct {
	hoveredNodes []*Node
	pressedNodes []*pressedNode
}

func (s *currentState) isHovered(n *Node) bool {
	for _, hn := range s.hoveredNodes {
		if hn == n {
			return true
		}
	}
	return false
}

func (s *currentState) addHovered(n *Node) {
	s.hoveredNodes = append(s.hoveredNodes, n)
}

func (s *currentState) removeHovered(n *Node) {
	for i, hn := range s.hoveredNodes {
		if hn == n {
			s.hoveredNodes = append(s.hoveredNodes[:i], s.hoveredNodes[i+1:]...)
			return
		}
	}
}

func (s *currentState) isPressed(n *Node, id int) bool {
	for _, pn := range s.pressedNodes {
		if pn.node == n && (id == -1 || pn.id == id) {
			return true
		}
	}
	return false
}

func (s *currentState) addPressed(n *Node, id int) {
	s.pressedNodes = append(s.pressedNodes, &pressedNode{n, id})
}

func (s *currentState) removePressed(n *Node, id int) {
	for i, pn := range s.pressedNodes {
		if pn.node == n && (id == -1 || pn.id == id) {
			s.pressedNodes = append(s.pressedNodes[:i], s.pressedNodes[i+1:]...)
			return
		}
	}
}

func (s *currentState) removePressedID(id int) {
	for i := len(s.pressedNodes) - 1; i >= 0; i-- {
		pn := s.pressedNodes[i]
		if pn.id == id {
			s.pressedNodes = append(s.pressedNodes[:i], s.pressedNodes[i+1:]...)
		}
	}
}
