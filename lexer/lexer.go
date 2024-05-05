package lexer

type Lexer struct {
	input string
	position int // current position
	readPosition int // future position
	ch byte //
}

func New(input string)  *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar
