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

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = "<"

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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
