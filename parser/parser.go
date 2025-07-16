package parser

import (
	"fmt"

	"github.com/santos-404/myte/ast"
	"github.com/santos-404/myte/lexer"
	"github.com/santos-404/myte/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
	postfixParseFn func(ast.Expression) ast.Expression
)

type Parser struct {
	l *lexer.Lexer
	errors []string
	
	currentToken token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns map[token.TokenType]infixParseFn
	postfixParseFns map[token.TokenType]postfixParseFn
}


func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		errors: []string{},
	}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)

	// This way we set both current and peek tokens
	p.nextToken()
	p.nextToken()

	return p
}


func (p *Parser) Errors () []string {
	return p.errors
}

func (p *Parser) ParseProgram() *ast.Program { 
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	
	for p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
		
	return program 
}


func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekError(expectedType token.TokenType) {
	msg := fmt.Sprintf("expected next token to be: %s, got: %s instead. Line: %d, column: %d",
		expectedType, p.peekToken.Type, p.peekToken.Line, p.peekToken.Column)
	p.errors = append(p.errors, msg)
}

func (p *Parser) peekCompareThenAdvance(expectedType token.TokenType) bool {
	/* 
	At the beginning I thought was not a good idea to check the following token and 
	also advance to the next token in the same function. As this is a really used
	thing on this parser, I'm gonna implement it is a function. This can be discussed.
	*/ 
	if p.peekToken.Type != expectedType {
		p.peekError(expectedType)
		return false
	}

	p.nextToken()
	return true
}

// These following three are just helper methods to add things to our (pre/in/post)fix maps
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) registerPostfix(tokenType token.TokenType, fn postfixParseFn) {
	p.postfixParseFns[tokenType] = fn
}
