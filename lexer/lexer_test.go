package lexer

import (
	"testing"

	"github.com/santos-404/myte/token"
)

func TestNextTokenLarge(t *testing.T) {
	// I cannot be sure now if this will be the final syntax of the lang. 
	input := `const five = 5;
const ten = 10;

const addIfEqual = fn(x, y) {
	if x == y {
		return x + y;
	}
	return 0;
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
		{token.IDENT, "addIfEqual"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IF, "if"},
		{token.IDENT, "x"},
		{token.EQ, "=="},
		{token.IDENT, "y"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RETURN, "return"},
		{token.INT, "0"},
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

func TestBangOperators(t *testing.T) {
	input := `
	if x != y {
		return !false;
	}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.IDENT, "x"},
		{token.NOT_EQ, "!="},
		{token.IDENT, "y"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.BANG, "!"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("TestBangOperators[%d] - token type wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("TestBangOperators[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestComparisonOperators(t *testing.T) {
	input := `
	if 5 < 10 {
		return 10 > 5;
	}
	if 1 == 1 {
		return 1 != 2;
	}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},

		{token.IF, "if"},
		{token.INT, "1"},
		{token.EQ, "=="},
		{token.INT, "1"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.INT, "1"},
		{token.NOT_EQ, "!="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("TestComparisonOperators[%d] - token type wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("TestComparisonOperators[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
