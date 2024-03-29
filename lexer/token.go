package lexer

type Token struct {
	Type    int
	Lexeme  string
	Literal any
	Line    int
}
