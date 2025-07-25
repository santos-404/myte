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
	if x == y or y == x {
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
		{token.OR, "or"},
		{token.IDENT, "y"},
		{token.EQ, "=="},
		{token.IDENT, "x"},
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
		{token.NOTEQ, "!="},
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
		{token.NOTEQ, "!="},
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



func TestNextTokenWithQuotes(t *testing.T) {
	input := `
const singleQuote = 'hello world';
const doubleQuote = "hello world";
const mixedQuotes = "'Hello', she said.";
`

	tests := []struct {
		expectedType   token.TokenType
		expectedLiteral string
	}{
		{token.CONST, "const"},
		{token.IDENT, "singleQuote"},
		{token.ASSIGN, "="},
		{token.STRING, "'hello world'"},
		{token.SEMICOLON, ";"},

		{token.CONST, "const"},
		{token.IDENT, "doubleQuote"},
		{token.ASSIGN, "="},
		{token.STRING, `"hello world"`},
		{token.SEMICOLON, ";"},

		{token.CONST, "const"},
		{token.IDENT, "mixedQuotes"},
		{token.ASSIGN, "="},
		{token.STRING, `"'Hello', she said."`},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextTokenWithFloat(t *testing.T) {
	input := `
const startingWithNumber = 123.456;
const startingWithPoint= .987;

const operationWithFloats = fn (x, y) {
	return 12.03 - .98;
}
`

	tests := []struct {
		expectedType   token.TokenType
		expectedLiteral string
	}{
		{token.CONST, "const"},
		{token.IDENT, "startingWithNumber"},
		{token.ASSIGN, "="},
		{token.FLOAT, "123.456"},
		{token.SEMICOLON, ";"},

		{token.CONST, "const"},
		{token.IDENT, "startingWithPoint"},
		{token.ASSIGN, "="},
		{token.FLOAT, ".987"},
		{token.SEMICOLON, ";"},

		{token.CONST, "const"},
		{token.IDENT, "operationWithFloats"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FLOAT, "12.03"},
		{token.MINUS, "-"},
		{token.FLOAT, ".98"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestLineAndColumn(t *testing.T) {
	input := `
if x != y {
	return !false;
}
	`

	tests := []struct {
		expectedLine int
		expectedColumn int
	}{
		{1, 1},
		{1, 4},
		{1, 6},
		{1, 9},
		{1, 11},
		{2, 5},
		{2, 12},
		{2, 13},
		{2, 18},
		{3, 1},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Line != tt.expectedLine{
			t.Fatalf("TestLineAndColumn[%d] - token line wrong. expected=%d, got=%d",
				i, tt.expectedLine, tok.Line)
		}
		if tok.Column!= tt.expectedColumn{
			t.Fatalf("TestLineAndColumn[%d] - token column wrong. expected=%d, got=%d",
				i, tt.expectedColumn, tok.Column)
		}
	}
}

func TestComments(t *testing.T) {
	input := `
# This is a full line comment
const varWithComment = 0;  # A comment at the end of the line

const varToTestIfTheEndOfTheCommentWorks = 1;

# Comment with some weird values like 'example', / or ! .
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.COMMENT, "# This is a full line comment"},
		{token.CONST, "const"},
		{token.IDENT, "varWithComment"},
		{token.ASSIGN, "="},
		{token.INT, "0"},
		{token.SEMICOLON, ";"},
		{token.COMMENT, "# A comment at the end of the line"},
		{token.CONST, "const"},
		{token.IDENT, "varToTestIfTheEndOfTheCommentWorks"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.COMMENT, "# Comment with some weird values like 'example', / or ! ."},
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

