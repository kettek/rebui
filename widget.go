package rebui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui/widgets/assigners"
	"github.com/kettek/rebui/widgets/getters"
	"github.com/kettek/rebui/widgets/receivers"
)

// Widget is the interface that all widgets must implement.
type Widget interface {
	Draw(*ebiten.Image, *ebiten.DrawImageOptions)
}

// LayoutWidget is an optional interface that widgets can implement to allow for changing their size.
type LayoutWidget interface {
	Widget
	Layout(width, height float64) (float64, float64)
}

// AssignerBackgroundColor is an alias.
type AssignerBackgroundColor = assigners.BackgroundColor

// AssignerForegroundColor is an alias.
type AssignerForegroundColor = assigners.ForegroundColor

// AssignerBorderColor is an alias.
type AssignerBorderColor = assigners.BorderColor

// AssignerVerticalAlignment is an alias.
type AssignerVerticalAlignment = assigners.VerticalAlignment

// AssignerHorizontalAlignment is an alias.
type AssignerHorizontalAlignment = assigners.HorizontalAlignment

// AssignerText is an alias.
type AssignerText = assigners.Text

// AssignerTextWrap is an alias.
type AssignerTextWrap = assigners.TextWrap

// AssignerFontFace is an alias.
type AssignerFontFace = assigners.FontFace

// AssignerFontSize is an alias.
type AssignerFontSize = assigners.FontSize

// AssignerImageStretch is an alias.
type AssignerImageStretch = assigners.ImageStretch

// AssignerImage is an alias.
type AssignerImage = assigners.Image

// GetterX is an alias.
type GetterX = getters.X

// AssignerX is an alias.
type AssignerX = assigners.X

// GetterY is an alias.
type GetterY = getters.Y

// AssignerY is an alias.
type AssignerY = assigners.Y

// AssignerOriginX is an alias.
type AssignerOriginX = assigners.OriginX

// AssignerOriginY is an alias.
type AssignerOriginY = assigners.OriginY

// AssignerWidth is an alias.
type AssignerWidth = assigners.Width

// AssignerHeight is an alias.
type AssignerHeight = assigners.Height

// AssignerDisable is an alias.
type AssignerDisable = assigners.Disable

// GetterDisabled is an alias.
type GetterDisabled = getters.Disabled

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
