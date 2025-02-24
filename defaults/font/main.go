package font

import (
	"bytes"

	// Linter needs to stfu
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kettek/rebui"
)

//go:embed x12y12pxMaruMinyaM.ttf
var fontBytes []byte

func init() {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fontBytes))
	if err != nil {
		panic(err)
	}
	rebui.DefaultTheme.FontFace = &text.GoTextFace{
		Source: s,
		Size:   12,
	}
}
