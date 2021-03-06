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

var whitespace = map[byte]struct{}{
	' ':  {},
	'\t': {},
	'\n': {},
	'\r': {},
}

func (lex *Lexer) peekChar() byte {
	var chIndex byte // zero value is 0
	if lex.readPosition < len(lex.input) {
		chIndex = lex.input[lex.readPosition]
	}
	return chIndex
}

func (lex *Lexer) skipWhitespace() {
	for _, inSet := whitespace[lex.ch]; inSet; {
		lex.readChar()
		_, inSet = whitespace[lex.ch]
	}
}

// NextToken returns the next token converted from Lexer.input
func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	lex.skipWhitespace()

	switch lex.ch {
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '=':
		tok = newToken(token.ASSIGN, lex.ch)
		if isEqualSign(lex.peekChar()) {
			tok.Type = token.EQ
			tok.Literal = lex.read(isEqualSign)
		}
	case '+':
		tok = newToken(token.PLUS, lex.ch)
	case '-':
		tok = newToken(token.MINUS, lex.ch)
	case '*':
		tok = newToken(token.ASTERISK, lex.ch)
	case '/':
		tok = newToken(token.SLASH, lex.ch)
	case '<':
		tok = newToken(token.LT, lex.ch)
	case '>':
		tok = newToken(token.GT, lex.ch)
	case '!':
		tok = newToken(token.BANG, lex.ch)
		if isEqualSign(lex.peekChar()) {
			tok.Type = token.NOT_EQ
			tok.Literal = lex.readWithOffset(isEqualSign, 1)
		}
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
	default:
		if isLetter(lex.ch) {
			tok.Literal = lex.read(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lex.ch) {
			tok.Literal = lex.read(isDigit)
			tok.Type = token.INT
			return tok
		}
		tok = newToken(token.ILLEGAL, lex.ch)
	}

	// advance lexer to next character, if available
	lex.readChar()
	return tok
}

// read returns a substring containing characters that satisfy boolFunc.
func (lex *Lexer) read(boolFunc func(byte) bool) string {
	return lex.readWithOffset(boolFunc, 0)
}

// readWithOffset returns a substring containing characters that satisfy
// boolFunc, except for possibly the first n characters specified by offset.
func (lex *Lexer) readWithOffset(boolFunc func(byte) bool, offset int) string {
	position := lex.position
	for i := 0; i < offset; i++ {
		lex.readChar()
	}

	for boolFunc(lex.ch) {
		lex.readChar()
	}
	return lex.input[position:lex.position]
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

func isEqualSign(ch byte) bool {
	return ch == '='
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
