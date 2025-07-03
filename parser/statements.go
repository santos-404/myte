package parser

import (
	"github.com/santos-404/myte/ast"
	"github.com/santos-404/myte/token"
)


func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.VAR:
		return p.parseVarStatement()
	case token.RETURN:
		return p.parseReturnStatement()
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
	// is being initialized. If it's just being declared, then a semicolon is ok.I 
	for p.currentToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt 
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	// We don't do anything about ensuring the following token is whatever
	// because I don't think it makes any sense. It might be a lot of things
	p.nextToken()
	
	// TODO: Once again, we shouldn't skip everything til a semicolon is found
	for p.currentToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

