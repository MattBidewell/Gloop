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
	MINUS		=	"-"
	BANG		= "!"
	ASTRIX	= "*"
	SLASH		= "/"

	EQ 			= "=="
	NOT_EQ	= "!="
	LT 			= "<"
	GT 			= ">"
	GT_EQ		=	">="
	LT_EQ		=	"<="

	// delimiters
	COMMA				= ","
	SEMICOLON		= ";"

	LPAREN 	= "("
	RPAREN 	= ")"
	LBRACE 	= "{"
	RBRACE 	= "}"

	// keywords
	FUNCTION	= "FUNCTION"
	LET				= "LET"
	TRUE			= "TRUE"
	FALSE			= "FALSE"
	IF				= "IF"
	ELSE			= "ELSE"
	RETURN		= "RETURN"
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}