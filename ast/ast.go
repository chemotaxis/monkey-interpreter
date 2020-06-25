/*
Package ast implements the abstract syntax tree for the Monkey programming
language
*/
package ast

import (
	"bytes"
	"monkey/token"
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
