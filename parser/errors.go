package parser

import (
	"fmt"

	"github.com/santos-404/myte/token"
)


func (p *Parser) peekError(expectedType token.TokenType) {
	msg := fmt.Sprintf("expected next token to be: %s, got: %s instead. Line: %d, column: %d",
		expectedType, p.peekToken.Type, p.peekToken.Line, p.peekToken.Column)
	p.errors = append(p.errors, msg)
}

func (p *Parser) noPrefixParseFunctionError(tt token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", p.currentToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parsingLiteralError(parseTo, literal string) {
	msg := fmt.Sprintf("could not parse %q as %s)", p.currentToken.Literal, parseTo)
	p.errors = append(p.errors, msg)
}
