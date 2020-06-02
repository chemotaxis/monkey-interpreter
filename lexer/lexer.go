package lexer

import "monkey/token"

// Lexer turns string characters into tokens
type Lexer struct {
	input        string // the source code
	position     int    // the position of the last-read character
	readPosition int    // the position of the next character we will read
	ch           byte   // the current character whose position is position
}

// New returns a reference to a new Lexer
func New(input string) *Lexer {
	lex := &Lexer{input: input}
	lex.readChar()
	return lex
}

// NextToken returns the next token converted from Lexer.input
func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lex.ch {
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '=':
		tok = newToken(token.ASSIGN, lex.ch)
	case '+':
		tok = newToken(token.PLUS, lex.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lex.ch)
	case ',':
		tok = newToken(token.COMMA, lex.ch)
	case '(':
		tok = newToken(token.LPAREN, lex.ch)
	case ')':
		tok = newToken(token.RPAREN, lex.ch)
	case '{':
		tok = newToken(token.LBRACE, lex.ch)
	case '}':
		tok = newToken(token.RBRACE, lex.ch)
	}

	// advance lexer to next character, if available
	lex.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (lex *Lexer) readChar() {
	// 0 is a byte to indicate eof or no character
	lex.ch = 0
	if lex.readPosition < len(lex.input) {
		lex.ch = lex.input[lex.readPosition]
	}

	lex.position = lex.readPosition
	lex.readPosition++
}
