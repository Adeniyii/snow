package ast

import "snow/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement implements the Node interface and the Statement interface.
// It is the AST Node representation of a let statement.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// ReturnStatement implements the Node interface and the Statement interface.
// It is the AST Node representation of a return statement.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

type Identifier struct {
	Token token.Token
}

func (i *Identifier) statementNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
