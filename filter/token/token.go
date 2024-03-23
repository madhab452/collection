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
	BANG   = "!"
	DQUOTE = "\""

	// Delimiters
	LPAREN = "("
	RPAREN = ")"

	// Keywords
	AND   = "AND"
	OR    = "OR"
	IN    = "IN"
	TRUE  = "TRUE"
	FALSE = "FALSE"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"and":   AND,
	"or":    OR,
	"in":    IN,
	"true":  TRUE,
	"false": FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
