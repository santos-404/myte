package parser

import (
	"testing"

	"github.com/santos-404/myte/ast"
	"github.com/santos-404/myte/lexer"
)

/*
There are two options here:
	1. We mock the lexer and we just test the parser
	2. We use the lexer so we can have readable test cases

I choose the second options here. I find really useful to
got readable tests, especially if anyone want to contribute
here.
*/

func TestVarStatements(t *testing.T) {
	input := `
var foo = 1;
var bar = 2;

var foobar = 123.45;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"foo"},
		{"bar"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testVarStatement(t, stmt, tt.expectedIdentifier) {
			return 
		}
	}
}


func testVarStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral not 'var'. got=%T", statement)
		return false
	}

	varStmt, ok := statement.(*ast.VarStatement)
	if !ok {
 		t.Errorf("s not *ast.VarStatement. got=%T", statement)
		return false
	}

	// TODO: Dereferencing invalid memory address
	if varStmt.Name.Value != name {
		t.Errorf("varStmt.Name.Value not %s. got=%s", name, varStmt.Name.Value)
		return false
	}

	if varStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, varStmt.Name)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()	
	if len(errors) == 0 {
		return 
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)	
	}
	t.FailNow()
}


func TestReturnStatement(t *testing.T) {
	input := `
return 12;

return foo();
return bar;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
			returnStmt.TokenLiteral())
		}
	}
}


func TestIdentifierExpressions(t *testing.T) {
	input := "foobar";

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not 1 statement. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an ast.ExpressionStatemnt. got=%T", 
			program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier) 
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got=%T", stmt.Expression)	
	}
	if ident.Value != "foobar" {
		t.Fatalf("ident.Value not %s. got=%s", "foobar", ident.Value)	
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not 1 statement. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an ast.ExpressionStatemnt. got=%T", 
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral) 
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got=%T", stmt.Expression)	
	}
	if literal.Value != 5 {
		t.Fatalf("ident.Value not %s. got=%d", "5", literal.Value)	
	}
}
