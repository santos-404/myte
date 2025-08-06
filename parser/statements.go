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
	p.nextToken()	

	if p.currentToken.Type == token.ASSIGN {
		p.nextToken()	
		stmt.Value = p.parseExpression(LOWEST)
	} else {
		stmt.Value = &ast.NilLiteral{Token: token.Token{Type: token.NIL}}
	}

	p.nextToken()  // We go to the ";" in both cases
	
	return stmt 
}

func (p *Parser) parseConstStatement() *ast.ConstStatement {
	stmt := &ast.ConstStatement{Token: p.currentToken}

	if !p.peekCompareThenAdvance(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	p.nextToken()
	
	if p.currentToken.Type == token.ASSIGN {
		p.nextToken()	
		stmt.Value = p.parseExpression(LOWEST)
	} else {
		stmt.Value = &ast.NilLiteral{Token: token.Token{Type: token.NIL}}
	}

	p.nextToken()  // We go to the ";" in both cases
	return stmt 
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()
	
	stmt.ReturnValue = p.parseExpression(LOWEST)

	p.nextToken()
	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	/*
	The block statement starts with a '{' Token.
	*/
	block := &ast.BlockStatement{Token: p.currentToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for p.currentToken.Type != token.RBRACE && p.currentToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

