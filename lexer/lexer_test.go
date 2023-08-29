package lexer

import (
	"testing"

	"snow/token"
)

func TestNextToken(t *testing.T) {
	input := `
let foo = 6;
let bar = 9;

let add = fn(x, y) {
  return x + y;
};
let result = add(foo, bar);
  `
	l := New(input)

	tests := []token.Token{
		{Literal: "let", Type: token.LET},
		{Literal: "foo", Type: token.IDENT},
		{Literal: "=", Type: token.ASSIGN},
		{Literal: "6", Type: token.INT},
		{Literal: ";", Type: token.SEMICOLON},
		{Literal: "let", Type: token.LET},
		{Literal: "bar", Type: token.IDENT},
		{Literal: "=", Type: token.ASSIGN},
		{Literal: "9", Type: token.INT},
		{Literal: ";", Type: token.SEMICOLON},
		{Literal: "let", Type: token.LET},
		{Literal: "add", Type: token.IDENT},
		{Literal: "=", Type: token.ASSIGN},
		{Literal: "fn", Type: token.FUNCTION},
		{Literal: "(", Type: token.LPAREN},
		{Literal: "x", Type: token.IDENT},
		{Literal: ",", Type: token.COMMA},
		{Literal: "y", Type: token.IDENT},
		{Literal: ")", Type: token.RPAREN},
		{Literal: "{", Type: token.LBRACE},
		{Literal: "return", Type: token.RETURN},
		{Literal: "x", Type: token.IDENT},
		{Literal: "+", Type: token.PLUS},
		{Literal: "y", Type: token.IDENT},
		{Literal: ";", Type: token.SEMICOLON},
		{Literal: "}", Type: token.RBRACE},
		{Literal: ";", Type: token.SEMICOLON},
		{Literal: "let", Type: token.LET},
		{Literal: "result", Type: token.IDENT},
		{Literal: "=", Type: token.ASSIGN},
		{Literal: "add", Type: token.IDENT},
		{Literal: "(", Type: token.LPAREN},
		{Literal: "foo", Type: token.IDENT},
		{Literal: ",", Type: token.COMMA},
		{Literal: "bar", Type: token.IDENT},
		{Literal: ")", Type: token.RPAREN},
		{Literal: ";", Type: token.SEMICOLON},
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
