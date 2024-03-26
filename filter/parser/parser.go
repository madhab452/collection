package parser

import (
	"github.com/madhab452/collection/filter/lexer"
	"github.com/madhab452/collection/filter/token"
)

type Operator token.TokenType

func getValidOps() []Operator {
	return []Operator{
		Operator(token.EQ),
		Operator(token.NOT_EQ),
		Operator(token.LT),
		Operator(token.GT),
		Operator(token.IN),
	}
}

type AndStatement struct {
	Field    string
	Operator Operator
	Value    *Value
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

func (p *Parser) Errors() []string {
	return p.errors
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
		if stmt == nil {
			return nil
		}

		filter.Statements = append(filter.Statements, stmt)

		p.nextToken()
	}

	return filter
}

func (p *Parser) expectPeek(peekToken token.Token, ops []Operator) bool {
	for _, opv := range ops {
		if opv == Operator(peekToken.Type) {
			return true
		}
	}
	p.errors = append(p.errors, "unexpected peek token %+v", peekToken.Literal)

	return false
}

func (p *Parser) parseStmt() *AndStatement {
	andStmt := &AndStatement{
		Field: p.curToken.Literal,
	}

	if !p.expectPeek(p.peekToken, getValidOps()) {
		return nil
	}
	p.nextToken()

	andStmt.Operator = Operator(p.curToken.Type)

	p.nextToken()

	v := p.parseValue()

	andStmt.Value = v

	if p.peekToken.Type == token.AND {
		p.nextToken()
	}

	return andStmt
}

func (p *Parser) parseValue() *Value {
	switch p.curToken.Type {
	case token.TRUE, token.FALSE:
		return newValue(TypeBool, p.curToken.Literal)
	case token.DQUOTE:
		sentence := p.readSentence()
		p.nextToken()
		for p.curToken.Type != token.DQUOTE {
			p.nextToken()
		}
		p.nextToken()
		return newValue(TypeString, sentence)
	}

	return nil
}

func (p *Parser) readSentence() string {
	return "todo: read sentence"
}
