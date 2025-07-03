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
