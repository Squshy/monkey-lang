package token

type TokenType string

type Token struct {
	Type TokenType
	// an int or a byte would be better for performance, but for simplicity
	// this is fine
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" // Something we don't know about
	EOF     = "EOF"

	// Identifies + literals
	IDENT = "IDENT" // add, foobar, x, y, etc
	INT   = "INT"   // 12345

	// Operands
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
