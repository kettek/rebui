package rebui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// BasicElement provides the core functionality for positioning and testing for hits.
type BasicElement struct {
	X, Y, Width, Height float64
	OriginX, OriginY    float64
}

// Hit returns true if the given x and y coordinates are within the bounds of the element.
func (b *BasicElement) Hit(x, y float64) bool {
	return x >= b.X && x <= b.X+b.Width && y >= b.Y && y <= b.Y+b.Height
}

// SetWidth sets the width of the element.
func (b *BasicElement) SetWidth(w float64) {
	b.Width = w
}

// SetHeight sets the height of the element.
func (b *BasicElement) SetHeight(h float64) {
	b.Height = h
}

// SetX sets the x position of the element.
func (b *BasicElement) SetX(x float64) {
	b.X = x
}

// SetY sets the y position of the element.
func (b *BasicElement) SetY(y float64) {
	b.Y = y
}

// SetOriginX sets the origin x position of the element.
func (b *BasicElement) SetOriginX(x float64) {
	b.OriginX = x
}

// SetOriginY sets the origin y position of the element.
func (b *BasicElement) SetOriginY(y float64) {
	b.OriginY = y
}

// Element is the interface that all elements must implement.
type Element interface {
	Draw(*ebiten.Image)
}

// BackgroundColorSetter is used to set the background color of the given element.
type BackgroundColorSetter interface {
	SetBackgroundColor(color.Color)
}

// ForegroundColorSetter is used to set the foreground color of the given element.
type ForegroundColorSetter interface {
	SetForegroundColor(color.Color)
}

// BorderColorSetter is used to set the border color of the given element.
type BorderColorSetter interface {
	SetBorderColor(color.Color)
}

// VerticalAlignmentSetter is used to set the vertical alignment of the given element.
type VerticalAlignmentSetter interface {
	SetVerticalAlignment(Alignment)
}

// HorizontalAlignmentSetter is used to set the horizontal alignment of the given element.
type HorizontalAlignmentSetter interface {
	SetHorizontalAlignment(Alignment)
}

// TextSetter is used to set the text of the given element.
type TextSetter interface {
	SetText(string)
}

// FontFaceSetter is used to set the font face that the given element should use. This is generally derived from Theme, but may be overridden.
type FontFaceSetter interface {
	SetFontFace(text.Face)
}

// FontSizeSetter is used to set the font size of the given element.
type FontSizeSetter interface {
	SetFontSize(float64)
}

// ImageScaleSetter is used to set the image scale style of the given element.
type ImageScaleSetter interface {
	SetImageScale(ImageScale)
}

// ImageSetter is used to set the image of the given element.
type ImageSetter interface {
	SetImage(*ebiten.Image)
}

// XSetter is used to set the x position of the given element.
type XSetter interface {
	SetX(float64)
}

// YSetter is used to set the y position of the given element.
type YSetter interface {
	SetY(float64)
}

// OriginXSetter is used to set the origin x position of the given element.
type OriginXSetter interface {
	SetOriginX(float64)
}

// OriginYSetter is used to set the origin y position of the given element.
type OriginYSetter interface {
	SetOriginY(float64)
}

// WidthSetter is used to set the width of the given element.
type WidthSetter interface {
	SetWidth(float64)
}

// HeightSetter is used to set the height of the given element.
type HeightSetter interface {
	SetHeight(float64)
}

// PointerMoveReceiver is used to receive pointer move events.
type PointerMoveReceiver interface {
	HandlePointerMove(PointerMoveEvent)
}

// GlobalMoveReceiver is used to receive global pointer move events. Note that these move events will only be received if the element received a press event (there does not have to be a handler).
type GlobalMoveReceiver interface {
	HandlePointerGlobalMove(PointerMoveEvent)
}

// PointerInReceiver is used to receive pointer in events.
type PointerInReceiver interface {
	HandlePointerIn(PointerInEvent)
}

// PointerOutReceiver is used to receive pointer out events.
type PointerOutReceiver interface {
	HandlePointerOut(PointerOutEvent)
}

// PressReceiver is used to receive pointer press events. This occurs when a mouse button is pressed on the element.
type PressReceiver interface {
	HandlePointerPress(PointerPressEvent)
}

// ReleaseReceiver is used to receive pointer release events. This occurs when a mouse button is released on the element.
type ReleaseReceiver interface {
	HandlePointerRelease(PointerReleaseEvent)
}

// GlobalReleaseReceiver is used to receive global pointer release events. Note that these release events will only be received if the element received a press event (there does not have to be a handler).
type GlobalReleaseReceiver interface {
	HandlePointerGlobalRelease(PointerReleaseEvent)
}

// PressedReceiver is used to receive a press and subsequent release on the same element. This is akin to a "Click" event.
type PressedReceiver interface {
	HandlePointerPressed(PointerPressedEvent)
}

// HitChecker checks if the given coordinate hits the target element.
type HitChecker interface {
	Hit(x, y float64) bool
}
