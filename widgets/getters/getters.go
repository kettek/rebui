package getters

// Width is an interface for getting the width of an element. It is used during layout calculation to see if the user has changed the width of the element manually, and if not, to SetWidth.
type Width interface {
	GetWidth() float64
}

// Height is an interface for getting the height of an element. It is used during layout calculation to see if the user has changed the height of the element manually, and if not, to SetHeight.
type Height interface {
	GetHeight() float64
}

// X is an interface for getting the x position of an element. It is used during layout calculation to see if the user has changed the X position of the element manually, and if not, to SetX.
type X interface {
	GetX() float64
}

// Y is an interface for getting the y position of an element. It is used during layout calculation to see if the user has changed the Y position of the element manually, and if not, to SetY.
type Y interface {
	GetY() float64
}

// Obfuscated is an interface for getting the obfuscation state of an element.
type Obfuscated interface {
	GetObfuscation() bool
}

// BorderWidth is an interface for getting the border width of an element.
type BorderWidth interface {
	GetBorderWidth() float64
}
