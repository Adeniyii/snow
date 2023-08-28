package lexer

import (
	"testing"

	"snow/token"
)

func TestNextToken(t *testing.T) {
	input := "========"
	l := New(input)

	tests := []token.Token{
		{Literal: "=", Type: token.ASSIGN},
	}

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Literal != tt.Literal {
			t.Fatalf("tests[%d] - literal wrong.\n\t\t-> Expected [%v] Got [%v]", i, tt.Literal, tok.Literal)
		}

		if tok.Type != tt.Type {
			t.Fatalf("tests[%d] - type wrong.\n\t\t-> Expected [%v] Got [%v]", i, tt.Type, tok.Type)
		}
	}
}
