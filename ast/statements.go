package ast

import (
	"bytes"

	"github.com/santos-404/myte/token"
)

type Statement interface {
	Node
	statementNode()
}


// This is a bit weird; I didn't knoe where to place this.
// It's related w/ expressions but it's its statement 
type ExpressionStatement struct {  
	Token token.Token	
	Expression Expression
}

func (es *ExpressionStatement) statementNode()			{}
func (es *ExpressionStatement) TokenLiteral() string 	{ return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
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

type ConstStatement struct {
	Token token.Token  // This is the token.CONST
	Name *Identifier
	Value Expression 
}

func (cs *ConstStatement) statementNode()			{}
func (cs *ConstStatement) TokenLiteral() string 	{ return cs.Token.Literal }
func (cs *ConstStatement) String() string {
	var out bytes.Buffer

	out.WriteString(cs.TokenLiteral() + " ")
	out.WriteString(cs.Name.String())
	out.WriteString(" = ")

	if cs.Value != nil {
		out.WriteString(cs.Value.String())
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

