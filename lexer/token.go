package lexer


type Token struct {
	Kind string
	Text string
	Pos int
	Line int
	Column int
	Class string
	Id string
}
