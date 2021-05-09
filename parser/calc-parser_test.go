package parser

import (
	"testing"

	"github.com/gmmads/Calculator/constants"
	"github.com/gmmads/Calculator/lexer"
	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	parser := NewCalcParser()
	tokens := []lexer.Token{
		{TokenType: constants.NUMBER, Value: 3},
	}

	ast, err := parser.Parse(tokens)
	assert.Nil(t, err)
	assert.NotNil(t, ast)
	expected := &AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 3}}
	assert.Equal(t, expected, ast)
}

func TestTwoNumbersError(t *testing.T) {
	parser := NewCalcParser()
	tokens := []lexer.Token{
		{TokenType: constants.NUMBER, Value: 3},
		{TokenType: constants.NUMBER, Value: 42},
	}

	_, err := parser.Parse(tokens)
	assert.NotNil(t, err)
}

func TestAddition(t *testing.T) {
	parser := NewCalcParser()
	tokens := []lexer.Token{
		{TokenType: constants.NUMBER, Value: 1},
		{TokenType: constants.PLUS},
		{TokenType: constants.NUMBER, Value: 2},
	}

	ast, err := parser.Parse(tokens)
	assert.Nil(t, err)
	assert.NotNil(t, ast)

	left := &AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 1}}
	right := &AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 2}}
	expected := &AstNode{NodeType: constants.BINOP, Token: lexer.Token{TokenType: constants.PLUS}, Children: []*AstNode{left, right}}
	assert.Equal(t, expected, ast)
}

func TestMult(t *testing.T) {
	parser := NewCalcParser()
	tokens := []lexer.Token{
		{TokenType: constants.NUMBER, Value: 5},
		{TokenType: constants.MULT},
		{TokenType: constants.NUMBER, Value: 43},
	}

	ast, err := parser.Parse(tokens)
	assert.Nil(t, err)
	assert.NotNil(t, ast)

	left := &AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 5}}
	right := &AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 43}}
	expected := &AstNode{NodeType: constants.BINOP, Token: lexer.Token{TokenType: constants.MULT}, Children: []*AstNode{left, right}}
	assert.Equal(t, expected, ast)
}

func TestPrecedence(t *testing.T) {
	parser := NewCalcParser()
	tokens := []lexer.Token{
		{TokenType: constants.NUMBER, Value: 5},
		{TokenType: constants.MINUS},
		{TokenType: constants.NUMBER, Value: 10},
		{TokenType: constants.DIV},
		{TokenType: constants.NUMBER, Value: 2},
	}

	ast, err := parser.Parse(tokens)
	assert.Nil(t, err)
	assert.NotNil(t, ast)

	rootLeft := &AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 5}}

	divLeft := &AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 10}}
	divRight := &AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 2}}

	rootRight := &AstNode{NodeType: constants.BINOP, Token: lexer.Token{TokenType: constants.DIV}, Children: []*AstNode{divLeft, divRight}}

	expected := &AstNode{NodeType: constants.BINOP, Token: lexer.Token{TokenType: constants.MINUS}, Children: []*AstNode{rootLeft, rootRight}}
	assert.Equal(t, expected, ast)
}

func TestParentheses(t *testing.T) {
	parser := NewCalcParser()
	tokens := []lexer.Token{
		{TokenType: constants.LPAREN},
		{TokenType: constants.NUMBER, Value: 5},
		{TokenType: constants.MINUS},
		{TokenType: constants.NUMBER, Value: 10},
		{TokenType: constants.RPAREN},
		{TokenType: constants.DIV},
		{TokenType: constants.NUMBER, Value: 2},
	}

	ast, err := parser.Parse(tokens)
	assert.Nil(t, err)
	assert.NotNil(t, ast)

	minusLeft := &AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 5}}
	minusRight := &AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 10}}
	rootLeft := &AstNode{NodeType: constants.BINOP, Token: lexer.Token{TokenType: constants.MINUS}, Children: []*AstNode{minusLeft, minusRight}}

	rootRight := &AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 2}}

	expected := &AstNode{NodeType: constants.BINOP, Token: lexer.Token{TokenType: constants.DIV}, Children: []*AstNode{rootLeft, rootRight}}
	assert.Equal(t, expected, ast)
}

func TestUminus(t *testing.T) {
	parser := NewCalcParser()
	tokens := []lexer.Token{
		{TokenType: constants.MINUS},
		{TokenType: constants.NUMBER, Value: 2},
	}

	ast, err := parser.Parse(tokens)
	assert.Nil(t, err)
	assert.NotNil(t, ast)

	child := &AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 2}}
	expected := &AstNode{NodeType: constants.UMINUS, Token: lexer.Token{TokenType: constants.MINUS}, Children: []*AstNode{child}}

	assert.Equal(t, expected, ast)
}

func TestExtraRPARENError(t *testing.T) {
	parser := NewCalcParser()
	tokens := []lexer.Token{
		{TokenType: constants.LPAREN},
		{TokenType: constants.NUMBER, Value: 5},
		{TokenType: constants.RPAREN},
		{TokenType: constants.RPAREN},
	}

	_, err := parser.Parse(tokens)
	assert.NotNil(t, err)
}

func TestExtraLPARENError(t *testing.T) {
	parser := NewCalcParser()
	tokens := []lexer.Token{
		{TokenType: constants.LPAREN},
		{TokenType: constants.LPAREN},
		{TokenType: constants.NUMBER, Value: 5},
		{TokenType: constants.RPAREN},
	}

	_, err := parser.Parse(tokens)
	assert.NotNil(t, err)
}
