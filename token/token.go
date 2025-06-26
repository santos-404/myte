package token

// Update this to an int or a byte might be a good option for the future
type TokenType string

const (
	ILLEGAL 	= "ILLEGAL"
	EOF 		= "EOF"
	
	IDENT 		= "IDENT"
	INT 		= "INT"

	ASSIGN   	= "="
	PLUS     	= "+"
	MINUS    	= "-"
	BANG 		= "!"
	ASTERISK 	= "*"
	SLASH    	= "/"

	LT 			= "<"
	GT 			= ">"

	EQ     		= "=="
	NOT_EQ 		= "!="

	COMMA 		= ","
	SEMICOLON 	= ";"

	LPAREN 		= "("
	RPAREN 		= ")"
	LBRACE 		= "{"
	RBRACE 		= "}"

	FUNCTION 	= "FUNCTION"
	VAR      	= "VAR"
	CONST 		= "CONST"  // I ain't sure I will be able to use this.
	TRUE     	= "TRUE"
	FALSE    	= "FALSE"
	IF       	= "IF"
	ELSE     	= "ELSE"
	RETURN   	= "RETURN"
	FOR 		= "FOR"
)

type Token struct {
	Type 	TokenType
	Literal string
}

// This is useful to tell user-defined indetifiers apart from language keywords
var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"const": CONST,
	"var": VAR,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
	"for": FOR,
}

func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT 
}
