package rebui

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var imageLoader func(path string) (*ebiten.Image, error)

// SetImageLoader sets the function to load an image by path.
func SetImageLoader(loader func(path string) (*ebiten.Image, error)) {
	imageLoader = loader
}

// LoadImage loads the given path using the loader set in SetImageLoader.
func LoadImage(path string) (*ebiten.Image, error) {
	if imageLoader != nil {
		return imageLoader(path)
	}
	return nil, ErrNoImageLoader
}

var fontLoader func(name string) (text.Face, error)

// SetFontLoader sets the function to load a font by path.
func SetFontLoader(loader func(path string) (text.Face, error)) {
	fontLoader = loader
}

// LoadFont loads the given path using the loader set in SetFontLoader.
func LoadFont(path string) (text.Face, error) {
	if fontLoader != nil {
		return fontLoader(path)
	}
	return nil, ErrNoFontLoader
}

var templateLoader func(name string) (Nodes, error)

// SetTemplateLoader sets the function to load a template by path.
func SetTemplateLoader(loader func(path string) (Nodes, error)) {
	templateLoader = loader
}

// LoadTemplate loads the given path using the loader set in SetTemplateLoader
func LoadTemplate(path string) (Nodes, error) {
	if templateLoader != nil {
		return templateLoader(path)
	}
	return nil, ErrNoTemplateLoader
}

// Errors
var (
	ErrNoImageLoader    = errors.New("no image loader set")
	ErrNoFontLoader     = errors.New("no font loader set")
	ErrNoTemplateLoader = errors.New("no template loader set")
)
