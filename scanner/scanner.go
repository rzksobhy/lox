package scanner

import (
	"fmt"
	"strconv"

	"github.com/rzksobhy/lox/utils"
)

type Scanner struct {
	start   int
	current int
	line    int
	source  []rune
	tokens  []Token
}

func NewScanner(source string) Scanner {
	return Scanner{start: 0, current: 0, line: 1, source: []rune(source)}
}

func (scanner *Scanner) isAtEnd() bool {
	return scanner.current >= len(scanner.source)
}

func (scanner *Scanner) advance() rune {
	result := scanner.source[scanner.current]
	scanner.current += 1
	return result
}

func (scanner *Scanner) peek() rune {
	if scanner.isAtEnd() {
		return 0
	}

	return scanner.source[scanner.current]
}

func (scanner *Scanner) peekNext() rune {
	idx := scanner.current + 1

	if idx >= len(scanner.source) {
		return 0
	}

	return scanner.source[idx]
}

func (scanner *Scanner) match(expected rune) bool {
	if scanner.isAtEnd() {
		return false
	}

	if scanner.source[scanner.current] != expected {
		return false
	}

	scanner.current += 1
	return true
}

func (scanner *Scanner) addToken(type_ int, literal interface{}) {
	text := string(scanner.source[scanner.start:scanner.current])
	scanner.tokens = append(scanner.tokens, Token{
		Type:    type_,
		Lexeme:  text,
		Literal: literal,
		Line:    scanner.line,
	})
}

func (scanner *Scanner) string() {
	for scanner.peek() != '"' && !scanner.isAtEnd() {
		scanner.advance()
	}

	if scanner.isAtEnd() {
		panic(fmt.Sprintf("[%v] Error: Unterminated string.", scanner.line))
	}

	// consume the closing '"'
	scanner.advance()

	value := string(scanner.source[scanner.start-1 : scanner.current-1])
	scanner.addToken(STRING, value)
}

func (scanner *Scanner) number() {
	for utils.IsDigit(scanner.peek()) {
		scanner.advance()
	}

	if scanner.peek() == '.' && utils.IsDigit(scanner.peekNext()) {
		// consume the "."
		scanner.advance()

		for utils.IsDigit(scanner.peek()) {
			scanner.advance()
		}
	}

	value, err := strconv.ParseFloat(string(scanner.source[scanner.start:scanner.current]), 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse float at %v", scanner.line))
	}

	scanner.addToken(NUMBER, value)
}

func (scanner *Scanner) identifier() {
	for utils.IsAlphaNumeric(scanner.peek()) {
		scanner.advance()
	}

	value := string(scanner.source[scanner.start:scanner.current])
	type_, found := keywords[value]
	if found {
		scanner.addToken(type_, nil)
	} else {
		scanner.addToken(IDENTIFIER, value)
	}
}

func (scanner *Scanner) scanToken() {
	c := scanner.advance()

	switch c {
	case '(':
		scanner.addToken(LEFT_PAREN, nil)
		break
	case ')':
		scanner.addToken(RIGHT_PAREN, nil)
		break
	case '{':
		scanner.addToken(LEFT_BRACE, nil)
		break
	case '}':
		scanner.addToken(RIGHT_BRACE, nil)
		break
	case ',':
		scanner.addToken(COMMA, nil)
		break
	case '.':
		scanner.addToken(DOT, nil)
		break
	case '-':
		scanner.addToken(MINUS, nil)
		break
	case '+':
		scanner.addToken(PLUS, nil)
		break
	case ';':
		scanner.addToken(SEMICOLON, nil)
		break
	case '*':
		scanner.addToken(STAR, nil)
		break

		//> two-char-tokens
	case '!':
		if scanner.match('=') {
			scanner.addToken(BANG_EQUAL, nil)
		} else {
			scanner.addToken(BANG, nil)
		}

		break
	case '=':
		if scanner.match('=') {
			scanner.addToken(EQUAL_EQUAL, nil)
		} else {
			scanner.addToken(EQUAL, nil)
		}

		break
	case '<':
		if scanner.match('=') {
			scanner.addToken(LESS_EQUAL, nil)
		} else {
			scanner.addToken(LESS, nil)
		}

		break
	case '>':
		if scanner.match('=') {
			scanner.addToken(GREATER_EQUAL, nil)
		} else {
			scanner.addToken(GREATER, nil)
		}

		break

		// slash
	case '/':
		if scanner.match('/') {
			for scanner.peek() != '\n' && !scanner.isAtEnd() {
				scanner.advance()
			}
		} else {
			scanner.addToken(SLASH, nil)
		}

		break

		// ignore the whitespaces
	case ' ':
	case '\r':
	case '\t':
		break

	case '\n':
		scanner.line += 1
		break

		// strings
	case '"':
		scanner.string()
		break

	default:
		if utils.IsDigit(c) {
			scanner.number()
		} else if utils.IsAlpha(c) {
			scanner.identifier()
		} else {
			panic(fmt.Sprintf("[%v] Error: Unexpected character '%v'.", scanner.line, string(c)))
		}

		break
	}
}

func (scanner *Scanner) ScanTokens() []Token {
	for !scanner.isAtEnd() {
		scanner.start = scanner.current

		scanner.scanToken()
	}

	scanner.addToken(EOF, nil)
	return scanner.tokens
}
