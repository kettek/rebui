package rebui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui/widgets/getters"
	"github.com/kettek/rebui/widgets/receivers"
	"github.com/kettek/rebui/widgets/setters"
)

// Widget is the interface that all widgets must implement.
type Widget interface {
	Draw(*ebiten.Image, *ebiten.DrawImageOptions)
}

// SetterBackgroundColor is an alias.
type SetterBackgroundColor = setters.BackgroundColor

// SetterForegroundColor is an alias.
type SetterForegroundColor = setters.ForegroundColor

// SetterBorderColor is an alias.
type SetterBorderColor = setters.BorderColor

// SetterVerticalAlignment is an alias.
type SetterVerticalAlignment = setters.VerticalAlignment

// SetterHorizontalAlignment is an alias.
type SetterHorizontalAlignment = setters.HorizontalAlignment

// SetterText is an alias.
type SetterText = setters.Text

// SetterTextWrap is an alias.
type SetterTextWrap = setters.TextWrap

// SetterFontFace is an alias.
type SetterFontFace = setters.FontFace

// SetterFontSize is an alias.
type SetterFontSize = setters.FontSize

// SetterImageScale is an alias.
type SetterImageScale = setters.ImageScale

// SetterImage is an alias.
type SetterImage = setters.Image

// GetterX is an alias.
type GetterX = getters.X

// SetterX is an alias.
type SetterX = setters.X

// GetterY is an alias.
type GetterY = getters.Y

// SetterY is an alias.
type SetterY = setters.Y

// SetterOriginX is an alias.
type SetterOriginX = setters.OriginX

// SetterOriginY is an alias.
type SetterOriginY = setters.OriginY

// SetterWidth is an alias.
type SetterWidth = setters.Width

// SetterHeight is an alias.
type SetterHeight = setters.Height

// ReceiverPointerMove is an alias.
type ReceiverPointerMove = receivers.PointerMove

// ReceiverPointerGlobalMove is an alias.
type ReceiverPointerGlobalMove = receivers.PointerGlobalMove

// ReceiverPointerIn is an alias.
type ReceiverPointerIn = receivers.PointerIn

// ReceiverPointerOut is an alias.
type ReceiverPointerOut = receivers.PointerOut

// ReceiverPointerPress is an alias.
type ReceiverPointerPress = receivers.PointerPress

// ReceiverPointerRelease is an alias.
type ReceiverPointerRelease = receivers.PointerRelease

// ReceiverGlobalRelease is an alias.
type ReceiverGlobalRelease = receivers.PointerGlobalRelease

// ReceiverPointerPressed is an alias.
type ReceiverPointerPressed = receivers.PointerPressed

// HitChecker checks if the given coordinate hits the target element.
type HitChecker interface {
	Hit(x, y float64) bool
}
