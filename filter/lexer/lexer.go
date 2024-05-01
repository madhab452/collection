package lexer

import "github.com/madhab452/collection/filter/token"

type Lexer struct {
	input        string
	position     int // current position in input. (points to current character)
	readPosition int // current reading position in input. (after current char)
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) peekChar() byte {
	if l.readPosition > len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '!':
		if peekChar := l.peekChar(); peekChar == '=' {
			tok = newToken(token.NOT_EQ, string(l.ch)+string(peekChar))
			l.readChar() // advance the cursor as we already read the next char
		} else {
			tok = newToken(token.BANG, string(l.ch))
		}
	case '=':
		tok = newToken(token.EQ, string(l.ch))
	case '<':
		tok = newToken(token.LT, string(l.ch))
	case '>':
		tok = newToken(token.GT, string(l.ch))
	case ',':
		tok = newToken(token.COMMA, string(l.ch))
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		switch {
		case isLetter(l.ch):
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		case isDigit(l.ch):
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		case string(l.ch) == string(token.DQUOTE):
			tok.Type = token.STRING
			tok.Literal = l.readString()
			return tok
		}
	}

	l.readChar()

	return tok
}

func newToken(tokType token.TokenType, val string) token.Token {
	return token.Token{
		Type:    tokType,
		Literal: string(val),
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isAlnum(l.ch) { // alpha numeric
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	if ch >= 'a' && ch <= 'z' ||
		ch >= 'A' && ch <= 'Z' ||
		ch == '_' ||
		ch == '.' {
		return true
	}

	return false
}
func isAlnum(ch byte) bool {
	if isLetter(ch) || isDigit(ch) {
		return true
	}

	return false
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readString() string {
	l.readChar() // skip two quotes start and end
	defer l.readChar()
	pos := l.position

	for l.ch != '"' {
		l.readChar()
	}
	return l.input[pos:l.position]
}
