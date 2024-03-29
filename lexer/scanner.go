package lexer

import (
	"fmt"
	"strconv"

	"github.com/rzksobhy/lox/utils"
)

type Lexer struct {
	start   int
	current int
	line    int
	source  []rune
	tokens  []Token
}

func NewLexer(source string) Lexer {
	return Lexer{start: 0, current: 0, line: 1, source: []rune(source)}
}

func (lexer *Lexer) isAtEnd() bool {
	return lexer.current >= len(lexer.source)
}

func (lexer *Lexer) advance() rune {
	result := lexer.source[lexer.current]
	lexer.current += 1
	return result
}

func (lexer *Lexer) peek() rune {
	if lexer.isAtEnd() {
		return 0
	}

	return lexer.source[lexer.current]
}

func (lexer *Lexer) peekNext() rune {
	idx := lexer.current + 1

	if idx >= len(lexer.source) {
		return 0
	}

	return lexer.source[idx]
}

func (lexer *Lexer) match(expected rune) bool {
	if lexer.isAtEnd() {
		return false
	}

	if lexer.source[lexer.current] != expected {
		return false
	}

	lexer.current += 1
	return true
}

func (lexer *Lexer) addToken(type_ int, literal interface{}) {
	text := string(lexer.source[lexer.start:lexer.current])
	lexer.tokens = append(lexer.tokens, Token{
		Type:    type_,
		Lexeme:  text,
		Literal: literal,
		Line:    lexer.line,
	})
}

func (lexer *Lexer) string() {
	for lexer.peek() != '"' && !lexer.isAtEnd() {
		lexer.advance()
	}

	if lexer.isAtEnd() {
		panic(fmt.Sprintf("[%v] Error: Unterminated string.", lexer.line))
	}

	// consume the closing '"'
	lexer.advance()

	value := string(lexer.source[lexer.start-1 : lexer.current-1])
	lexer.addToken(STRING, value)
}

func (lexer *Lexer) number() {
	for utils.IsDigit(lexer.peek()) {
		lexer.advance()
	}

	if lexer.peek() == '.' && utils.IsDigit(lexer.peekNext()) {
		// consume the "."
		lexer.advance()

		for utils.IsDigit(lexer.peek()) {
			lexer.advance()
		}
	}

	value, err := strconv.ParseFloat(string(lexer.source[lexer.start:lexer.current]), 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse float at %v", lexer.line))
	}

	lexer.addToken(NUMBER, value)
}

func (lexer *Lexer) identifier() {
	for utils.IsAlphaNumeric(lexer.peek()) {
		lexer.advance()
	}

	value := string(lexer.source[lexer.start:lexer.current])
	type_, found := keywords[value]
	if found {
		lexer.addToken(type_, nil)
	} else {
		lexer.addToken(IDENTIFIER, value)
	}
}

func (lexer *Lexer) scanToken() {
	c := lexer.advance()

	switch c {
	case '(':
		lexer.addToken(LEFT_PAREN, nil)
		break
	case ')':
		lexer.addToken(RIGHT_PAREN, nil)
		break
	case '{':
		lexer.addToken(LEFT_BRACE, nil)
		break
	case '}':
		lexer.addToken(RIGHT_BRACE, nil)
		break
	case ',':
		lexer.addToken(COMMA, nil)
		break
	case '.':
		lexer.addToken(DOT, nil)
		break
	case '-':
		lexer.addToken(MINUS, nil)
		break
	case '+':
		lexer.addToken(PLUS, nil)
		break
	case ';':
		lexer.addToken(SEMICOLON, nil)
		break
	case '*':
		lexer.addToken(STAR, nil)
		break

		//> two-char-tokens
	case '!':
		if lexer.match('=') {
			lexer.addToken(BANG_EQUAL, nil)
		} else {
			lexer.addToken(BANG, nil)
		}

		break
	case '=':
		if lexer.match('=') {
			lexer.addToken(EQUAL_EQUAL, nil)
		} else {
			lexer.addToken(EQUAL, nil)
		}

		break
	case '<':
		if lexer.match('=') {
			lexer.addToken(LESS_EQUAL, nil)
		} else {
			lexer.addToken(LESS, nil)
		}

		break
	case '>':
		if lexer.match('=') {
			lexer.addToken(GREATER_EQUAL, nil)
		} else {
			lexer.addToken(GREATER, nil)
		}

		break

		// slash
	case '/':
		if lexer.match('/') {
			for lexer.peek() != '\n' && !lexer.isAtEnd() {
				lexer.advance()
			}
		} else {
			lexer.addToken(SLASH, nil)
		}

		break

		// ignore the whitespaces
	case ' ':
	case '\r':
	case '\t':
		break

	case '\n':
		lexer.line += 1
		break

		// strings
	case '"':
		lexer.string()
		break

	default:
		if utils.IsDigit(c) {
			lexer.number()
		} else if utils.IsAlpha(c) {
			lexer.identifier()
		} else {
			panic(fmt.Sprintf("[%v] Error: Unexpected character '%v'.", lexer.line, string(c)))
		}

		break
	}
}

func (lexer *Lexer) ScanTokens() []Token {
	for !lexer.isAtEnd() {
		lexer.start = lexer.current

		lexer.scanToken()
	}

	lexer.addToken(EOF, nil)
	return lexer.tokens
}
