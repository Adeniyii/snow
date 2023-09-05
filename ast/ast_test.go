package ast

import (
	"testing"

	"snow/token"
)

func TestString(t *testing.T) {
	test := "let foobar = barfoo;"
	program := Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "foobar"},
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "barfoo"},
				},
			},
		},
	}
	if program.String() != test {
		t.Errorf("program.String() wrong. got=%q. expected %q", program.String(), test)
	}
}
