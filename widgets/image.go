package widgets

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/kettek/rebui"
)

type Image struct {
	Basic
	scale       rebui.ImageScale
	image       *ebiten.Image
	valign      rebui.Alignment
	halign      rebui.Alignment
	borderColor color.Color
}

func (w *Image) SetImage(image *ebiten.Image) {
	w.image = image
}

func (w *Image) SetImageScale(scale rebui.ImageScale) {
	w.scale = scale
}

func (w *Image) SetVerticalAlignment(align rebui.Alignment) {
	w.valign = align
}

func (w *Image) SetHorizontalAlignment(align rebui.Alignment) {
	w.halign = align
}

func (w *Image) SetBorderColor(clr color.Color) {
	w.borderColor = clr
}

func (w *Image) Draw(screen *ebiten.Image) {
	if w.image != nil {
		op := &ebiten.DrawImageOptions{}

		iw, ih := float64(w.image.Bounds().Dx()), float64(w.image.Bounds().Dy())
		sw, sh := float64(w.Width), float64(w.Height)

		switch w.scale {
		case rebui.ImageScaleFill:
			op.GeoM.Scale(sw/iw, sh/ih)
			iw *= sw / iw
			ih *= sh / ih
		case rebui.ImageScaleCover:
			if iw < ih {
				op.GeoM.Scale(sh/ih, sh/ih)
				iw *= sh / ih
				ih *= sh / ih
			} else {
				op.GeoM.Scale(sw/iw, sw/iw)
				iw *= sw / iw
				ih *= sw / iw
			}
		case rebui.ImageScaleNearest:
			scale := math.Floor(sw / iw)
			if sh/ih < scale {
				scale = math.Floor(sh / ih)
			}
			op.GeoM.Scale(scale, scale)
			iw *= scale
			ih *= scale
		}

		if w.halign == rebui.AlignCenter {
			op.GeoM.Translate(w.X+w.Width/2, 0)
			op.GeoM.Translate(-float64(iw)/2, 0)
		} else if w.halign == rebui.AlignRight {
			op.GeoM.Translate(w.X+w.Width, 0)
			op.GeoM.Translate(-float64(iw), 0)
		} else {
			op.GeoM.Translate(w.X, 0)
		}

		if w.valign == rebui.AlignMiddle {
			op.GeoM.Translate(0, w.Y+w.Height/2)
			op.GeoM.Translate(0, -float64(ih)/2)
		} else if w.valign == rebui.AlignBottom {
			op.GeoM.Translate(0, w.Y+w.Height)
			op.GeoM.Translate(0, -float64(ih))
		} else {
			op.GeoM.Translate(0, w.Y)
		}

		screen.DrawImage(w.image, op)
	}

	if w.borderColor != nil {
		vector.StrokeRect(screen, float32(w.X), float32(w.Y), float32(w.Width), float32(w.Height), 1, w.borderColor, false)
	}
}

func init() {
	rebui.RegisterWidget("Image", &Image{})
}
