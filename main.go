package main

import (
	"fmt"
	"os"

	"github.com/rzksobhy/lox/scanner"
)

func init() {
	scanner.InitKeywords()
}

func main() {
	data, err := os.ReadFile("./example.txt")
	if err != nil {
		panic(err)
	}

	source := string(data)
	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}
