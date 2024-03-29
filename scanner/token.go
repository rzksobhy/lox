package scanner

type Token struct {
	Type   int
	Lexeme  string
	Literal interface{}
	Line    int
}
