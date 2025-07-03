package ast

import (
	"bytes"

	"github.com/santos-404/myte/token"
)

type Node interface {
	TokenLiteral() string  // This is here only for debugging reasons
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}


type VarStatement struct {
	Token token.Token  // This is the token.VAR
	Name *Identifier
	Value Expression 
}

func (vs *VarStatement) statementNode()			{}
func (vs *VarStatement) TokenLiteral() string 	{ return vs.Token.Literal }
func (vs *VarStatement) String() string {
	var out bytes.Buffer

	out.WriteString(vs.TokenLiteral() + " ")
	out.WriteString(vs.Name.String())
	out.WriteString(" = ")

	if vs.Value != nil {
		out.WriteString(vs.Value.String())
	}

	out.WriteString(";")

	return out.String()
}


type Identifier struct {
	Token token.Token  // This one; the token.IDENT
	Value string 
}

func (i *Identifier) expressionNode()			{}
func (i *Identifier) TokenLiteral() string 		{ return i.Token.Literal }
func (i *Identifier) String() string 			{ return i.Value }

