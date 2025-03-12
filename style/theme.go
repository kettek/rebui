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

	ActiveBackgroundColor color.Color
	ActiveForegroundColor color.Color
	ActiveBorderColor     color.Color

	HoverBackgroundColor color.Color
	HoverForegroundColor color.Color
	HoverBorderColor     color.Color

	FontFace text.Face
}

// NewTheme makes a new theme, wow.
func NewTheme() *Theme {
	return &Theme{
		FontFace: defaultFontFace,
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

	DefaultTheme.ActiveBackgroundColor = color.RGBA{160, 160, 160, 255}
	DefaultTheme.ActiveForegroundColor = color.RGBA{255, 255, 255, 255}
	DefaultTheme.ActiveBorderColor = color.RGBA{255, 255, 255, 255}

	DefaultTheme.HoverBackgroundColor = color.RGBA{128, 128, 128, 255}
	DefaultTheme.HoverForegroundColor = color.RGBA{255, 255, 255, 255}
	DefaultTheme.HoverBorderColor = color.RGBA{200, 200, 200, 255}
}
