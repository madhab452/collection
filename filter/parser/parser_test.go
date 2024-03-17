package parser

import (
	"testing"

	"github.com/madhab452/collection/filter/lexer"
)

func TestParse(t *testing.T) {

	input := `
    field1 = 100 and field2 > 10 
  `

	p := New(lexer.New(input))
	pf := p.ParseFilter()

	if pf == nil {
		t.Fatalf("ParseFilter() returned nil")
	}
	if len(pf.Statements) != 2 {
		t.Errorf("pf.AndStatements does not contain 2 statements. got=%d", len(pf.Statements))
		t.Fatalf("got: %+v", pf.Statements[0])
	}

}
