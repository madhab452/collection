package filter

import (
	"github.com/madhab452/collection/filter/lexer"
	"github.com/madhab452/collection/filter/parser"
)

func Parse(in string) *parser.Parser {
	lexer := lexer.New(in)

	return parser.New(lexer)
}
