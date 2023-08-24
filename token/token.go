package token

// type alias `TokenType` to a string
// An int or Byte would be more performant
type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
