package filter

import (
	"errors"

	"github.com/madhab452/collection/filter/lexer"
	"github.com/madhab452/collection/filter/parser"
)

func Parse(in string) (*parser.Filter, error) {
	l := lexer.New(in)

	p := parser.New(l)
	if len(p.Errors()) > 0 {
		return nil, errors.New(p.Errors()[0])
	}

	return p.Parse(), nil
}
