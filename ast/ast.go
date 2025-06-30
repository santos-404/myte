package ast

import "github.com/santos-404/myte/token"

type Node interface {
	TokenLiteral() string  // This is here only for debugging reasons
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}


type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}


type VarStatement struct {
	Token token.Token  // This is the token.VAR
	Name *Identifier
	Value Expression 
}

func (vs *VarStatement) statementNode()			{}
func (vs *VarStatement) TokenLiteral() string 	{ return vs.Token.Literal }


type Identifier struct {
	Token token.Token  // This one; the token.IDENT
	Value string 
}

func (i *Identifier) expressionNode()			{}
func (i *Identifier) TokenLiteral() string 		{ return i.Token.Literal }

