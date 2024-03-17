package parser

import (
	"testing"

	"github.com/madhab452/collection/filter/lexer"
)

func TestParse(t *testing.T) {
	input := `field_123 = 10 and fieldy > 10 and field.z != 0`

	p := New(lexer.New(input))
	pf := p.ParseFilter()

	if pf == nil {
		t.Fatalf("ParseFilter() returned nil")
	}
	expectedStmtLen := 3
	if len(pf.Statements) != expectedStmtLen {
		t.Errorf("pf.AndStatements does not contain %d statements. got=%d", expectedStmtLen, len(pf.Statements))
		t.Fatalf("got: %+v", pf.Statements[0])
	}
}
