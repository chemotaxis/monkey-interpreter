/*
Package ast implements the abstract syntax tree for the Monkey programming
language
*/
package ast

import (
	"bytes"
	"monkey/token"
	"strings"
)

// Node represents a node in the abstract syntax tree.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents a statement.  Currently, there are only two types of
// statements: let (for assignment) and return (for returning expressions)
// statements.
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression that is to be evaluated to a value.
type Expression interface {
	Node
	expressionNode()
}

// Identifier represents a variable.
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the text character used for this identifier.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String returns the string representation of the identifier.
func (i *Identifier) String() string {
	return i.Value
}

// IntegerLiteral represents a parsed integer.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}

// TokenLiteral returns the text character used for the integer.
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

// String returns the string representation of the integer.
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// PrefixExpression respresents a parsed prefix operator and its operand.  The
// operand is always to the right of the prefix operator (e.g !isFull, -5).
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}

// TokenLiteral returns the literal prefix operator
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

// String returns a string representation of the operator and its operand. For
// example, (-5).
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression represents a parsed infix expression (e.g. 5 + 4).
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}

// TokenLiteral returns the literal for the infix operator.
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

// Program represents the whole syntax tree for a program.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal for the root node in the program.
func (p *Program) TokenLiteral() string {
	var literal string
	if len(p.Statements) > 0 {
		literal = p.Statements[0].TokenLiteral()
	}

	return literal
}

// String returns a string representation of the program.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// LetStatement represents the statement used to assign values to identifiers
// (ie variables).  E.g. let x = 34;
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the text character used to represent the let keyword.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String writes the let statement to a string.
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement represents a return statement (e.g. return <expression>).
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral returns the actual text character used to represent the return
// token.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String returns string representation of the return statement.
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement represents a statement that evaluates to a value.  In
// Monkey, an expression on its own line is perfectly acceptable (e.g. 1 + 3;).
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns the text characters used to represent the expression
// statement.
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String returns a string representation of the expression
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// Boolean is a boolean literal.  It can take the place of any expression
// because it satisfies the Expression interface.
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral returns the actual token used represent the boolean.
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

// String returns the string representation of the boolean.
func (b *Boolean) String() string {
	return b.Token.Literal
}

// IfExpression represents an if expression that has a condition, a consequence,
// and an optional alternative.
type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral returns the string representation of token used for the if
// expression.
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// BlockStatement represents a series of individual statements delimited by
// braces "{}".
type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}

// TokenLiteral returns the string representation of the token used for a block
// statement.
func (bs *BlockStatement) TokenLiteral() string {
	return bs.TokenLiteral()
}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// FunctionLiteral represents a function expression.
type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode() {}

// TokenLiteral returns a string representation of the FunctionLiteral token.
func (fl *FunctionLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())

	return out.String()
}

// CallExpression represents an expression that evaluates to a function and its
// arguments.  The token is actually the left parenthesis "(" and Function can
// be an identifier or a FunctionLiteral.
type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

// TokenLiteral returns a string representation of the token used for
// CallExpression.
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
