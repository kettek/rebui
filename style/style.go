package style

// Alignment is used to determine how text, images, or otherwise are aligned.
type Alignment string

// Our various alignments.
const (
	Left   Alignment = "left"
	Center Alignment = "center"
	Right  Alignment = "right"
	Top    Alignment = "top"
	Middle Alignment = "middle"
	Bottom Alignment = "bottom"
)

// Wrap is used to determine how text is word wrapped.
type Wrap string

// Our various wrapping.
const (
	NoWrap Wrap = "nowrap"
	Word   Wrap = "word"
	Rune   Wrap = "rune"
)

// ImageStretch is used to determine how images stretch within their element.
type ImageStretch string

// Our various stretching.
const (
	// None doesn't apply any stretching.
	None ImageStretch = "none"
	// Fill will stretch the image to fit the element.
	Fill ImageStretch = "fill"
	// Cover will fit the image within the element, maintaining aspect ratio.
	Cover ImageStretch = "cover"
	// Nearest works like Cover, but to nearest whole multiple.
	Nearest ImageStretch = "nearest"
)
