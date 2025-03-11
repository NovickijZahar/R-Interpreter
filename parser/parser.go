package parser

import (
	"fmt"
	"interpreter/lexer"
)

// Parser структура парсера
type Parser struct {
	tokens []lexer.Token
	pos    int
}

// NewParser создает новый парсер с фильтрацией skip-токенов
func NewParser(tokens []lexer.Token) *Parser {
	filtered := make([]lexer.Token, 0)
	filtered = append(filtered, tokens...)
	return &Parser{tokens: filtered, pos: 0}
}

// currentToken возвращает текущий токен
func (p *Parser) currentToken() *lexer.Token {
	if p.pos >= len(p.tokens) {
		return nil
	}
	return &p.tokens[p.pos]
}

// consumeToken переходит к следующему токену
func (p *Parser) consumeToken() {
	p.pos++
}

// Parse начинает разбор
func (p *Parser) Parse() []ASTNode {
	return p.parseStatements()
}

// parseAssignment разбирает присваивание или выражение
func (p *Parser) parseAssignment() ASTNode {
	if p.currentToken() == nil {
		return nil
	}

	// Пробуем прочитать идентификатор и оператор присваивания
	if p.currentToken().Kind == "ident" {
		ident := &IdentifierNode{Name: p.currentToken().Text}
		savePos := p.pos
		p.consumeToken()

		if p.currentToken() != nil && p.currentToken().Kind == "assignment" && p.currentToken().Text == "<-" {
            p.consumeToken()

            // Если следующий токен - ключевое слово "function", это объявление функции
            if p.currentToken() != nil && p.currentToken().Kind == "function" && p.currentToken().Text == "function" {
                functionNode := p.parseFunctionDeclaration()
				if functionNode == nil {
					p.errorf("Ожидается выражение после '<-'")
				}
                return &AssignNode{Left: ident, Right: functionNode}
            }

            // Иначе это обычное присваивание
            expr := p.parseExpression()
			if expr == nil {
				p.errorf("Ожидается выражение после '<-'")
			}
            return &AssignNode{Left: ident, Right: expr}
        }
		// Откат, если это не присваивание
		p.pos = savePos
	}

	return p.parseExpression()
}

// parseExpression разбирает выражения уровня сложения/вычитания
func (p *Parser) parseExpression() ASTNode {
    node := p.parseTerm()

    for {
        current := p.currentToken()
        if current == nil {
            break
        }

        // Поддержка арифметических операций и операторов сравнения
        if current.Kind == "arithmetic" || current.Kind == "comparison" || 
			(current.Kind == "miscellaneous" && current.Text == ":") {
            p.consumeToken()
            right := p.parseTerm()
			if right == nil {
                p.errorf("Ожидается выражение после оператора '%s'", current.Text)
            }
            node = &BinOpNode{Op: current.Text, Left: node, Right: right}
        } else {
            break
        }
    }

    return node
}

// parseTerm разбирает выражения уровня умножения/деления
func (p *Parser) parseTerm() ASTNode {
	node := p.parseFactor()

	for {
		current := p.currentToken()
		if current == nil || current.Kind != "arithmetic" || (current.Text != "*" && current.Text != "/") {
			break
		}

		p.consumeToken()
		right := p.parseFactor()
		if right == nil {
            p.errorf("Ожидается выражение после оператора '%s'", current.Text)
        }
		node = &BinOpNode{Op: current.Text, Left: node, Right: right}
	}

	return node
}

// parseFactor разбирает возведение в степень и унарные операторы
func (p *Parser) parseFactor() ASTNode {
	node := p.parsePrimary()

	if p.currentToken() != nil && p.currentToken().Kind == "arithmetic" && p.currentToken().Text == "^" {
		p.consumeToken()
		right := p.parseFactor() // Правоассоциативность
		node = &BinOpNode{Op: "^", Left: node, Right: right}
	}

	return node
}

