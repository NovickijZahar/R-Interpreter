package parser

import (
	"fmt"
	"interpreter/lexer"
)

type Parser struct {
	Tokens []lexer.Token
	Pos int
	
}


func (p* Parser) Match(expected ...lexer.TokenType) (bool, lexer.Token) {
	if (p.Pos < len(p.Tokens)) {
		currentToken := p.Tokens[p.Pos]
		for _, tokenType := range expected {
			if tokenType.Name == currentToken.Kind {
				p.Pos++
				return true, currentToken
			}
		}
	}
	return false, lexer.Token{}
}


func (p* Parser) Require(expected ...lexer.TokenType) (lexer.Token, error) {
	res, token := p.Match(expected...)
	if !res {
		return lexer.Token{}, fmt.Errorf("ошибка на позиции %d, ожидалось %s", p.Pos, expected[0].Name)
	}
	return token, nil
}


// func (p* Parser) ParseCode() ExpressionNode {
// 	root := StatementNode {}
// 	for p.Pos < len(p.Tokens) {
// 		codeStringNode := p.parseExpression()
// 		root.AddNode(codeStringNode)
// 	}
// 	return root
// } 

