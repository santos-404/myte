package parser

import (
	"fmt"
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
	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("ident.TokenLiteral() not %s. got=%s", 
		"foobar", ident.TokenLiteral())	
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
		t.Fatalf("ident.Value not %d. got=%d", 5, literal.Value)	
	}
	if literal.TokenLiteral() != "5" {
		t.Fatalf("ident.TokenLiteral() not %s. got=%s", 
		"5", literal.TokenLiteral())	
	}
}

func TestFloatLiteralExpression(t *testing.T) {
	input := "5.4;"

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

	literal, ok := stmt.Expression.(*ast.FloatLiteral) 
	if !ok {
		t.Fatalf("exp not *ast.FloatLiteral. got=%T", stmt.Expression)	
	}
	if literal.Value != 5.4 {
		t.Fatalf("ident.Value not %f. got=%f", 5.4, literal.Value)	
	}
	if literal.TokenLiteral() != "5.4" {
		t.Fatalf("ident.TokenLiteral() not %s. got=%s", 
		"5.4", literal.TokenLiteral())	
	}
}

func TestStringLiteralExpression(t *testing.T) {
	input := "'foobar';"

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

	literal, ok := stmt.Expression.(*ast.StringLiteral) 
	if !ok {
		t.Fatalf("exp not *ast.StringLiteral . got=%T", stmt.Expression)	
	}
	if literal.Value != "'foobar'" {
		t.Fatalf("ident.Value not %s. got=%s", "'foobar'", literal.Value)	
	}
	if literal.TokenLiteral() != "'foobar'" {
		t.Fatalf("ident.TokenLiteral() not %s. got=%s", 
		"foobar", literal.TokenLiteral())	
	}
}

func TestBooleanLiteralExpression(t *testing.T) {
	input := "true"

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

	literal, ok := stmt.Expression.(*ast.BooleanLiteral) 
	if !ok {
		t.Fatalf("exp not *ast.BooleanLiteral. got=%T", stmt.Expression)	
	}
	if literal.Value != true {
		t.Fatalf("ident.Value not %t. got=%t", true, literal.Value)	
	}
	if literal.TokenLiteral() != "true" {
		t.Fatalf("ident.TokenLiteral() not %s. got=%s", 
		"true", literal.TokenLiteral())	
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct{
		input 			string
		operator 		string
		integerValue 	interface{}	
	} {
		{"!69;", "!", 69},
		{"-96;", "-", 96},
		{"!true", "!", true},
		{"!false", "!", false},
		{"--5", "--", 5},
		{"++5", "++", 5},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
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
		
		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s. got=%s", tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.integerValue) {
		return 
		}
	}
}

func testIntegerLiteral(t* testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)

	if !ok {
		t.Fatalf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Fatalf("integ.Value not %d. got=%d", value, integ.Value)
		return false }

	return true
}


func TestParsingInfixExpressions(t *testing.T) {
	infixTests:= []struct{
		input 			string
		leftValue 		interface{}
		operator 		string
		rightValue 		interface{}
	} {
		{"5 + 5", 5, "+", 5},
		{"5 - 5", 5, "-", 5},
		{"5 * 5", 5, "*", 5},
		{"5 ** 5", 5, "**", 5},
		{"5 / 5", 5, "/", 5},
		{"5 // 5", 5, "//", 5},
		{"5 % 5", 5, "%", 5},
		{"5 > 5", 5, ">", 5},
		{"5 >= 5", 5, ">=", 5},
		{"5 < 5", 5, "<", 5},
		{"5 <= 5", 5, "<=", 5},
		{"5 == 5", 5, "==", 5},
		{"5 != 5", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for _, tt := range infixTests{
		l := lexer.New(tt.input)
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
		
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.InfixExpression. got=%T", stmt.Expression)
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not %s. got=%s", tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Left, tt.leftValue)  {
			return 
		}
		if !testLiteralExpression(t, exp.Right, tt.rightValue)  {
			return 
		}
		
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input string
		expected string
	} {
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 ** 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 ** 5)))",
		},
		{
			"true",
			"true",
		},
		{
			"false",
			"false",
		},
		{
			"3 > 5 == false",
			"((3 > 5) == false)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"!(true == true)",
			"(!(true == true))",
		},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)

		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)	
		}
	}
}

