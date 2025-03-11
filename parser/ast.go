package parser

import (
	"fmt"
)

type ASTNode interface {
	String(indent string) string
}

// AssignNode узел присваивания
type AssignNode struct {
	Left  ASTNode
	Right ASTNode
}

func (a *AssignNode) String(indent string) string {
    leftStr := "nil"
    if a.Left != nil {
        leftStr = a.Left.String(indent + "  ")
    }
    rightStr := "nil"
    if a.Right != nil {
        rightStr = a.Right.String(indent + "  ")
    }
    return fmt.Sprintf("%sAssign:\n%sLeft: %s\n%sToken(<-)\n%sRight: %s",
    indent, indent+"  ", leftStr, indent+"  ", indent+"  ", rightStr)
}

// BinOpNode бинарная операция
type BinOpNode struct {
	Op    string
	Left  ASTNode
	Right ASTNode
}

func (b *BinOpNode) String(indent string) string {
    if b == nil {
        return fmt.Sprintf("%sBinOp(nil)", indent)
    }
    leftStr := "nil"
    if b.Left != nil {
        leftStr = b.Left.String(indent + "  ")
    }
    rightStr := "nil"
    if b.Right != nil {
        rightStr = b.Right.String(indent + "  ")
    }
    return fmt.Sprintf("%sBinOp(%s)\n%s%s\n%sToken(%s)\n%s%s",
        indent, b.Op, indent+"  ", leftStr,
        indent+"  ", b.Op,
        indent+"  ", rightStr)
}

// IdentifierNode идентификатор
type IdentifierNode struct {
	Name string
}

func (i *IdentifierNode) String(indent string) string {
	return fmt.Sprintf("%sIdentifier(%s)", indent, i.Name)
}

// NumberNode числовой литерал
type NumberNode struct {
	Value string
}

func (n *NumberNode) String(indent string) string {
	return fmt.Sprintf("%sNumeric(%s)", indent, n.Value)
}

type IntegerNode struct {
	Value string
}

func (n *IntegerNode) String(indent string) string {
	return fmt.Sprintf("%sInteger(%s)", indent, n.Value)
}

type ComplexNode struct {
	Value string
}

func (n *ComplexNode) String(indent string) string {
	return fmt.Sprintf("%sComplex(%s)", indent, n.Value)
}

type LogicalNode struct {
	Value string
}

func (n *LogicalNode) String(indent string) string {
	return fmt.Sprintf("%sLogical(%s)", indent, n.Value)
}

type CharacterNode struct {
	Value string
}

func (n *CharacterNode) String(indent string) string {
	return fmt.Sprintf("%sCharacterNode(%s)", indent, n.Value)
}

// GroupNode выражение в скобках
type GroupNode struct {
	Expr ASTNode
}

func (g *GroupNode) String(indent string) string {
	return fmt.Sprintf("%sGroup:\n\tToken('(')\n%s\n\tToken(')')", indent, g.Expr.String(indent+"  "))
}

// IfNode узел для if-выражения
type IfNode struct {
    Condition ASTNode  // Условие
    Then      []ASTNode  // Блок Then
    ElseIf    []*IfNode  // Цепочка else if (опционально)
    Else      []ASTNode  // Блок Else (опционально)
}

func (i *IfNode) String(indent string) string {
    s := fmt.Sprintf("%sIf:\n%sCondition: %s\n%sThen:\n", 
        indent, indent+"  ", i.Condition.String(indent+"  "), indent+"  ")
    for _, stmt := range i.Then {
        s += stmt.String(indent + "    ") + "\n"
    }

    // Добавляем цепочку else if
    for _, elseIf := range i.ElseIf {
        s += elseIf.String(indent + "  ") + "\n"
    }

    // Добавляем блок else
    if len(i.Else) > 0 {
        s += fmt.Sprintf("%sElse:\n", indent+"  ")
        for _, stmt := range i.Else {
            s += stmt.String(indent + "    ") + "\n"
        }
    }
    return s
}

