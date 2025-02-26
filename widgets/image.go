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

func (el *Image) SetImage(image *ebiten.Image) {
	el.image = image
}

func (el *Image) SetImageScale(scale rebui.ImageScale) {
	el.scale = scale
}

func (el *Image) SetVerticalAlignment(align rebui.Alignment) {
	el.valign = align
}

func (el *Image) SetHorizontalAlignment(align rebui.Alignment) {
	el.halign = align
}

func (el *Image) SetBorderColor(clr color.Color) {
	el.borderColor = clr
}

func (el *Image) Draw(screen *ebiten.Image) {
	if el.image != nil {
		op := &ebiten.DrawImageOptions{}

		iw, ih := float64(el.image.Bounds().Dx()), float64(el.image.Bounds().Dy())
		sw, sh := float64(el.Width), float64(el.Height)

		switch el.scale {
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

		if el.halign == rebui.AlignCenter {
			op.GeoM.Translate(el.X+el.Width/2, 0)
			op.GeoM.Translate(-float64(iw)/2, 0)
		} else if el.halign == rebui.AlignRight {
			op.GeoM.Translate(el.X+el.Width, 0)
			op.GeoM.Translate(-float64(iw), 0)
		} else {
			op.GeoM.Translate(el.X, 0)
		}

		if el.valign == rebui.AlignMiddle {
			op.GeoM.Translate(0, el.Y+el.Height/2)
			op.GeoM.Translate(0, -float64(ih)/2)
		} else if el.valign == rebui.AlignBottom {
			op.GeoM.Translate(0, el.Y+el.Height)
			op.GeoM.Translate(0, -float64(ih))
		} else {
			op.GeoM.Translate(0, el.Y)
		}

		screen.DrawImage(el.image, op)
	}

	if el.borderColor != nil {
		vector.StrokeRect(screen, float32(el.X), float32(el.Y), float32(el.Width), float32(el.Height), 1, el.borderColor, false)
	}
}

func init() {
	rebui.RegisterElement("Image", &Image{})
}
