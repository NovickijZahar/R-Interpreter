package lexer


type TokenType struct {
	Name string
	Regex string
	Class string
}


var keywords = []TokenType {
	{"if", `if`, "keyword"},
	{"else", `else`, "keyword"},
	{"while", `while`, "keyword"},
	{"repeat", `repeat`, "keyword"},
	{"in", `in`, "keyword"},
	{"for", `for`, "keyword"},
	{"function", `function`, "keyword"},
	{"next", `next`, "keyword"},
	{"break", `break`, "keyword"},
	{"NULL", `NULL`, "keyword"},
	{"Inf", `Inf`, "keyword"},
	{"NaN", `NaN`, "keyword"},
	{"NA", `NA`, "keyword"},
}

var operators = []TokenType {
	{"assignment", `(<-)`, "operator"},
	{"comparison", `(==|!=|<=|>=|>|<)`, "operator"},
	{"assignm_in_func", `=`, "operator"},
	{"logical", `(&&|\|\||&|\||!)`, "operator"},
	{"miscellaneous", `(:|%in%|%\*%)`, "operator"},
	{"arithmetic", `(\+|-|\*|\/|\^|%%|%\/%)`, "operator"},
	{"ellipsis", `(\.\.\.)`, "operator"},
}

var variables = []TokenType {
	{"ident", "([a-zA-Z_]\\w*|`%.*%`|%.*%)", "variable"},
}

var constants = []TokenType {
	{"complex", `[+-]?(\d+(\.\d*)?|(\.\d+))?\s*[+-]?\s*(\d+(\.\d*)?|(\.\d+))i`, "constant"},
	{"integer", `[+-]?\d+L`, "constant"},
	{"numeric", `[+-]?(\d+(\.\d*)?|\.\d+)`, "constant"},
	{"logical", `(TRUE|T|FALSE|F)`, "constant"}, 
	{"character", `("[^"]*"|'[^']*')`, "constant"},
}

var punctuations = []TokenType {
	{"lpar", `\(`, "punctuation"},
	{"rpar", `\)`, "punctuation"},
	{"begin", `{`, "punctuation"},
	{"end", `}`, "punctuation"},
	{"comma", `,`, "punctuation"},
	{"access", `(\.|@|\$)`, "punctuation"},
}

var skip = []TokenType {
	{"SPACE", `\s`, "skip"},
	{"COMMENT", `#.*($|\n)`, "skip"},
}


func getTokenTypesList() []TokenType {
	tokenTypesList := append(keywords, constants...)
	tokenTypesList = append(tokenTypesList, operators...)
	tokenTypesList = append(tokenTypesList, variables...)
	tokenTypesList = append(tokenTypesList, punctuations...)
	tokenTypesList = append(tokenTypesList, skip...)

	return tokenTypesList
}