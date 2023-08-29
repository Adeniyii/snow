package token

type TokenType string

// Token is the final representation
// of each character of our source code after transformation by the lexer
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // identifies unknown characters
	EOF     = "EOF"     // signifies the end of file

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	DIVIDE  = "/"
	PRODUCT = "*"

	COMMA     = ","
	SEMICOLON = ";"
	PERIOD    = "."

	LPAREN  = "("
	RPAREN  = ")"
	LBRACE  = "{"
	RBRACE  = "}"
	LSQUARE = "["
	RSQUARE = "]"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
