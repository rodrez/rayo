package token

type TokenType string

type Token struct {
  Type TokenType
  Literal string
}


const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers
	IDENT = "IDENT"
	INT = "INT"
	BOOL = "BOOL"

	// Operators
	PLUS = "PLUS"
	MINUS = "MINUS"
	ASSIGN = "ASSIGN"
    MULTIPLY = "MULTIPLY"
    DIVIDE = "DIVIDE"

	// Delimiters
	COMMA = "COMMA"
	SEMICOLON = "SEMICOLON"
	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
    RBRACE = "RBRACE"
    LSQUARE = "LSQUARE"
    RSQUARE = "RSQUARE"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)



