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

type Filter struct {
	AndStatements []*AndStatement
}

type AndStatement struct {
	Field    string
	Operator Operator
	Value    interface{}
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

func (p *Parser) Parse() *Filter {
	var filter = &Filter{}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStmt()
		if stmt == nil {
			return nil
		}

		filter.AndStatements = append(filter.AndStatements, stmt)

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

func (p *Parser) readSentence() string {
	return p.curToken.Literal
}

func (p *Parser) parseValue() interface{} {
	return p.curToken.Literal
}
