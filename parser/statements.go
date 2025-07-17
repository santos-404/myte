package parser

import (
	"github.com/santos-404/myte/ast"
	"github.com/santos-404/myte/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.VAR:
		return p.parseVarStatement()
	case token.CONST:
		return p.parseConstStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()	
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currentToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.currentToken}

	if !p.peekCompareThenAdvance(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	
	// TODO: We don't have to wait for a semicolon if the var is being 
	// initialized. If it's just being declared, then a semicolon is ok
	for p.currentToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return stmt 
}

func (p *Parser) parseConstStatement() *ast.ConstStatement {
	stmt := &ast.ConstStatement{Token: p.currentToken}

	if !p.peekCompareThenAdvance(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	
	// TODO: We don't have to wait for a semicolon if the var is being 
	// initialized. If it's just being declared, then a semicolon is ok
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