// parsePrimary разбирает базовые элементы
func (p *Parser) parsePrimary() ASTNode {
	current := p.currentToken()
	if current == nil {
		return nil
	}

	switch current.Kind {
	case "ident":
		// Сначала считаем идентификатор
		var left ASTNode = &IdentifierNode{Name: current.Text} // Используем тип ASTNode
        p.consumeToken()

        // Проверяем, является ли следующий токен точкой
        for p.currentToken() != nil && p.currentToken().Kind == "access" {
            p.consumeToken() // Пропускаем '.'

            // Следующий токен должен быть идентификатором
            if p.currentToken() == nil || p.currentToken().Kind != "ident" {
                p.errorf("Ожидается идентификатор после '.'")
            }

            // Создаем узел доступа
            right := &IdentifierNode{Name: p.currentToken().Text}
            p.consumeToken()

            // Обновляем left для поддержки цепочки доступа (например, a.b.c)
            left = &AccessNode{Left: left, Right: right} // Теперь left имеет тип ASTNode
        }

        // Проверяем, является ли следующий токен открывающей скобкой '('
        if p.currentToken() != nil && p.currentToken().Kind == "lpar" {
            p.consumeToken() // Пропускаем '('

            // Разбираем список аргументов
            args := p.parseArguments()

            // Проверяем наличие закрывающей скобки ')'
            if p.currentToken() == nil || p.currentToken().Kind != "rpar" {
                p.errorf("Ожидается ')' после списка аргументов")
            }
            p.consumeToken() // Пропускаем ')'

            // Возвращаем узел вызова функции
            return &CallNode{FuncName: left, Args: args}
        }

        // Если это не вызов функции, возвращаем узел доступа или идентификатор
        return left

	case "numeric":
		node := &NumberNode{Value: current.Text}
		p.consumeToken()
		return node

	case "integer":
		node := &IntegerNode{Value: current.Text}
		p.consumeToken()
		return node

	case "complex":
		node := &ComplexNode{Value: current.Text}
		p.consumeToken()
		return node

	case "logical":
		node := &LogicalNode{Value: current.Text}
		p.consumeToken()
		return node

	case "character":
		node := &CharacterNode{Value: current.Text}
		p.consumeToken()
		return node

	case "lpar":
		p.consumeToken() // Пропускаем '('
		expr := p.parseExpression()
		if p.currentToken() == nil || p.currentToken().Kind != "rpar" {
			p.errorf("Ожидается ')'")
		}
		p.consumeToken() // Пропускаем ')'
		return &GroupNode{Expr: expr}

	case "arithmetic":
		// Унарные + и -
		if current.Text == "+" || current.Text == "-" {
			op := current.Text
			p.consumeToken()
			expr := p.parsePrimary()
			return &BinOpNode{Op: op, Left: &NumberNode{Value: "0"}, Right: expr}
		}
		p.errorf("Неожиданный оператор: %s", current.Text)

	default:
		p.errorf("Неожиданный токен: %s", current.Kind)
	}

	panic("")
}

// parseStatements разбирает все выражения до конца списка токенов
func (p *Parser) parseStatements() []ASTNode {
    statements := make([]ASTNode, 0)

    for p.currentToken() != nil {

        if p.currentToken() == nil {
            break
        }

        var stmt ASTNode

        // Проверяем, является ли текущий токен ключевым словом "if"
        if p.currentToken().Kind == "if" && p.currentToken().Text == "if" {
            stmt = p.parseIfStatement()
        } else if p.currentToken().Kind == "for" && p.currentToken().Text == "for" {
            stmt = p.parseFor()
        } else if p.currentToken().Kind == "while" && p.currentToken().Text == "while" {
            stmt = p.parseWhile()
        } else if p.currentToken().Kind == "repeat" && p.currentToken().Text == "repeat" {
            stmt = p.parseRepeat()
        } else if p.currentToken().Kind == "break" && p.currentToken().Text == "break" {
            stmt = p.parseBreak()
        } else if p.currentToken().Kind == "next" && p.currentToken().Text == "next" {
            stmt = p.parseNext()
        } else {
            stmt = p.parseAssignment()
        }

        if stmt != nil {
            statements = append(statements, stmt)
        }

    }

    return statements
}

