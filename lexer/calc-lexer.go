package lexer

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gmmads/Calculator/constants"
)

type CalcLexer struct{}

func NewCalcLexer() Lexer {
	return &CalcLexer{}
}

func (*CalcLexer) Lex(expr string) ([]Token, error) {
	if expr == "" {
		err := errors.New("the expression is empty")
		return nil, err
	}
	runes := []rune(expr)
	var tokens []Token
	i := 0

	for i < len(runes) {
		var token Token
		switch r := runes[i]; r {
		case ' ': // Whitespace
			i++
			continue
		case '+':
			token.TokenType = constants.PLUS
		case '-':
			token.TokenType = constants.MINUS
		case '*':
			token.TokenType = constants.MULT
		case '/':
			token.TokenType = constants.DIV
		case '(':
			token.TokenType = constants.LPAREN
		case ')':
			token.TokenType = constants.RPAREN
		default:
			if r == '0' && (i+1 >= len(runes) || !isDigit(runes[i+1])) {
				token.TokenType = constants.NUMBER
				token.Value = 0
				tokens = append(tokens, token)
				i++
				continue
			}
			if !isNonZeroDigit(r) {
				return nil, errors.New("lexing error. Unexpected symbol: '" + string(r) + "'")
			}
			// Else, it is a non-zero number
			str, newIndex := lexNumber(runes, i)
			value, err := strconv.Atoi(str)
			if err != nil {
				return nil, errors.New("lexing error. Error lexing number: " + err.Error())
			}
			token.TokenType = constants.NUMBER
			token.Value = value
			tokens = append(tokens, token)
			i = newIndex
			continue
		}

		tokens = append(tokens, token)
		i++
	}

	return tokens, nil
}

func isNonZeroDigit(r rune) bool {
	return r == '1' ||
		r == '2' ||
		r == '3' ||
		r == '4' ||
		r == '5' ||
		r == '6' ||
		r == '7' ||
		r == '8' ||
		r == '9'
}

func isDigit(r rune) bool {
	return r == '0' || isNonZeroDigit(r)
}

// Starts at runes[i], and lexes a number starting from there. Returns a string representing the number, and the new index
func lexNumber(runes []rune, i int) (string, int) {
	var sb strings.Builder
	for i < len(runes) && isDigit(runes[i]) {
		sb.WriteRune(runes[i])
		i++
	}
	return sb.String(), i
}
