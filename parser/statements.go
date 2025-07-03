package parser

import (
	"github.com/santos-404/myte/ast"
	"github.com/santos-404/myte/token"
)


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

