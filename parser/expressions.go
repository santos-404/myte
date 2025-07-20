package parser

import (
	"strconv"

	"github.com/santos-404/myte/ast"
	"github.com/santos-404/myte/token"
)


func (p *Parser) parseExpression(precedence int) ast.Expression {
	// This first prefix can be a number for instance 
	prefixParseFunction := p.prefixParseFns[p.currentToken.Type]
	if prefixParseFunction == nil {
		p.noPrefixParseFunctionError()
		return nil
	}
	leftExp := prefixParseFunction()


	for p.peekToken.Type != token.SEMICOLON && precedence < p.peekPrecedence() {
		infixParseFunction := p.infixParseFns[p.peekToken.Type]
		if infixParseFunction == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infixParseFunction(leftExp)
	}
		
	return leftExp
}


func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.currentToken}

	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		p.parsingLiteralError("integer")
		return nil
	}

	lit.Value = value
	return lit 
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	lit := &ast.FloatLiteral{Token: p.currentToken}

	value, err := strconv.ParseFloat(p.currentToken.Literal, 64)
	if err != nil {
		p.parsingLiteralError("float")
		return nil
	}

	lit.Value = value
	return lit 
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) parseBooleanLiteral() ast.Expression {
	lit := &ast.BooleanLiteral{Token: p.currentToken}

	value, err := strconv.ParseBool(p.currentToken.Literal) 
	if err != nil {
		p.parsingLiteralError("bool")
		return nil
	}

	lit.Value = value
	return lit
}


func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token: p.currentToken,
		Operator: p.currentToken.Literal,
	}

	p.nextToken()
	
	exp.Right = p.parseExpression(PREFIX)

	return exp 
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token: p.currentToken,
		Left: left,
		Operator: p.currentToken.Literal,
	}
	precedence := p.currentPrecedence()
	p.nextToken()
	exp.Right = p.parseExpression(precedence)

	return exp
}

// THIS IS FUCKING MAGICAL
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()
	
	exp := p.parseExpression(LOWEST)

	if !p.peekCompareThenAdvance(token.RPAREN) { 
		return nil
	}

	return exp
}

func (p *Parser) parseIfExpression() ast.Expression {
	exp := &ast.IfExpression{Token: p.currentToken}

	if !p.peekCompareThenAdvance(token.LPAREN) {
		return nil  // TODO: Throw an error here
	}

	p.nextToken()
	exp.Condition = p.parseExpression(LOWEST)

	if !p.peekCompareThenAdvance(token.RPAREN) { 
		return nil  // TODO: errrrrrrrrrrr
	}

	if !p.peekCompareThenAdvance(token.LBRACE) {
		return nil  // TODO: errrrrrrrrrrrr
	}

	exp.Consequence = p.parseBlockStatement()

	if p.peekToken.Type == token.ELSE { 
		p.nextToken()
		if !p.peekCompareThenAdvance(token.LBRACE) {
			return nil
		}

		exp.Alternative = p.parseBlockStatement()
	}

	return exp 
}


func (p *Parser) parseFnLiteral() ast.Expression {
	exp := &ast.FunctionLiteral{Token: p.currentToken}

	if !p.peekCompareThenAdvance(token.LPAREN) {
		return nil	
	}

	exp.Parameters = p.parseParameters()

	exp.Body = p.parseBlockStatement()
	
	return exp
}

func (p *Parser) parseParameters() []*ast.Identifier {
	var params []*ast.Identifier
	p.nextToken()  // We start at '('

	for p.currentToken.Type != token.RPAREN {
		literal := &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
		params = append(params, literal)

		p.nextToken()
		if p.currentToken.Literal == "," {
			p.nextToken()
		}
	}

	p.nextToken()  // We are at ')'
	return params
}
