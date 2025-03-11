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

// Wrap is a type alias for style.Wrap.
type Wrap = style.Wrap

// Our wrapping types. See style package for more info.
const (
	WrapNone = style.NoWrap
	WrapWord = style.Word
	WrapRune = style.Rune
)

// ImageStretch is a type alias for style.ImageStretch.
type ImageStretch = style.ImageStretch

// Our image scaling types. See style package for more info.
const (
	ImageStretchNone    = style.None
	ImageStretchFill    = style.Fill
	ImageStretchCover   = style.Cover
	ImageStretchNearest = style.Nearest
)
