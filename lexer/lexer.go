package lexer

import (
	"fmt"
	"regexp"
	"strings"	
)

var names_index = 0
var table_names = make(map[string]int)
var operators_index = 0
var table_operators = make(map[string]int)
var punctuations_index = 0
var table_punctuations = make(map[string]int)
var keywords_index = 0
var table_keywords = make(map[string]int)


type Lexer struct {
	Code string
	Pos int
	TokenList []Token
	NamesTokenList []Token
	KeywordsTokenList []Token
	OperatorsTokenList []Token
	PunctuationsTokenList []Token
}


func (l* Lexer) LexAnalyze() []Token {
	for res, err := l.NextToken(); res; {
		if err != nil {
			fmt.Println(err)
			break
		} 
		res, err = l.NextToken()
	}
	//line, col := charToLineCol(l.Code, l.Pos - 1)
	//l.TokenList = append(l.TokenList, Token {Kind: "EOF", Text: "EOF", Line: line, Column: col, Class: "", Id: "" })
	return l.TokenList
}


func (l* Lexer) NextToken() (bool, error) {
	if l.Pos >= len(l.Code) {
		return false, nil
	}

	var text string

	for _, token := range getTokenTypesList() {
		regex := regexp.MustCompile(`^` + token.Regex)
		text = l.Code[l.Pos:]
		firstMatch := regex.FindString(text)

		if len(firstMatch) != 0 {
			regex_lookahead := regexp.MustCompile("^" + token.Regex + `([^\w\.]|$)`)
			lookahead := regex_lookahead.FindString(text)
			if (len(lookahead) == 0 && (
				token.Class == "keyword" || token.Class == "constant")) {
				continue
			}
			var s string
			var typ string
			var flag_names bool
			var flag_punctuations bool
			var flag_operators bool
			var flag_keywords bool

			if token.Class == "variable" || token.Class == "constant" {
				temp_s := strings.ReplaceAll(firstMatch, "`", "")
				if value, exists := table_names[temp_s]; exists {
					s = fmt.Sprint(value)
					} else {
						flag_names = true
						s = fmt.Sprint(names_index)
						table_names[temp_s] = names_index 
						names_index++
				}
				typ = "N"
			}

			if token.Class == "punctuation" {
				if value, exists := table_punctuations[firstMatch]; exists {
					s = fmt.Sprint(value)
					} else {
						flag_punctuations = true
						s = fmt.Sprint(punctuations_index)
						table_punctuations[firstMatch] = punctuations_index 
						punctuations_index++
				}
				typ = "P"
			}

			if token.Class == "operator" {
				if value, exists := table_operators[firstMatch]; exists {
					s = fmt.Sprint(value)
					} else {
						flag_operators = true
						s = fmt.Sprint(operators_index)
						table_operators[firstMatch] = operators_index 
						operators_index++
				}
				typ = "O"
			}

			if token.Class == "keyword" {
				if value, exists := table_keywords[firstMatch]; exists {
					s = fmt.Sprint(value)
					} else {
						flag_keywords = true
						s = fmt.Sprint(keywords_index)
						table_keywords[firstMatch] = keywords_index 
						keywords_index++
				}
				typ = "K"
			}


			line, col := charToLineCol(l.Code, l.Pos)
			newToken := Token {Kind: token.Name, Text: firstMatch, Line: line, Column: col, Class: token.Class, Id: typ + ":" + s}
			l.Pos += len(firstMatch)
			// if strings.Contains(firstMatch, "\n") && token.Name == "SPACE" {
			// 	newToken.Text = "NEWLINE"
			// 	l.TokenList = append(l.TokenList, newToken)
			// }
			if token.Class != "skip" {
				l.TokenList = append(l.TokenList, newToken)
			}
			switch token.Class {
			case "keyword": 
				if flag_keywords {
					l.KeywordsTokenList = append(l.KeywordsTokenList, newToken)
				}
			case "operator":
				if flag_operators {
					l.OperatorsTokenList = append(l.OperatorsTokenList, newToken)
				}
			case "variable":
				if flag_names{
					l.NamesTokenList = append(l.NamesTokenList, newToken)
				}
			case "constant":
				if flag_names {
					l.NamesTokenList = append(l.NamesTokenList, newToken)
				}
			case "punctuation":
				if flag_punctuations {
					l.PunctuationsTokenList = append(l.PunctuationsTokenList, newToken)
				}
			}
			return true, nil
		}
    }
	err_regex := regexp.MustCompile(`.*[\w|$]`)
	err := err_regex.FindString(text)
	err = strings.Trim(err, " ")
	line, col := charToLineCol(l.Code, l.Pos)
	return true, fmt.Errorf("ошибка на позиции %d %d: %s", line, col, err)
}


func charToLineCol(s string, charIndex int) (line int, col int) {
	if charIndex < 0 || charIndex >= len(s) {
		return -1, -1
	}

	line = 1
	col = 1

	for i := 0; i < charIndex; i++ {
		if s[i] == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}

	return line, col
}