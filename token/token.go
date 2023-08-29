package token

type TokenType string

// Token is the final representation
// of each character of our source code after transformation by the lexer
type Token struct {
	Type    TokenType
	Literal string
}

var keywordMap = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"return": RETURN,
}

// LookupKeyword checks for a keyword match in keywordMap
// and return a corresponding token type if a match was found,
// otherwise it return an IDENT token type.
func LookupKeyword(literal string) TokenType {
	if tok, ok := keywordMap[literal]; ok {
		return tok
	}
	return IDENT
}

// token types
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
	RETURN   = "RETURN"
)
