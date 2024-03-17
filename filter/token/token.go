package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + literals
	IDENT = "IDENT" // fieldx, fieldy
	INT   = "INT"

	// Operators
	EQ     = "="
	NOT_EQ = "!=" // todo: not suported yet
	LT     = "<"
	GT     = ">"
	COMMA  = ","

	// Delimiters
	LPAREN = "("
	RPAREN = ")"

	// Keywords
	AND = "AND"
	OR  = "OR"
	IN  = "IN"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"and": AND,
	"or":  OR,
	"in":  IN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
