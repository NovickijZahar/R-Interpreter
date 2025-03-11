package main

import (
	"fmt"
	"interpreter/lexer"
	"interpreter/parser"
	"os"
)


func main() {

	code, _ := os.ReadFile("R/1.R")

	lex := lexer.Lexer {Code: string(code)}

	lex.LexAnalyze()

	tokens := lex.TokenList

	//fmt.Println(tokens)

	p := parser.NewParser(tokens)
	ast := p.Parse()

	for _, a := range ast {
		if a != nil {
			fmt.Println(a.String(""))
			fmt.Println()
		} else {
			fmt.Println("Ошибка разбора")
		}
	}

}

