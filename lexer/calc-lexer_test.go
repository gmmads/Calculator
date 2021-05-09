package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyExprError(t *testing.T) {
	lexer := NewCalcLexer()
	expr := ""

	result, err := lexer.Lex(expr)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "the expression is empty", err.Error())
}

func TestIllegalSymbol(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "A"

	result, err := lexer.Lex(expr)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "lexing error. Unexpected symbol: 'A'", err.Error())
}

func TestPlus(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "+"

	expectedResult := []Token{{TokenType: "+"}}
	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

func TestMinus(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "-"

	expectedResult := []Token{{TokenType: "-"}}
	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

func TestMult(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "*"

	expectedResult := []Token{{TokenType: "*"}}
	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

func TestDiv(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "/"

	expectedResult := []Token{{TokenType: "/"}}
	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

func TestLPAREN(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "("

	expectedResult := []Token{{TokenType: "("}}
	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

func TestRPAREN(t *testing.T) {
	lexer := NewCalcLexer()
	expr := ")"

	expectedResult := []Token{{TokenType: ")"}}
	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

func TestMultipleTokens(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "+(*"

	expectedResult := []Token{
		{TokenType: "+"},
		{TokenType: "("},
		{TokenType: "*"},
	}

	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

func TestWhitespace(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "+ ( *"

	expectedResult := []Token{
		{TokenType: "+"},
		{TokenType: "("},
		{TokenType: "*"},
	}

	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

func TestDigit(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "3"

	expectedResult := []Token{{TokenType: "NUM", Value: 3}}
	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

func TestZero(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "0"

	expectedResult := []Token{{TokenType: "NUM", Value: 0}}
	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

func TestBiggerNumber(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "42"

	expectedResult := []Token{{TokenType: "NUM", Value: 42}}
	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

func TestExpressionWithZero(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "0 + 3"

	expectedResult := []Token{{TokenType: "NUM", Value: 0}, {TokenType: "+"}, {TokenType: "NUM", Value: 3}}
	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}

func TestNumberStartingWithZero(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "042"

	result, err := lexer.Lex(expr)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "lexing error. Unexpected symbol: '0'", err.Error())
}

func TestComplicatedExpression(t *testing.T) {
	lexer := NewCalcLexer()
	expr := "11 + 2 * (33-24)"

	expectedResult := []Token{
		{TokenType: "NUM", Value: 11},
		{TokenType: "+"},
		{TokenType: "NUM", Value: 2},
		{TokenType: "*"},
		{TokenType: "("},
		{TokenType: "NUM", Value: 33},
		{TokenType: "-"},
		{TokenType: "NUM", Value: 24},
		{TokenType: ")"},
	}

	result, err := lexer.Lex(expr)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedResult, result)
}
