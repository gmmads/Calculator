package lexer

type Token struct {
	TokenType string
	Value     int
}

type Lexer interface {
	Lex(expr string) ([]Token, error)
}
