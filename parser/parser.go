package parser

import (
	"fmt"

	"github.com/santos-404/myte/ast"
	"github.com/santos-404/myte/lexer"
	"github.com/santos-404/myte/token"
)


type Parser struct {
	l *lexer.Lexer
	
	currentToken token.Token
	peekToken token.Token

	errors []string
}


func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		errors: []string{},
	}

	// This way we set both current and peek tokens
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
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


func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.VAR:
		return p.parseVarStatement()
	default:
		return nil
	}
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.currentToken}

	if !p.peekCompareThenAdvance(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	
	// TODO: We don't have to wait for a semicolon, but for an assignment if the var
	// is being initialized. If it's just being declared, then a semicolon is ok.
	for p.currentToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt 
}


func (p *Parser) Errors () []string {
	return p.errors
}

func (p *Parser) peekError(expectedType token.TokenType) {
	msg := fmt.Sprintf("expected next token to be: %s, got: %s instead",
		expectedType, p.peekToken.Type)
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

