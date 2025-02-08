package lexer

import (
	"fmt"
	"regexp"
	"strings"	
)

var index = 0
var ids = make(map[string]int)


type Lexer struct {
	Code string
	Pos int
	TokenList []Token
	VariablesTokenList []Token
	KeywordsTokenList []Token
	OperatorsTokenList []Token
	ConstantsTokenList []Token
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
			regex_lookahead := regexp.MustCompile("^" + token.Regex + `([^\w]|$)`)
			lookahead := regex_lookahead.FindString(text)
			if (len(lookahead) == 0 && (
				token.Class == "keyword" || token.Class == "constant")) {
				continue
			}
			var s string
			if token.Name == "id" {
				temp_s := strings.ReplaceAll(firstMatch, "`", "")
				if value, exists := ids[temp_s]; exists {
					s = fmt.Sprint(value)
					} else {
						s = fmt.Sprint(index)
						ids[temp_s] = index 
						index++
					}
				}
			line, col := charToLineCol(l.Code, l.Pos)
			newToken := Token {Kind: token.Name + s, Text: firstMatch, Line: line, Column: col}
			l.Pos += len(firstMatch)
			if token.Class != "skip" {
				l.TokenList = append(l.TokenList, newToken)
			}
			switch token.Class {
			case "keyword": 
				l.KeywordsTokenList = append(l.KeywordsTokenList, newToken)
			case "operator":
				l.OperatorsTokenList = append(l.OperatorsTokenList, newToken)
			case "id":
				l.VariablesTokenList = append(l.VariablesTokenList, newToken)
			case "constant":
				l.ConstantsTokenList = append(l.ConstantsTokenList, newToken)
			case "punctuation":
				l.PunctuationsTokenList = append(l.PunctuationsTokenList, newToken)
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