package ast

import (
	"bytes"
	"strings"

	"github.com/santos-404/myte/token"
)

type Expression interface {
	Node
	expressionNode()
}


type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }


type FloatLiteral struct {
	Token token.Token
	Value float64 
}

func (fl *FloatLiteral) expressionNode()      {}
func (fl *FloatLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FloatLiteral) String() string       { return fl.Token.Literal }


type StringLiteral struct {
	Token token.Token
	Value string 
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }


type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BooleanLiteral) String() string       { return bl.Token.Literal }


type NilLiteral struct {
	Token token.Token
}

func (nl *NilLiteral) expressionNode()      {}
func (nl *NilLiteral) TokenLiteral() string { return nl.Token.Literal }
func (nl *NilLiteral) String() string       { return nl.Token.Literal }


type PrefixExpression struct {
	Token token.Token  // Prefix token, e.g.: ! | -
	Operator string 
	Right Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }
func (pe *PrefixExpression) String() string       {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}


type InfixExpression struct {
	Token token.Token  // Infix token, e.g.: + | ==
	Left Expression
	Operator string 
	Right Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string       {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}


type IfExpression struct {
	Token token.Token  
	Condition Expression
	Consequence *BlockStatement
	Alternative *IfExpression // The alternative is else-if/else expression. on else, condition is true
}

func (ie *IfExpression) expressionNode()      {}
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *IfExpression) String() string       {
	var out bytes.Buffer

	out.WriteString(ie.Token.Literal)
	out.WriteString(" ")

	if ie.Token.Type == token.IF {
		out.WriteString(ie.Condition.String())
		out.WriteString(" ")
	}

	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("\n")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}


type FunctionLiteral struct {
	Token token.Token  // The 'fn' token
	Parameters []*Identifier
	Body *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string       {
	var out bytes.Buffer
	var params []string

	for _, param := range fl.Parameters {
		params = append(params, param.String())	
	}

	out.WriteString(fl.Token.Literal)	
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}


type ForExpression struct {
	Token token.Token  // The 'for' token
	Condition Expression
	Body *BlockStatement
}

func (fe *ForExpression) expressionNode()      {}
func (fe *ForExpression) TokenLiteral() string { return fe.Token.Literal }
func (fe *ForExpression) String() string       {
	var out bytes.Buffer

	out.WriteString(fe.Token.Literal)
	out.WriteString(fe.Condition.String())
	out.WriteString(" ")
	out.WriteString(fe.Body.String())

	return out.String()
}

type CallExpression struct {
	Token token.Token  // The '(' token
	Function Expression	
	Arguments []Expression	
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string       {
	var out bytes.Buffer
	var args []string

	out.WriteString(ce.Function.String())
	out.WriteString(ce.Token.Literal)

	for _, arg := range ce.Arguments {
		args = append(args, arg.String())	
	}
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}


type CommentExpression struct {
	Token token.Token  
}

func (ce *CommentExpression) expressionNode()      {}
func (ce *CommentExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CommentExpression) String() string       {
	var out bytes.Buffer
	out.WriteString(ce.Token.Literal)
	return out.String()
}
