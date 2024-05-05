package lexer

import (
	"github.com/rodrez/rayo/token"
)

type Lexer struct {
	input        string
	position     int  // current position
	readPosition int  // future position
	ch           byte //
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// readPosition is past the length of the input, we got EOF
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		// Shift the character to the next position
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition

	l.readPosition += 1

}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.DIVIDE, l.ch)
	case '*':
		tok = newToken(token.MULTIPLY, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '[':
		tok = newToken(token.LSQUARE, l.ch)
	case ']':
		tok = newToken(token.RSQUARE, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case 0:
		tok.Literal = "0"
		tok.Type = token.EOF

	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {

	return token.Token{Type: tokenType, Literal: string(ch)}
}
