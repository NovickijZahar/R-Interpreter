package parser

import "interpreter/lexer"

type ExpressionNode interface {
}

type StatementNode struct {
	codeString []ExpressionNode
}

func (s *StatementNode) AddNode(e ExpressionNode) {
	s.codeString = append(s.codeString, e)
}

type VariableNode struct {
	variable lexer.Token
}

type NumberNode struct {
	number lexer.Token
}

type BinOperatorNode struct {
	operator lexer.Token
	leftNode ExpressionNode
	rightNode ExpressionNode
}

type UnarOperatorNode struct {
	operator lexer.Token
	operand ExpressionNode
}