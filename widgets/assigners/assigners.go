package assigners

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kettek/rebui/style"
)

// BackgroundColor is used to set the background color of the given element.
type BackgroundColor interface {
	AssignBackgroundColor(color.Color)
}

// ForegroundColor is used to set the foreground color of the given element.
type ForegroundColor interface {
	AssignForegroundColor(color.Color)
}

// BorderColor is used to set the border color of the given element.
type BorderColor interface {
	AssignBorderColor(color.Color)
}

// VerticalAlignment is used to set the vertical alignment of the given element.
type VerticalAlignment interface {
	AssignVerticalAlignment(style.Alignment)
}

// HorizontalAlignment is used to set the horizontal alignment of the given element.
type HorizontalAlignment interface {
	AssignHorizontalAlignment(style.Alignment)
}

// Text is used to set the text of the given element.
type Text interface {
	AssignText(string)
}

// TextWrap is used to set the text wrap of the given element.
type TextWrap interface {
	AssignTextWrap(style.Wrap)
}

// FontFace is used to set the font face that the given element should use. This is generally derived from Theme, but may be overridden.
type FontFace interface {
	AssignFontFace(text.Face)
}

// FontSize is used to set the font size of the given element.
type FontSize interface {
	AssignFontSize(float64)
}

// ImageScale is used to set the image scale style of the given element.
type ImageScale interface {
	AssignImageScale(style.ImageScale)
}

// Image is used to set the image of the given element.
type Image interface {
	AssignImage(*ebiten.Image)
}

// X is used to set the x position of the given element.
type X interface {
	AssignX(float64)
}

// Y is used to set the y position of the given element.
type Y interface {
	AssignY(float64)
}

// OriginX is used to set the origin x position of the given element.
type OriginX interface {
	AssignOriginX(float64)
}

// OriginY is used to set the origin y position of the given element.
type OriginY interface {
	AssignOriginY(float64)
}

// Width is used to set the width of the given element.
type Width interface {
	AssignWidth(float64)
}

// Height is used to set the height of the given element.
type Height interface {
	AssignHeight(float64)
}