// parseIfStatement разбирает if-выражение
func (p *Parser) parseIfStatement() ASTNode {
    // Пропускаем ключевое слово "if"
    if p.currentToken().Kind != "if" || p.currentToken().Text != "if" {
        p.errorf("Ожидается ключевое слово 'if'")
    }
    p.consumeToken()

    // Проверяем наличие открывающей скобки '('
    if p.currentToken().Kind != "lpar" {
        p.errorf(
            "Ожидается '(' после 'if', но получен токен %s (текст: '%s')",
            p.currentToken().Kind,
            p.currentToken().Text,
        )
    }
    p.consumeToken() // Пропускаем '('

    // Разбираем условие
    condition := p.parseExpression()

    // Проверяем наличие закрывающей скобки ')'
    if p.currentToken().Kind != "rpar" {
        p.errorf(
            "Ожидается ')' после условия, но получен токен %s (текст: '%s')",
            p.currentToken().Kind,
            p.currentToken().Text,
        )
    }
    p.consumeToken() // Пропускаем ')'

    // Разбираем блок Then
    thenBlock := p.parseBlock()

    // Разбираем цепочку else if и блок else
    var elseIfBlocks []*IfNode
    var elseBlock []ASTNode

    for p.currentToken() != nil && p.currentToken().Kind == "else" && p.currentToken().Text == "else" {
        p.consumeToken() // Пропускаем 'else'

        // Если после 'else' идет 'if', разбираем как else if
        if p.currentToken() != nil && p.currentToken().Kind == "if" && p.currentToken().Text == "if" {
            elseIfNode := p.parseIfStatement() // Рекурсивно разбираем else if
            elseIfBlocks = append(elseIfBlocks, elseIfNode.(*IfNode))
        } else {
            // Иначе разбираем как блок else
            elseBlock = p.parseBlock()
            break
        }
    }

    return &IfNode{
        Condition: condition,
        Then:      thenBlock,
        ElseIf:    elseIfBlocks,
        Else:      elseBlock,
    }
}

// parseBlock разбирает блок кода в фигурных скобках
func (p *Parser) parseBlock() []ASTNode {
    if p.currentToken().Kind != "begin" {
        p.errorf("Ожидается '{' для начала блока")
    }
    p.consumeToken() // Пропускаем '{'

    var statements []ASTNode
    for p.currentToken() != nil && p.currentToken().Kind != "end" {

        if p.currentToken() == nil || p.currentToken().Kind == "end" {
            break
        }

        stmt := p.parseStatement()
        if stmt != nil {
            statements = append(statements, stmt)
        }
    }

    if p.currentToken().Kind != "end" {
        p.errorf("Ожидается '}' для завершения блока")
    }
    p.consumeToken() // Пропускаем '}'

    return statements
}

// parseStatement разбирает одно выражение
func (p *Parser) parseStatement() ASTNode {
    if p.currentToken() == nil {
        return nil
    }

    // Проверяем, является ли текущий токен ключевым словом "if"
    if p.currentToken().Kind == "if" && p.currentToken().Text == "if" {
        return p.parseIfStatement()
    } else if p.currentToken().Kind == "for" && p.currentToken().Text == "for" {
        return p.parseFor()
    } else if p.currentToken().Kind == "while" && p.currentToken().Text == "while" {
        return p.parseWhile()
    } else if p.currentToken().Kind == "repeat" && p.currentToken().Text == "repeat" {
        return p.parseRepeat()
    } else if p.currentToken().Kind == "break" && p.currentToken().Text == "break" {
        return p.parseBreak()
    } else if p.currentToken().Kind == "next" && p.currentToken().Text == "next" {
        return p.parseNext()
    }


    // По умолчанию разбираем как присваивание или выражение
    return p.parseAssignment()
}

