/*
Package token defines the token type and all the different types and keywords
supported by the Monkey language.
*/
package token

// TokenType represents a generic token
type TokenType string

// Token is a struct that has the field Type that will be used to categorize
// what the Literal is.
type Token struct {
	Type    TokenType
	Literal string
}

// Define keywords
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent returns the token for either a keyword or an identifier
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

// Defining the various token types.  Update Lexer.nextToken() in lexer/lexer.go
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // end of file

	// identifier or variable: foobar, x, y, etc.
	IDENT = "IDENT"

	// primitive types
	INT = "INT"

	// operators
	ASSIGN = "="
	PLUS   = "+"

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET" // declares identifiers
)
