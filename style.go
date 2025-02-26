package rebui

import "github.com/kettek/rebui/style"

// Alignment is a type alias for style.Alignment.
type Alignment = style.Alignment

// Our alignment types. See style package for more info.
const (
	AlignLeft   = style.Left
	AlignCenter = style.Center
	AlignRight  = style.Right
	AlignTop    = style.Top
	AlignMiddle = style.Middle
	AlignBottom = style.Bottom
)

// ImageScale is a type alias for style.ImageScale.
type ImageScale = style.ImageScale

// Our image scaling types. See style package for more info.
const (
	ImageScaleNone    = style.None
	ImageScaleFill    = style.Fill
	ImageScaleCover   = style.Cover
	ImageScaleNearest = style.Nearest
)
