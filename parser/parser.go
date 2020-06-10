/*
Package parser implements the parser for the Monkey programming language.  It is
modeled after a Pratt parser
*/
package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
)

// Parser parses tokens
//
// curToken points to the current token being parsed.  peekToken points to the
// next token in order to know what to do with curToken, if needed.
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New creates a new Parser given a lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram creates an ast from a list of tokens
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
