package parser

import (
	"testing"

	"github.com/madhab452/collection/filter/lexer"
)

func TestParse(t *testing.T) {
	input := `field_123 = 10 and fieldy > 10 and field.z != 0`

	p := New(lexer.New(input))
	pf := p.Parse()

	if pf == nil {
		t.Fatalf("Parse() returned nil")
	}
	expectedStmtLen := 3
	if len(pf.AndStatements) != expectedStmtLen {
		t.Errorf("pf.AndStatements does not contain %d statements. got=%d", expectedStmtLen, len(pf.AndStatements))
		t.Fatalf("got: %+v", pf.AndStatements[0])
	}
}
