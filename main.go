package main

import (
	"interpreter/lexer"
	"fmt"
	"os"
)

func main() {
	code, _ := os.ReadFile("R/3.R")

	lex := lexer.Lexer {Code: string(code)}

	lex.LexAnalyze()


	main_file, _ := os.Create("results/result.txt")
	keywords_file, _ := os.Create("results/keywords.txt")
	operators_file, _ := os.Create("results/operators.txt")
	variables_file, _ := os.Create("results/variables.txt")
	constants_file, _ := os.Create("results/constants.txt")
	punctuations_file, _ := os.Create("results/punctuations.txt")

	defer main_file.Close()
	defer keywords_file.Close()
	defer operators_file.Close()
	defer variables_file.Close()
	defer constants_file.Close()
	defer punctuations_file.Close()
	
	main_file.WriteString(fmt.Sprintf("%-25s %-15s %-5s %-5s\n", "Лексема", "Токен", "Строка", "Столбец"))
	main_file.WriteString("=========================================================\n")

	keywords_file.WriteString(fmt.Sprintf("%-25s %-15s %-5s %-5s\n", "Лексема", "Токен", "Строка", "Столбец"))
	keywords_file.WriteString("=========================================================\n")

	operators_file.WriteString(fmt.Sprintf("%-25s %-15s %-5s %-5s\n", "Лексема", "Токен", "Строка", "Столбец"))
	operators_file.WriteString("=========================================================\n")

	variables_file.WriteString(fmt.Sprintf("%-25s %-15s %-5s %-5s\n", "Лексема", "Токен", "Строка", "Столбец"))
	variables_file.WriteString("=========================================================\n")

	constants_file.WriteString(fmt.Sprintf("%-25s %-15s %-5s %-5s\n", "Лексема", "Токен", "Строка", "Столбец"))
	constants_file.WriteString("=========================================================\n")

	punctuations_file.WriteString(fmt.Sprintf("%-25s %-15s %-5s %-5s\n", "Лексема", "Токен", "Строка", "Столбец"))
	punctuations_file.WriteString("=========================================================\n")


	for _, v := range lex.TokenList {
		main_file.WriteString(fmt.Sprintf("%-25s %-15s %-5d %-5d\n", v.Text, v.Kind, v.Line, v.Column))
	}
	for _, v := range lex.KeywordsTokenList {
		keywords_file.WriteString(fmt.Sprintf("%-25s %-15s %-5d %-5d\n", v.Text, v.Kind, v.Line, v.Column))
	}
	for _, v := range lex.OperatorsTokenList {
		operators_file.WriteString(fmt.Sprintf("%-25s %-15s %-5d %-5d\n", v.Text, v.Kind, v.Line, v.Column))
	}
	for _, v := range lex.VariablesTokenList {
		variables_file.WriteString(fmt.Sprintf("%-25s %-15s %-5d %-5d\n", v.Text, v.Kind, v.Line, v.Column))
	}
	for _, v := range lex.ConstantsTokenList {
		constants_file.WriteString(fmt.Sprintf("%-25s %-15s %-5d %-5d\n", v.Text, v.Kind, v.Line, v.Column))
	}
	for _, v := range lex.PunctuationsTokenList {
		punctuations_file.WriteString(fmt.Sprintf("%-25s %-15s %-5d %-5d\n", v.Text, v.Kind, v.Line, v.Column))
	}
}