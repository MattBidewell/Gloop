package token

type TokenType string

type Token struct {
	Type		TokenType
	Literal	string
}


const (
	ILLEGAL = "ILLEGAL"
	EOF			= "EOF"

	// Identifiers + literals
	IDENT	= "IDENT" // add, x, y, etc
	INT		= "INT"

	// operators
	ASSIGN 	= "="
	PLUS		=	"+"

	// delimiters
	COMMA				= ","
	SEMICOLON		= ";"

	LPAREN 	= "("
	RPAREN 	= ")"
	LBRACE 	= "{"
	RBRACE 	= "}"

	FUNCTION	= "FUNCTION"
	LET				= "LET"
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}