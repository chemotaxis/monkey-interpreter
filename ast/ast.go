/*
Package ast implements the abstract syntax tree for the Monkey programming
language
*/
package ast

import "monkey/token"

// Node represents a node in the abstract syntax tree.
type Node interface {
	TokenLiteral() string
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
