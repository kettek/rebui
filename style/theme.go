package style

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

// Theme is a collection of defaults for how elements should be rendered.
type Theme struct {
	BackgroundColor color.Color
	ForegroundColor color.Color
	BorderColor     color.Color
	FontFace        text.Face
}

// NewTheme makes a new theme, wow.
func NewTheme() *Theme {
	return &Theme{
		BackgroundColor: color.RGBA{0, 0, 0, 0},
		ForegroundColor: color.RGBA{0, 0, 0, 0},
		BorderColor:     color.RGBA{0, 0, 0, 0},
		FontFace:        defaultFontFace,
	}
}

// DefaultTheme is the default fallback theme.
var DefaultTheme = NewTheme()

var globalTheme *Theme
var defaultFontFace text.Face

func SetGlobalTheme(theme *Theme) {
	globalTheme = theme
}

func CurrentTheme() *Theme {
	if globalTheme != nil {
		return globalTheme
	}
	return DefaultTheme
}

func init() {
	DefaultTheme.BackgroundColor = color.RGBA{96, 96, 96, 255}
	DefaultTheme.ForegroundColor = color.RGBA{200, 200, 200, 255}
	DefaultTheme.BorderColor = color.RGBA{150, 150, 150, 255}
}