func testIdentifier (t * testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident.TokenLiteral())
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	bo, ok := exp.(*ast.BooleanLiteral)	
	if !ok {
		t.Errorf("exp not *ast.Boolean. got=%T", exp)
		return false
	}

	if bo.Value != value {
		t.Errorf("bo.Value not %t, got=%t", value, bo.Value)
		return false
	}

	if bo.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("bo.TokenLiteral not %t, got=%s", value, bo.TokenLiteral())
		return false
	}

	return true
}


func testLiteralExpression(
	t* testing.T, 
	exp ast.Expression, 
	expected interface{},
) bool {

	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)	
	case bool:
		return testBooleanLiteral(t, exp, v)
	}

	t.Errorf("type of exp not handled. got=%T", exp)
	return false 
}


func testInfixExpression(
	t *testing.T, 
	exp ast.Expression, 
	left interface{},
	operator string,
	right interface{},
) bool {

	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.OperatorExpressoin. got=%T(%s)", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s', got=%q", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false 
	}

	return true
}


func TestIfExpression(t *testing.T) {
	input := "if (x < y) { x }"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}
	
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an ast.ExpressionStatemnt. got=%T", 
			program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.IfExpression) 
	if !ok {
		t.Fatalf("exp not *ast.IfExpression. got=%T", stmt.Expression)	
	}

	testInfixExpression(t, exp.Condition, "x", "<", "y")

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d", len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statments[0] is not an ast.ExpressionStatement. got=%T",
			exp.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return
	}

	if exp.Alternative != nil {
		t.Errorf("exp.Alternative.Statements was not nil. got=%+v", exp.Alternative)
	}
}

func TestIfElseExpression(t *testing.T) {
	input := "if (x < y) { x } else { y }"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}
	
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an ast.ExpressionStatemnt. got=%T", 
			program.Statements[0])
	}

	exp, ok := stmt.Expression.(*ast.IfExpression) 
	if !ok {
		t.Fatalf("exp not *ast.IfExpression. got=%T", stmt.Expression)	
	}

	testInfixExpression(t, exp.Condition, "x", "<", "y")

	if len(exp.Consequence.Statements) != 1 {
		t.Errorf("consequence is not 1 statements. got=%d", len(exp.Consequence.Statements))
	}

	consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Statments[0] is not an ast.ExpressionStatement. got=%T",
			exp.Consequence.Statements[0])
	}

	if !testIdentifier(t, consequence.Expression, "x") {
		return 
	}

	if len(exp.Alternative.Statements) != 1 {
		t.Errorf("alternateive is not 1 statements. got=%d", len(exp.Consequence.Statements))
	}

	alternative, ok := exp.Alternative.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Alternative.Statments[0] is not an ast.ExpressionStatement. got=%T",
			exp.Alternative.Statements[0])
	}

	if !testIdentifier(t, alternative.Expression, "y") {
		return 
	}
}


func TestFunctionLiteralParsing(t *testing.T) {
	input := "fn(x, y) { x + y; }"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}
	
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an ast.ExpressionStatemnt. got=%T", 
			program.Statements[0])
	}	

	function, ok := stmt.Expression.(*ast.FunctionLiteral) 
	if !ok {
		t.Fatalf("stmt.Expression not *ast.FunctionLiteral. got=%T", stmt.Expression)	
	}

	if len(function.Parameters) != 2 {
		t.Errorf("function literal parameter wrong. want=2. got=%d", len(function.Parameters))
	}

	if !testLiteralExpression(t, function.Parameters[0], "x") {
		return 
	}
	if !testLiteralExpression(t, function.Parameters[1], "y") {
		return 
	}

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statements has not 1 statements, got=%d", 
			len(function.Body.Statements))
	}

	bodyStmt, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function body stmt is not ast.ExpressionStatement. got=%T",
			function.Body.Statements[0])
	}

	if !testInfixExpression(t, bodyStmt.Expression, "x", "+", "y") {
		return 
	}
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct{
		inputFunction string
		expectedParams []string
	} {
		{"fn() {};", []string{}},
		{"fn(x) {};", []string{"x"}},
		{"fn(x, y, z) {};", []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		l := lexer.New(tt.inputFunction)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		function := stmt.Expression.(*ast.FunctionLiteral)

		if len(function.Parameters) != len(tt.expectedParams) {
			t.Errorf("length parameters wrong. want=%d. got=%d",
				len(tt.expectedParams), len(function.Parameters))
		}

		for i, ident := range tt.expectedParams {
			if !testLiteralExpression(t, function.Parameters[i], ident) {
				return 
			}
		}
	}
}
