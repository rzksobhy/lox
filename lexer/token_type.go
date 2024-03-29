package lexer

const (
	LEFT_PAREN  = iota
	RIGHT_PAREN = iota
	LEFT_BRACE  = iota
	RIGHT_BRACE = iota

	COMMA     = iota
	DOT       = iota
	MINUS     = iota
	PLUS      = iota
	SEMICOLON = iota
	SLASH     = iota
	STAR      = iota

	// One or two character tokens.
	BANG       = iota
	BANG_EQUAL = iota

	EQUAL       = iota
	EQUAL_EQUAL = iota

	GREATER       = iota
	GREATER_EQUAL = iota

	LESS       = iota
	LESS_EQUAL = iota

	// Literals.
	IDENTIFIER = iota
	STRING     = iota
	NUMBER     = iota

	// Keywords.
	AND   = iota
	CLASS = iota
	ELSE  = iota
	FALSE = iota
	FUN   = iota
	FOR   = iota
	IF    = iota
	NIL   = iota
	OR    = iota

	PRINT  = iota
	RETURN = iota
	SUPER  = iota
	THIS   = iota
	TRUE   = iota
	VAR    = iota
	WHILE  = iota

	EOF = iota
)
