// Package lexer implements functions and objects for transforming
// source code from raw input into a data structure containing
// generated tokens.
package lexer

import (
	"snow/token"
)

// A Lexer represents the initial transformation factory of the source code.
// It holds the input and provides methods to traverse and generate tokens
// for each character.
type Lexer struct {
	input        string
	position     int
	nextPosition int
	ch           byte
}

// New initializes and returns a new lexer object.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar advances the lexer's current character, and updates the
// current position and next position by 1.
func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

// eatWhiteSpace advances the current character past all whitespace characters.
func (l *Lexer) eatWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

// NextToken reads the current character then constructs
// and outputs a corresponding token.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhiteSpace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, '=')
	case '-':
		tok = newToken(token.MINUS, '-')
	case '+':
		tok = newToken(token.PLUS, '+')
	case '*':
		tok = newToken(token.ASTERISK, '*')
	case '/':
		tok = newToken(token.SLASH, '/')
	case '<':
		tok = newToken(token.LT, '<')
	case '>':
		tok = newToken(token.GT, '>')
	case ';':
		tok = newToken(token.SEMICOLON, ';')
	case '.':
		tok = newToken(token.PERIOD, '.')
	case ',':
		tok = newToken(token.COMMA, ',')
	case '!':
		tok = newToken(token.BANG, '!')
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
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdent()
			tok.Type = token.LookupKeyword(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			// TODO: handle floats, hexadecimal...
			tok.Literal = l.readDigit()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, literal byte) token.Token {
	tok := &token.Token{Literal: string(literal), Type: tokenType}
	return *tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// readIdent identifies and returns a string sequence
func (l *Lexer) readIdent() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readDigit identifies and returns a number sequence
func (l *Lexer) readDigit() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
