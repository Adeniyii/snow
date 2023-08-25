package token

// type alias `TokenType` to a string
// An int or Byte would be more performant
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // identifies unknown characters
	EOF     = "EOF"     // signifies the end of file

	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	// special characters
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
