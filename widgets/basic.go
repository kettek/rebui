package widgets

// Basic provides the core functionality for positioning and testing for hits.
type Basic struct {
	X, Y, Width, Height float64
	OriginX, OriginY    float64
	Disabled            bool
}

// Hit returns true if the given x and y coordinates are within the bounds of the element.
func (b *Basic) Hit(x, y float64) bool {
	if b.Disabled {
		return false
	}
	return x >= b.X && x <= b.X+b.Width && y >= b.Y && y <= b.Y+b.Height
}

// AssignWidth sets the width of the element.
func (b *Basic) AssignWidth(w float64) {
	b.Width = w
}

// GetWidth returns the width of the element.
func (b *Basic) GetWidth() float64 {
	return b.Width
}

// AssignHeight sets the height of the element.
func (b *Basic) AssignHeight(h float64) {
	b.Height = h
}

// GetHeight returns the height of the element.
func (b *Basic) GetHeight() float64 {
	return b.Height
}

// AssignX sets the x position of the element.
func (b *Basic) AssignX(x float64) {
	b.X = x
}

// GetX returns the x position of the element.
func (b *Basic) GetX() float64 {
	return b.X
}

// AssignY sets the y position of the element.
func (b *Basic) AssignY(y float64) {
	b.Y = y
}

// GetY returns the y position of the element.
func (b *Basic) GetY() float64 {
	return b.Y
}

// AssignOriginX sets the origin x position of the element.
func (b *Basic) AssignOriginX(x float64) {
	b.OriginX = x
}

// AssignOriginY sets the origin y position of the element.
func (b *Basic) AssignOriginY(y float64) {
	b.OriginY = y
}

// AssignDisabled sets the disabled state of the element.
func (b *Basic) AssignDisabled(v bool) {
	b.Disabled = v
}

// GetDisabled returns the disabled state of the element.
func (b *Basic) GetDisabled() bool {
	return b.Disabled
}
