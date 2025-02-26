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

// ImageScale is used to determine how images scale within their element.
type ImageScale string

// Our various scaling.
const (
	// None doesn't apply any scaling.
	None ImageScale = "none"
	// Fill will stretch the image to fit the element.
	Fill ImageScale = "fill"
	// Cover will fit the image within the element, maintaining aspect ratio.
	Cover ImageScale = "cover"
	// Nearest works like Cover, but to nearest whole multiple.
	Nearest ImageScale = "nearest"
)
