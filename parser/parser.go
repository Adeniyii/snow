// Package parser implements a parser for Snow source files. Input must be
// provided as a list of tokens; the output is an abstract syntax tree (AST)
// representing the Snow source. The parser is invoked through the `New` function.
package parser

import (
	"fmt"
	"strconv"

	"snow/ast"
	"snow/lexer"
	"snow/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// precedence order assignment
const (
	_ int = iota
	LOWEST
	EQS    // ==
	LLGG   // > or <
	SUM    // +
	PROD   // *
	PREFIX // !X or -X
	CALL   // fn()
)

// The parser structure holds the parser's internal state.
// currToken holds the token currently being parsed,
// nextToken holds the next token to be parsed.
type Parser struct {
	l         *lexer.Lexer
	currToken token.Token
	nextToken token.Token
	errors    []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.errors = []string{}

	p.readToken()
	p.readToken()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)

	return p
}

// parseIdentifier is a parsing function for the IDENT token type.
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}
}

// parseIdentifier is a parsing function for the INT token type.
// It converts the token literal to an int type for accurate representation
// on the Value field.
func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.currToken}

	v, err := strconv.ParseInt(p.currToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currToken.Literal)
		p.errors = append(p.errors, msg)
	}

	lit.Value = v

	return lit
}

// ParseProgram sets up the program structures and kicks off parsing of the available tokens.
// It populates the statements slice with the statement node returned from each parsing operation.
// It does this until it encounteres an EOF token indicating the end of the program.
// Any errors encountered during parsing will be stored in the errors slice.
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

// parseStatement is a small engine that decides which parsing function to
// should handle parsing of the current statement/expression.
func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseLetStatement handles parsing of let statements.
// if the statement is invalid, an error is registered on the parser state
// and the function returns nil, otherwise it returns a LetStatement node.
func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.currToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currToken, Value: p.currToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	for !p.currTokenIs(token.SEMICOLON) {
		p.readToken()
	}

	return stmt
}

// parseReturnStatement handles parsing of return statements.
// if the statement is invalid, an error is registered on the parser state
// and the function returns nil, otherwise it returns a ReturnStatement node.
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currToken}

	p.readToken()

	for !p.currTokenIs(token.SEMICOLON) { // skipping over the expressions for now
		p.readToken()
	}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmnt := &ast.ExpressionStatement{Token: p.currToken}
	stmnt.Expression = p.parseExpression(LOWEST)

	if p.nextTokenIs(token.SEMICOLON) { // semicolons are optional
		p.readToken()
	}
	return stmnt
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.nextTokenIs(t) {
		p.readToken()
		return true
	}
	p.pushError(t)
	return false
}

// readToken advances the currToken and nextToken fields of the Parser
// using the NextToken method of the Lexer.
func (p *Parser) readToken() {
	p.currToken = p.nextToken
	p.nextToken = p.l.NextToken()
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

func (p *Parser) registerPrefix(t token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[t] = fn
}

func (p *Parser) registerInfix(t token.TokenType, fn infixParseFn) {
	p.infixParseFns[t] = fn
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currToken.Type]
	if prefix == nil {
		return nil
	}
	leftexp := prefix()
	return leftexp
}
