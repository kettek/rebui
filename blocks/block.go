package blocks

// Block is our block interface, wow.
type Block interface{}

// GetSize gets the width and height of the blocks.
func GetSize(blocks []Block, cfg Config) (width, height float64) {
	lineH := cfg.Face.Metrics().HAscent + cfg.Face.Metrics().HDescent
	curWidth := 0.0
	for _, block := range blocks {
		switch b := block.(type) {
		case Text:
			curWidth += b.Width
		case Break:
			height += lineH
			if curWidth > width {
				width = curWidth
			}
			curWidth = 0
		}
	}
	height += lineH // Hmm, this seems wrong.
	return
}
