package lexer

import (
	"fmt"
	"testing"

	"github.com/rodrez/rayo/token"
)


func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {


		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tt.expectedLiteral != tok.Literal {
			// error
			t.Fatalf("test[%d] - Literal wrong. Expected=%q, got %q.",
		i, tt.expectedLiteral, tok.Literal)
		}

		if tt.expectedType != tok.Type {

			t.Fatalf("test[%d] - tokenType wrong. Expected=%q, got %q.",
		i, tt.expectedLiteral, tok.Literal)
		}
	}




}
