package lexer

import (
	"testing"

	"snow/token"
)

func TestNextToken(t *testing.T) {
	input := "=+-/*{[(;.)]}"
	l := New(input)

	tests := []token.Token{
		{Literal: "=", Type: token.ASSIGN},
		{Literal: "+", Type: token.PLUS},
		{Literal: "-", Type: token.MINUS},
		{Literal: "/", Type: token.DIVIDE},
		{Literal: "*", Type: token.PRODUCT},
		{Literal: "{", Type: token.LBRACE},
		{Literal: "[", Type: token.LSQUARE},
		{Literal: "(", Type: token.LPAREN},
		{Literal: ";", Type: token.SEMICOLON},
		{Literal: ".", Type: token.PERIOD},
		{Literal: ")", Type: token.RPAREN},
		{Literal: "]", Type: token.RSQUARE},
		{Literal: "}", Type: token.RBRACE},
		{Literal: "", Type: token.EOF},
	}

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Literal != tt.Literal {
			t.Fatalf("tests[%d] - literal wrong.\n\t\t-> Expected [%q] Got [%q]", i, tt.Literal, tok.Literal)
		}

		if tok.Type != tt.Type {
			t.Fatalf("tests[%d] - type wrong.\n\t\t-> Expected [%q] Got [%q]", i, tt.Type, tok.Type)
		}
	}
}
