// Package parser implements a parser for Snow source files. Input must be
// provided as a list of tokens; the output is an abstract syntax tree (AST)
// representing the Snow source. The parser is invoked through the `New` function.
package parser

import (
	"fmt"

	"snow/ast"
	"snow/lexer"
	"snow/token"
)

// The parser structure holds the parser's internal state.
// currToken holds the token currently being parsed,
// nextToken holds the next token to be parsed.
type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	nextToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.errors = []string{}
	p.readToken()
	p.readToken()
	return p
}

// readToken advances the currToken and nextToken fields of the Parser
// using the NextToken method of the Lexer.
func (p *Parser) readToken() {
	p.currToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.readToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currToken}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.currTokenIs(token.SEMICOLON) {
		p.readToken()
	}

	return stmt
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.nextTokenIs(t) {
		p.readToken()
		return true
	}
	p.pushError(t)
	return false
}

func (p *Parser) nextTokenIs(t token.TokenType) bool {
	return p.nextToken.Type == t
}

func (p *Parser) currTokenIs(t token.TokenType) bool {
	return p.currToken.Type == t
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) pushError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.nextToken.Type)
	p.errors = append(p.errors, msg)
}
