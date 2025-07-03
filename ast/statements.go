package ast

import (
	"bytes"

	"github.com/santos-404/myte/token"
)

type Statement interface {
	Node
	statementNode()
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


type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}
func (rs *ReturnStatement) statementNode() 			{}
func (rs *ReturnStatement) TokenLiteral() string	{ return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
