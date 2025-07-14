package parser

import "github.com/santos-404/myte/ast"

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()
	return leftExp
}