// CallNode узел вызова функции
type CallNode struct {
	FuncName ASTNode   // Имя функции
	Args     []ASTNode // Список аргументов
}

func (c *CallNode) String(indent string) string {
	s := fmt.Sprintf("%sCall(%s):\n", indent, c.FuncName.String(indent+"  "))
	for _, arg := range c.Args {
		s += arg.String(indent+"  ") + "\n"
	}
	return s
}

// ForNode узел для цикла for
type ForNode struct {
    Variable  ASTNode   // Переменная цикла
    Range     ASTNode   // Диапазон значений
    Body      []ASTNode // Тело цикла
}

func (f *ForNode) String(indent string) string {
    s := fmt.Sprintf("%sFor:\n%sVariable: %s\n%sRange: %s\n%sBody:\n",
        indent, indent+"  ", f.Variable.String(indent+"  "), indent+"  ", f.Range.String(indent+"  "), indent+"  ")
    for _, stmt := range f.Body {
        s += stmt.String(indent + "    ") + "\n"
    }
    return s
}

// WhileNode узел для цикла while
type WhileNode struct {
    Condition ASTNode   // Условие
    Body      []ASTNode // Тело цикла
}

func (w *WhileNode) String(indent string) string {
    s := fmt.Sprintf("%sWhile:\n%sCondition: %s\n%sBody:\n",
        indent, indent+"  ", w.Condition.String(indent+"  "), indent+"  ")
    for _, stmt := range w.Body {
        s += stmt.String(indent + "    ") + "\n"
    }
    return s
}

// RepeatNode узел для цикла repeat
type RepeatNode struct {
    Body []ASTNode // Тело цикла
}

func (r *RepeatNode) String(indent string) string {
    s := fmt.Sprintf("%sRepeat:\n%sBody:\n", indent, indent+"  ")
    for _, stmt := range r.Body {
        s += stmt.String(indent + "    ") + "\n"
    }
    return s
}

// NextNode узел для ключевого слова next
type NextNode struct{}

func (n *NextNode) String(indent string) string {
    return fmt.Sprintf("%sNext", indent)
}

// BreakNode узел для ключевого слова next
type BreakNode struct{}

func (n *BreakNode) String(indent string) string {
    return fmt.Sprintf("%sBreak", indent)
}

// FunctionNode узел для объявления функции
type FunctionNode struct {
    Name       ASTNode   // Имя функции
    Parameters []ASTNode // Список параметров
    Body       []ASTNode // Тело функции
}

func (f *FunctionNode) String(indent string) string {
    s := fmt.Sprintf("%sFunction:\n", indent)
    if f.Name != nil {
        s += fmt.Sprintf("%sName: %s\n", indent+"  ", f.Name.String(indent+"  "))
    }
    s += fmt.Sprintf("%sParameters:\n", indent+"  ")
    for _, param := range f.Parameters {
        s += param.String(indent + "    ") + "\n"
    }
    s += fmt.Sprintf("%sBody:\n", indent+"  ")
    for _, stmt := range f.Body {
        s += stmt.String(indent + "    ") + "\n"
    }
    return s
}

// ParameterNode узел для параметра функции
type ParameterNode struct {
    Name  ASTNode // Имя параметра
    Value ASTNode // Значение по умолчанию (может быть nil)
}

func (p *ParameterNode) String(indent string) string {
    if p.Value != nil {
        return fmt.Sprintf("%sParameter(%s = %s)", indent, p.Name.String(indent), p.Value.String(indent))
    }
    return fmt.Sprintf("%sParameter(%s)", indent, p.Name.String(indent))
}

// AccessNode узел для доступа через точку
type AccessNode struct {
    Left  ASTNode
    Right ASTNode
}

func (a *AccessNode) String(indent string) string {
    return fmt.Sprintf("%sAccess:\n%sLeft: %s\n%sRight: %s",
        indent, indent+"  ", a.Left.String(indent+"  "), indent+"  ", a.Right.String(indent+"  "))
}