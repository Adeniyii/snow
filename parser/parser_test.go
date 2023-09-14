package parser

import (
	"testing"

	"snow/ast"
	"snow/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let foo = 6;
let bar = 9;
let foobar = 69;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil.")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got %d", len(program.Statements))
	}
	tests := []struct {
		expectedIdentifier string
	}{
		{"foo"},
		{"bar"},
		{"foobar"},
	}
	for i, tt := range tests {
		stmnt := program.Statements[i]
		if !testLetStatement(t, stmnt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
return 6;
return 9;
return 6 * 9;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil.")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got %d", len(program.Statements))
	}

	for _, stmnt := range program.Statements {
		retStmnt, ok := stmnt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmnt not ast.ReturnStatement. got=%T", stmnt)
			continue
		}
		if retStmnt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				retStmnt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has invalid statement count. got=%d",
			len(program.Statements))
	}

	stmnt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
			program.Statements[0])
	}

	ident, ok := stmnt.Expression.(*ast.Identifier)

	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmnt.Expression)
	}
	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral not %s. got=%s", "foobar",
			ident.TokenLiteral())
	}
}

func testLetStatement(t *testing.T, stmnt ast.Statement, name string) bool {
	if stmnt.TokenLiteral() != "let" {
		t.Fatalf("stmnt.TokenLiteral() not 'let'. got %q", stmnt.TokenLiteral())
		return false
	}
	letStmnt, ok := stmnt.(*ast.LetStatement)
	if !ok {
		t.Fatalf("stmnt does not implement the LetStatement interface. got %T", stmnt)
		return false
	}
	if letStmnt.Name.Token.Literal != name {
		t.Fatalf("letStmnt.Name.Token.Literal not %s. got %s", name, letStmnt.Name.Token.Literal)
		return false
	}
	if letStmnt.Name.TokenLiteral() != name {
		t.Fatalf("letStmnt.Name.TokenLiteral() not %s. got %s", name, letStmnt.Name.TokenLiteral())
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors.", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error %q", msg)
	}
	t.FailNow()
}
