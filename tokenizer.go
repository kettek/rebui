package rebui

import (
	"github.com/bzick/tokenizer"
)

const (
	relationNone = iota
	relationAfter
	relationAt
	relationOf
)

const (
	unitPixels = iota
	unitPercentage
	unitVW
	unitVH
)

const (
	tUnit = iota + 1
	tRelation
)

var tokenParser *tokenizer.Tokenizer

func init() {
	tokenParser = tokenizer.New()
	tokenParser.DefineTokens(tRelation, []string{"after", "at", "of"})
	tokenParser.DefineTokens(tUnit, []string{"%", "vw", "vh"})
	tokenParser.AllowKeywordSymbols(tokenizer.Underscore, tokenizer.Numbers)
}