// parseArguments разбирает список аргументов функции
func (p *Parser) parseArguments() []ASTNode {
	var args []ASTNode

	for p.currentToken() != nil {

		// Если обнаружена закрывающая скобка, список аргументов завершается
		if p.currentToken() != nil && p.currentToken().Kind == "rpar" {
			break
		}

		// Разбираем аргумент: это может быть выражение или ключ=значение
		arg := p.parseArgument()
		if arg == nil {
			p.errorf("Ожидается выражение или ключ=значение в списке аргументов")
		}
		args = append(args, arg)


		// Если следующий токен - запятая, пропускаем её и продолжаем
		if p.currentToken() != nil && p.currentToken().Kind == "comma" {
			p.consumeToken()
		} else {
			break
		}
	}

	return args
}

func (p *Parser) parseArgument() ASTNode {
	// Проверяем, является ли текущий токен идентификатором
	if p.currentToken() != nil && p.currentToken().Kind == "ident" {
		ident := &IdentifierNode{Name: p.currentToken().Text}
		savePos := p.pos
		p.consumeToken()

		// Проверяем, идет ли за идентификатором оператор присваивания внутри вызова функции
		if p.currentToken() != nil && p.currentToken().Kind == "assignm_in_func" && p.currentToken().Text == "=" {
			p.consumeToken() // Пропускаем '='
			value := p.parseExpression()
			return &AssignNode{Left: ident, Right: value}
		}

		// Если это не присваивание, откатываемся назад
		p.pos = savePos
	}

	// Если это просто выражение (например, без присваивания), разбираем его
	return p.parseExpression()
}

// parseFor разбирает цикл for
func (p *Parser) parseFor() ASTNode {
    if p.currentToken().Kind != "for" || p.currentToken().Text != "for" {
        p.errorf("Ожидается ключевое слово 'for'")
    }
    p.consumeToken() // Пропускаем 'for'

    if p.currentToken().Kind != "lpar" {
        p.errorf("Ожидается '(' после 'for'")
    }
    p.consumeToken() // Пропускаем '('

    // Разбираем переменную цикла
    variable := p.parsePrimary()

    if p.currentToken().Kind != "in" || p.currentToken().Text != "in" {
        p.errorf("Ожидается ключевое слово 'in'")
    }
    p.consumeToken() // Пропускаем 'in'

    // Разбираем диапазон (например, 1:10)
    rangeExpr := p.parseExpression()

    if p.currentToken().Kind != "rpar" {
        p.errorf("Ожидается ')' после диапазона")
    }
    p.consumeToken() // Пропускаем ')'

    // Разбираем тело цикла
    body := p.parseBlock()

    return &ForNode{
        Variable: variable,
        Range:    rangeExpr,
        Body:     body,
    }
}

// parseWhile разбирает цикл while
func (p *Parser) parseWhile() ASTNode {
    if p.currentToken().Kind != "while" || p.currentToken().Text != "while" {
        p.errorf("Ожидается ключевое слово 'while'")
    }
    p.consumeToken() // Пропускаем 'while'

    if p.currentToken().Kind != "lpar" {
        p.errorf("Ожидается '(' после 'while'")
    }
    p.consumeToken() // Пропускаем '('

    // Разбираем условие
    condition := p.parseExpression()

    if p.currentToken().Kind != "rpar" {
        p.errorf("Ожидается ')' после условия")
    }
    p.consumeToken() // Пропускаем ')'

    // Разбираем тело цикла
    body := p.parseBlock()

    return &WhileNode{
        Condition: condition,
        Body:      body,
    }
}

// parseRepeat разбирает цикл repeat
func (p *Parser) parseRepeat() ASTNode {
    if p.currentToken().Kind != "repeat" || p.currentToken().Text != "repeat" {
        p.errorf("Ожидается ключевое слово 'repeat'")
    }
    p.consumeToken() // Пропускаем 'repeat'

    // Разбираем тело цикла
    body := p.parseBlock()

    return &RepeatNode{
        Body: body,
    }
}

