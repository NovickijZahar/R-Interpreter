package main

import (
	"interpreter/lexer"
	"fmt"
	"os"
)

func main() {
	code, _ := os.ReadFile("R/1.R")

	lex := lexer.Lexer {Code: string(code)}

	lex.LexAnalyze()


	main_file, _ := os.Create("results/result.txt")
	keywords_file, _ := os.Create("results/keywords.txt")
	operators_file, _ := os.Create("results/operators.txt")
	names_file, _ := os.Create("results/names.txt")
	punctuations_file, _ := os.Create("results/punctuations.txt")

	defer main_file.Close()
	defer keywords_file.Close()
	defer operators_file.Close()
	defer names_file.Close()
	defer punctuations_file.Close()
	
	main_file.WriteString(fmt.Sprintf("%-25s %-15s %-5s %-5s %-5s\n", "Лексема", "Тип токена", "Строка", "Столбец", "Id"))
	main_file.WriteString("=========================================================\n")

	keywords_file.WriteString(fmt.Sprintf("%-25s %-15s %-5s %-5s %-5s\n", "Лексема", "Тип токена", "Строка", "Столбец", "Id"))
	keywords_file.WriteString("=========================================================\n")

	operators_file.WriteString(fmt.Sprintf("%-25s %-15s %-5s %-5s %-5s\n", "Лексема", "Тип токена", "Строка", "Столбец", "Id"))
	operators_file.WriteString("=========================================================\n")

	names_file.WriteString(fmt.Sprintf("%-25s %-15s %-5s %-5s %-5s\n", "Лексема", "Тип токена", "Строка", "Столбец", "Id"))
	names_file.WriteString("=========================================================\n")

	punctuations_file.WriteString(fmt.Sprintf("%-25s %-15s %-5s %-5s %-5s\n", "Лексема", "Тип токена", "Строка", "Столбец", "Id"))
	punctuations_file.WriteString("=========================================================\n")


	for _, v := range lex.TokenList {
		main_file.WriteString(fmt.Sprintf("%-25s %-15s %-5d %-5d %-5s\n", v.Text, v.Kind, v.Line, v.Column, v.Id))
	}
	for _, v := range lex.KeywordsTokenList {
		keywords_file.WriteString(fmt.Sprintf("%-25s %-15s %-5d %-5d %-5s\n", v.Text, v.Kind, v.Line, v.Column, v.Id))
	}
	for _, v := range lex.OperatorsTokenList {
		operators_file.WriteString(fmt.Sprintf("%-25s %-15s %-5d %-5d %-5s\n", v.Text, v.Kind, v.Line, v.Column, v.Id))
	}
	for _, v := range lex.NamesTokenList {
		names_file.WriteString(fmt.Sprintf("%-25s %-15s %-5d %-5d %-5s\n", v.Text, v.Kind, v.Line, v.Column, v.Id))
	}
	for _, v := range lex.PunctuationsTokenList {
		punctuations_file.WriteString(fmt.Sprintf("%-25s %-15s %-5d %-5d %-5s\n", v.Text, v.Kind, v.Line, v.Column, v.Id))
	}
}