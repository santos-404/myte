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
		// Think about the first two and how to join em
		case '=':
			tok = newToken(token.ASSIGN, l.char)
		case '!':
			tok = newToken(token.EXCLAMATION, l.char)
		case '+':
			tok = newToken(token.PLUS, l.char)
		case '-':
			tok = newToken(token.MINUS, l.char)
		case '*':
			tok = newToken(token.ASTERISK, l.char)
		case '/':
			tok = newToken(token.SLASH, l.char)
		case '<':
			tok = newToken(token.LT, l.char)
		case '>':
			tok = newToken(token.GT, l.char)
		case ',':
			tok = newToken(token.COMMA, l.char)
		case ';':
			tok = newToken(token.SEMICOLON, l.char)
		case '(':
			tok = newToken(token.LPAREN, l.char)
		case ')':
			tok = newToken(token.RPAREN, l.char)
		case '{':
			tok = newToken(token.LBRACE, l.char)
		case '}':
			tok = newToken(token.RBRACE, l.char)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
		default:
			if isLetter(l.char) {
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				return tok  // We can return because readIdentifier() makes what we need from readChar()
			} else if isDigit(l.char) {
				tok.Literal = l.readNumber()
				tok.Type = token.INT
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.char)
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
}

func newToken (tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func (l *Lexer) readIdentifier() string {
	startPos := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

func (l *Lexer) readNumber() string {
	startPos := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[startPos:l.position]
}

// This is a impactful point for the performance of the lang. Might be improved
func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'  // This last is because to support snake_case
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}
