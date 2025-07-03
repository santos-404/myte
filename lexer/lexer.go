package lexer

import (
	"github.com/santos-404/myte/token"
)

/*
Perhaps it seems like readPosition (as a second pointer) is not needed, it is.
The reason is the fact that we will need to be able to "peek" further into the input
and look after the current char to see what comes up next.
*/
type Lexer struct {
	input 			string
	position 		int  // current position on input | points to current char
	readPosition 	int  // current reading position | after current char
	char 			byte
	line			int
	column			int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()  // We can do this easily cause Go sets everything to "zero" when declaring. 
	return l
}

// In case you don't know go this is "similar" to a OOP method. 
// Behind the scenes it's just syntactic sugar.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
		case '=':
			if l.peekNextChar() == '=' {
				tok = l.newComplexToken(token.EQ)
			} else {
				tok = l.newToken(token.ASSIGN, l.char)
			}
		case '!':
			if l.peekNextChar() == '=' {
				tok = l.newComplexToken(token.NOTEQ)
			} else {
				tok = l.newToken(token.BANG, l.char)
			}
		case '+':
			if l.peekNextChar() == '=' {
				tok = l.newComplexToken(token.PLUSEQUAL)
			} else if l.peekNextChar() == '+'{
				tok = l.newComplexToken(token.DOUBLEPLUS)
			} else {
				tok = l.newToken(token.PLUS, l.char)
			}
		case '-':
			if l.peekNextChar() == '=' {
				tok = l.newComplexToken(token.MINUSEQUAL)
			} else if l.peekNextChar() == '-'{
				tok = l.newComplexToken(token.DOUBLEMINUS)
			} else {
				tok = l.newToken(token.MINUS, l.char)
			}
		case '*':
			if l.peekNextChar() == '=' {
				tok = l.newComplexToken(token.STAREQUAL)
			} else if l.peekNextChar() == '*'{
				tok = l.newComplexToken(token.DOUBLESTAR)
			} else {
				tok = l.newToken(token.STAR, l.char)
			}
		case '/':
			if l.peekNextChar() == '=' {
				tok = l.newComplexToken(token.SLASHEQUAL)
			} else if l.peekNextChar() == '/'{
				tok = l.newComplexToken(token.DOUBLESLASH)
			} else {
				tok = l.newToken(token.SLASH, l.char)
			}
		case '%':
			tok = l.newToken(token.PERCENT, l.char)
		case '<':
			if l.peekNextChar() == '=' {
				tok = l.newComplexToken(token.LTEQUAL)
			} else {
				tok = l.newToken(token.LT, l.char)
			}
		case '>':
			if l.peekNextChar() == '=' {
				tok = l.newComplexToken(token.GTEQUAL)
			} else {
				tok = l.newToken(token.GT, l.char)
			}
		case '#':
			tok.Column = l.column
			tok.Line = l.line
			tok.Literal = l.readComment()
			tok.Type = token.COMMENT
		case ',':
			tok = l.newToken(token.COMMA, l.char)
		case ';':
			tok = l.newToken(token.SEMICOLON, l.char)
		case ':':
			tok = l.newToken(token.COLON, l.char)
		case '(':
			tok = l.newToken(token.LPAREN, l.char)
		case ')':
			tok = l.newToken(token.RPAREN, l.char)
		case '{':
			tok = l.newToken(token.LBRACE, l.char)
		case '}':
			tok = l.newToken(token.RBRACE, l.char)
		case '[':
			tok = l.newToken(token.LBRACKET, l.char)
		case ']':
			tok = l.newToken(token.RBRACKET, l.char)
		case '.':
			tok.Line = l.line
			tok.Column = l.column
			tok.Literal, tok.Type = l.readNumber()
			return tok
		case '"':
			tok.Column = l.column  // I did it first of all to store the position of the beginning
			tok.Line = l.line
			tok.Literal = l.readString('"')
			tok.Type = token.STRING
			return tok
		case '\'':
			tok.Column = l.column 
			tok.Line = l.line
			tok.Literal = l.readString('\'')
			tok.Type = token.STRING
			return tok
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			// It's really important we start checking for digits beacuse we've added support to digits
			// on isValidCharForIdent. Then, we don't want to go in that branch with an initial digit.
			if isDigit(l.char) {
				tok.Line = l.line
				tok.Column = l.column
				tok.Literal, tok.Type = l.readNumber()
				return tok
			} else if isValidCharForIdent(l.char) {
				tok.Column = l.column
				tok.Line = l.line
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				return tok  // We can return because readIdentifier() makes what we need from readChar()
			} else {
				tok = l.newToken(token.ILLEGAL, l.char)
			}
	}
	l.readChar()
	return tok
}

func (l * Lexer) readChar() {
	if l.readPosition >= len(l.input){
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]	
	}
	l.position = l.readPosition
	l.readPosition++

	if l.char == '\n' {
		l.line++
		l.column = 0
	} else if l.char == '\t'{
		l.column += 4  // My default tab size is gonna be 4. Maybe I must update smth here.
	} else {
		l.column++
	}
}

func (l *Lexer) newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{
		Type: tokenType, 
		Literal: string(char),
		Line: l.line,
		Column: l.column, 
	}
}

func (l* Lexer) newComplexToken(tokenType token.TokenType) token.Token {
	char := l.char
	startColumn := l.column
	l.readChar()
	literal := string(char) + string(l.char)
	return token.Token{
		Type: tokenType, 
		Literal: literal,
		Line: l.line,
		Column: startColumn, 
	}
}

func (l *Lexer) readString(quoteType byte) string {
	startPos := l.position	
	l.readChar()	
	for l.char != quoteType {
		l.readChar()	
	}
	l.readChar()	
	return l.input[startPos:l.position]
}

func (l *Lexer) readIdentifier() string {
	startPos := l.position
	for isValidCharForIdent(l.char) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func (l *Lexer) readNumber() (string, token.TokenType) {
	/*
	This function iterate through the number updating the lexer via readChar()
	It returns the literal of the number and the type. It can be either an int or a float
	*/
	startPos := l.position
	
	for isDigit(l.char) {
		l.readChar()
	}
	if (l.char != '.') {
		return l.input[startPos:l.position], token.INT
	} 

	l.readChar()
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[startPos:l.position], token.FLOAT
}

func (l* Lexer) readComment() string {
	startPos := l.position 
	for l.char != '\n' {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func isValidCharForIdent(char byte) bool {
	// The '_' is for having support to snake_case. The numbers are just here so 
	// we accept contained digits but not at the start of an identifier
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || 
		char == '_' || '0' <= char && char <= '9'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l* Lexer) peekNextChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