// parseNext разбирает ключевое слово next
func (p *Parser) parseNext() ASTNode {
    if p.currentToken().Kind != "next" || p.currentToken().Text != "next" {
        p.errorf("Ожидается ключевое слово 'next'")
    }
    p.consumeToken() // Пропускаем 'next'
    return &NextNode{}
}

// parseBreak разбирает ключевое слово break
func (p *Parser) parseBreak() ASTNode {
    if p.currentToken().Kind != "break" || p.currentToken().Text != "break" {
        p.errorf("Ожидается ключевое слово 'break'")
    }
    p.consumeToken() // Пропускаем 'next'
    return &BreakNode{}
}

// parseFunctionDeclaration разбирает объявление функции
func (p *Parser) parseFunctionDeclaration() ASTNode {
    // Пропускаем ключевое слово "function"
    if p.currentToken().Kind != "function" || p.currentToken().Text != "function" {
        p.errorf("Ожидается ключевое слово 'function'")
    }
    p.consumeToken() // Пропускаем 'function'

    // Проверяем наличие открывающей скобки '('
    if p.currentToken().Kind != "lpar" {
        p.errorf("Ожидается '(' после 'function'")
    }
    p.consumeToken() // Пропускаем '('

    // Разбираем список параметров
    parameters := p.parseParameters()

    // Проверяем наличие закрывающей скобки ')'
    if p.currentToken().Kind != "rpar" {
        p.errorf("Ожидается ')' после списка параметров")
    }
    p.consumeToken() // Пропускаем ')'

    // Разбираем тело функции
    body := p.parseBlock()

    return &FunctionNode{
        Parameters: parameters,
        Body:       body,
    }
}

// parseParameters разбирает список параметров функции
func (p *Parser) parseParameters() []ASTNode {
    var parameters []ASTNode

    for p.currentToken() != nil && p.currentToken().Kind != "rpar" {
        if p.currentToken().Kind == "ident" {
            // Читаем имя параметра
            paramName := &IdentifierNode{Name: p.currentToken().Text}
            p.consumeToken()

            // Проверяем, есть ли значение по умолчанию
            var defaultValue ASTNode
            if p.currentToken() != nil && p.currentToken().Kind == "assignm_in_func" && p.currentToken().Text == "=" {
                p.consumeToken() // Пропускаем '='
                defaultValue = p.parseExpression() // Разбираем значение по умолчанию
            }

            // Создаем узел параметра
            param := &ParameterNode{
                Name:  paramName,
                Value: defaultValue,
            }
            parameters = append(parameters, param)
        }

        // Если следующий токен - запятая, пропускаем её и продолжаем
        if p.currentToken() != nil && p.currentToken().Kind == "comma" {
            p.consumeToken()
        } else if p.currentToken() != nil && p.currentToken().Kind != "rpar" {
            // Если токен не запятая и не закрывающая скобка, это ошибка
            p.errorf("Ожидается ',' или ')', но получен токен %s (текст: '%s')", 
                p.currentToken().Kind, p.currentToken().Text)
        }
    }

    return parameters
}


func (p *Parser) errorf(format string, args ...interface{}) {
    var line, column int
    
    // Пытаемся получить текущий токен
    if current := p.currentToken(); current != nil {
        line = current.Line
        column = current.Column
    } else if len(p.tokens) > 0 {
        // Берем последний токен, если текущего нет
        last := p.tokens[len(p.tokens)-1]
        line = last.Line
        column = last.Column + len(last.Text) // Позиция после последнего символа токена
    } else {
        // Если токенов нет вообще (пустой ввод)
        line = 1
        column = 1
    }
    
    msg := fmt.Sprintf(format, args...)
    panic(fmt.Sprintf("[%d:%d] %s", line, column, msg))
}