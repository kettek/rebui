package widgets

// Basic provides the core functionality for positioning and testing for hits.
type Basic struct {
	X, Y, Width, Height float64
	OriginX, OriginY    float64
}

// Hit returns true if the given x and y coordinates are within the bounds of the element.
func (b *Basic) Hit(x, y float64) bool {
	return x >= b.X && x <= b.X+b.Width && y >= b.Y && y <= b.Y+b.Height
}

// SetWidth sets the width of the element.
func (b *Basic) SetWidth(w float64) {
	b.Width = w
}

// GetWidth returns the width of the element.
func (b *Basic) GetWidth() float64 {
	return b.Width
}

// SetHeight sets the height of the element.
func (b *Basic) SetHeight(h float64) {
	b.Height = h
}

// GetHeight returns the height of the element.
func (b *Basic) GetHeight() float64 {
	return b.Height
}

// SetX sets the x position of the element.
func (b *Basic) SetX(x float64) {
	b.X = x
}

// GetX returns the x position of the element.
func (b *Basic) GetX() float64 {
	return b.X
}

// SetY sets the y position of the element.
func (b *Basic) SetY(y float64) {
	b.Y = y
}

// GetY returns the y position of the element.
func (b *Basic) GetY() float64 {
	return b.Y
}

// SetOriginX sets the origin x position of the element.
func (b *Basic) SetOriginX(x float64) {
	b.OriginX = x
}

// SetOriginY sets the origin y position of the element.
func (b *Basic) SetOriginY(y float64) {
	b.OriginY = y
}
