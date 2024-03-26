package token

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"
	// Identifiers + literals
	IDENT TokenType = "IDENT" // fieldx, fieldy
	INT   TokenType = "INT"

	// Operators
	EQ     TokenType = "="
	NOT_EQ TokenType = "!=" // todo: not suported yet
	LT     TokenType = "<"
	GT     TokenType = ">"
	COMMA  TokenType = ","
	BANG   TokenType = "!"
	DQUOTE TokenType = "\""

	// Delimiters
	LPAREN TokenType = "("
	RPAREN TokenType = ")"

	// Keywords
	AND   TokenType = "AND"
	OR    TokenType = "OR"
	IN    TokenType = "IN"
	TRUE  TokenType = "TRUE"
	FALSE TokenType = "FALSE"
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
