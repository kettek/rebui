package widgets

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kettek/rebui"
)

type Image struct {
	Basic
	Border
	scale  rebui.ImageStretch
	image  *ebiten.Image
	valign rebui.Alignment
	halign rebui.Alignment
}

func (w *Image) AssignImage(image *ebiten.Image) {
	w.image = image
}

func (w *Image) AssignImageStretch(scale rebui.ImageStretch) {
	w.scale = scale
}

func (w *Image) AssignVerticalAlignment(align rebui.Alignment) {
	w.valign = align
}

func (w *Image) AssignHorizontalAlignment(align rebui.Alignment) {
	w.halign = align
}

func (w *Image) Draw(screen *ebiten.Image, sop *ebiten.DrawImageOptions) {
	if w.image != nil {
		op := &ebiten.DrawImageOptions{}

		iw, ih := float64(w.image.Bounds().Dx()), float64(w.image.Bounds().Dy())
		sw, sh := float64(w.Width), float64(w.Height)

		switch w.scale {
		case rebui.ImageStretchFill:
			op.GeoM.Scale(sw/iw, sh/ih)
			iw *= sw / iw
			ih *= sh / ih
		case rebui.ImageStretchCover:
			if iw < ih {
				op.GeoM.Scale(sh/ih, sh/ih)
				iw *= sh / ih
				ih *= sh / ih
			} else {
				op.GeoM.Scale(sw/iw, sw/iw)
				iw *= sw / iw
				ih *= sw / iw
			}
		case rebui.ImageStretchNearest:
			scale := math.Floor(sw / iw)
			if sh/ih < scale {
				scale = math.Floor(sh / ih)
			}
			op.GeoM.Scale(scale, scale)
			iw *= scale
			ih *= scale
		}

		op.GeoM.Concat(sop.GeoM)

		if w.halign == rebui.AlignCenter {
			op.GeoM.Translate(w.Width/2, 0)
			op.GeoM.Translate(-float64(iw)/2, 0)
		} else if w.halign == rebui.AlignRight {
			op.GeoM.Translate(w.Width, 0)
			op.GeoM.Translate(-float64(iw), 0)
		}

		if w.valign == rebui.AlignMiddle {
			op.GeoM.Translate(0, w.Height/2)
			op.GeoM.Translate(0, -float64(ih)/2)
		} else if w.valign == rebui.AlignBottom {
			op.GeoM.Translate(0, w.Height)
			op.GeoM.Translate(0, -float64(ih))
		}

		screen.DrawImage(w.image, op)
	}

	if w.borderColor != nil {
		x := sop.GeoM.Element(0, 2)
		y := sop.GeoM.Element(1, 2)
		w.drawBorder(screen, float32(x), float32(y), float32(w.Width), float32(w.Height))
	}
}

func init() {
	rebui.RegisterWidget("Image", &Image{})
}
