package token

// Update this to an int or a byte might be a good option for the future
type TokenType byte 

const (
	ILLEGAL TokenType = iota
	EOF
	
	IDENT
	INT
	FLOAT
	STRING

	ASSIGN
	PLUS
	DOUBLEPLUS
	PLUSEQUAL
	MINUS
	DOUBLEMINUS
	MINUSEQUAL
	BANG
	STAR
	DOUBLESTAR
	STAREQUAL
	SLASH
	DOUBLESLASH
	SLASHEQUAL
	PERCENT

	AND
	OR
	LT
	LTEQUAL
	GT
	GTEQUAL
	EQ
	NOTEQ

	COMMA
	SEMICOLON
	COLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	FUNCTION
	VAR
	CONST
	TRUE
	FALSE
	IF
	ELSE
	RETURN
	FOR
	BREAK
	CONTINUE
	NIL
	IMPORT  // Not sure if I will use this.
)


type Token struct {
	Type 	TokenType
	Literal string
	Line	int
	Column	int
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
	"break": BREAK,
	"continue": CONTINUE,
	"nil": NIL,
	"import": IMPORT,
}

func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT 
}


// This is gonna be used to make the process of debugging easier.
var tokenTypeStrings = [...]string{
	"ILLEGAL",
	"EOF",
	"IDENT",
	"INT",
	"FLOAT",
	"STRING",
	"=",
	"+",
	"++",
	"+=",
	"-",
	"--",
	"-=",
	"!",
	"*",
	"**",
	"*=",
	"/",
	"//",
	"/=",
	"%",
	"&&",
	"||",
	"<",
	"<=",
	">",
	">=",
	"==",
	"!=",
	",",
	";",
	":",
	"(",
	")",
	"{",
	"}",
	"FUNCTION",
	"VAR",
	"CONST",
	"TRUE",
	"FALSE",
	"IF",
	"ELSE",
	"RETURN",
	"FOR",
	"BREAK",
	"CONTINUE",
	"NIL",
	"IMPORT",
}

func (tt TokenType) String() string {
	if int(tt) < len(tokenTypeStrings) {
		return tokenTypeStrings[tt]
	}
	return "UNKNOWN"
}

