package lexer

import (
	"snow/token"
)

type Lexer struct {
	input        string
	position     int
	nextPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, '=')
	case '-':
		tok = newToken(token.MINUS, '-')
	case '+':
		tok = newToken(token.PLUS, '+')
	case '*':
		tok = newToken(token.PRODUCT, '*')
	case '/':
		tok = newToken(token.DIVIDE, '/')
	case ';':
		tok = newToken(token.SEMICOLON, ';')
	case '.':
		tok = newToken(token.PERIOD, '.')
	case '{':
		tok = newToken(token.LBRACE, '{')
	case '}':
		tok = newToken(token.RBRACE, '}')
	case '[':
		tok = newToken(token.LSQUARE, '[')
	case ']':
		tok = newToken(token.RSQUARE, ']')
	case '(':
		tok = newToken(token.LPAREN, '(')
	case ')':
		tok = newToken(token.RPAREN, ')')
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, literal byte) token.Token {
	tok := &token.Token{Literal: string(literal), Type: tokenType}
	return *tok
}
