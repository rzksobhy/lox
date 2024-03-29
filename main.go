package main

import (
	"fmt"
	"os"

	"github.com/rzksobhy/lox/lexer"
)

func init() {
	lexer.InitKeywords()
}

func main() {
	data, err := os.ReadFile("./example.txt")
	if err != nil {
		panic(err)
	}

	source := string(data)
	lexer := lexer.NewLexer(source)
	tokens := lexer.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}
