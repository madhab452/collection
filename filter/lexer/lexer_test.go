package lexer

import (
	"testing"

	"github.com/madhab452/collection/filter/token"
)

func TestNextToken(t *testing.T) {
	input := `field_123 = 10 and fieldy > 10`

	testCases := []struct {
		Literal string
		Type    token.TokenType
	}{
		{
			Literal: "field_123",
			Type:    token.IDENT,
		},
		{
			Literal: "=",
			Type:    token.EQ,
		},
		{
			Literal: "10",
			Type:    token.INT,
		},
		{
			Literal: "and",
			Type:    token.AND,
		},
		{
			Literal: "fieldy",
			Type:    token.IDENT,
		},
		{
			Literal: ">",
			Type:    token.GT,
		},
		{
			Literal: "10",
			Type:    token.INT,
		},
	}
	l := New(input)

	for i, tc := range testCases {
		t.Run(tc.Literal, func(t *testing.T) {
			i++
			nt := l.NextToken()
			if nt.Type != tc.Type {
				t.Errorf("test case #%d failed, expected type %q, got %q", i, tc.Type, nt.Type)
			}
			if nt.Literal != tc.Literal {
				t.Errorf("test case #%d failed, expected literal %q, got %q", i, tc.Literal, nt.Literal)
			}
		})
	}
}
