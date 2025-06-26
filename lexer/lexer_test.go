package lexer

import (
	"testing"

	"github.com/santos-404/myte/token"
)

func TestNextToken(t *testing.T) {
	// I cannot be sure now if this will be the final syntax of the lang. 
	input := `const five = 5;
const ten = 10;

const add = fn(x, y) {
	x + y;
}

const result = add(five, ten);
`


	tests := []struct {
		expectedType 	token.TokenType
		expectedLiteral string	
	}{
		{token.CONST, "const"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.CONST, "const"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
	
		{token.CONST, "const"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.CONST, "const"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected %q, got =%q)",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal!= tt.expectedLiteral{
			t.Fatalf("tests[%d] - literal wrong. expected %q, got =%q)",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
