package parser

import (
	"fmt"

	"github.com/madhab452/collection/filter/lexer"
	"github.com/madhab452/collection/filter/token"
)

type AndStatement struct {
	Field    token.Token
	Operator token.Token
	Value    interface{}
}

type ParsedFilter struct {
	Statements []*AndStatement
}

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token

	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// set both peek and next token
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseFilter() *ParsedFilter {
	filter := &ParsedFilter{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStmt()
		if stmt != nil {
			filter.Statements = append(filter.Statements, stmt)
		}

		p.nextToken()
	}

	return filter
}

func (p *Parser) parseStmt() *AndStatement {
	andStmt := &AndStatement{
		Field: p.curToken,
	}

	p.nextToken()

	switch p.curToken.Type {
	case token.EQ, token.GT, token.LT:
		andStmt.Operator = p.curToken
	default:
		p.errors = append(p.errors, fmt.Sprintf("invalid operator: got %v", p.curToken.Literal))
		return nil
	}

	p.nextToken()
	andStmt.Value = p.curToken.Literal

	if p.peekToken.Type == token.AND {
		p.nextToken()
	}

	return andStmt
}
