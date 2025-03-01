package blocks

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/kettek/rebui"
)

// Text is a block of text.
type Text struct {
	Width  float64
	Height float64
	Text   string
}

func FromText(txt string, cfg Config) []Block {
	var blocks []Block

	var currentBlock Text
	var currentWidth float64
	for _, r := range txt {
		if r == '\n' {
			currentBlock.Width, currentBlock.Height = text.Measure(currentBlock.Text, cfg.Face, 0)
			blocks = append(blocks, currentBlock)
			blocks = append(blocks, Break{})
			currentBlock = Text{}
			currentWidth = 0
			continue
		}
		currentBlock.Text += string(r)
		if cfg.Wrap == rebui.WrapNone {
			continue
		}
		width, _ := text.Measure(currentBlock.Text, cfg.Face, 0)
		if currentWidth+width >= cfg.Width {
			if cfg.Wrap == rebui.WrapWord {
				// Find our last previous space, if possible.
				didit := false
				for i := len(currentBlock.Text) - 1; i >= 0; i-- {
					if currentBlock.Text[i] == ' ' {
						txt := currentBlock.Text
						currentBlock.Width, currentBlock.Height = text.Measure(currentBlock.Text[:i], cfg.Face, 0)
						currentBlock.Text = txt[:i]
						blocks = append(blocks, currentBlock)
						blocks = append(blocks, Break{})
						currentBlock = Text{Text: txt[i+1:]}
						didit = true
						break
					}
				}
				// WrapRune when if we fail.
				if !didit {
					blocks, currentBlock = genToPreviousRune(cfg, blocks, currentBlock)
				}
			} else if cfg.Wrap == rebui.WrapRune {
				blocks, currentBlock = genToPreviousRune(cfg, blocks, currentBlock)
			}
		}
	}
	// Add last blockie.
	currentBlock.Width, currentBlock.Height = text.Measure(currentBlock.Text, cfg.Face, 0)
	blocks = append(blocks, currentBlock)

	return blocks
}

func genToPreviousRune(cfg Config, blocks []Block, currentBlock Text) ([]Block, Text) {
	for i := len(currentBlock.Text) - 1; i >= 0; i-- {
		width, height := text.Measure(currentBlock.Text[:i], cfg.Face, 0)
		if width < cfg.Width {
			txt := currentBlock.Text
			currentBlock.Width = width
			currentBlock.Height = height
			currentBlock.Text = txt[:i]
			blocks = append(blocks, currentBlock)
			blocks = append(blocks, Break{})
			currentBlock = Text{Text: txt[i:]}
			break
		}
	}
	return blocks, currentBlock
}
