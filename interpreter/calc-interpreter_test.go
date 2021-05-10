package interpreter

import (
	"testing"

	"github.com/gmmads/Calculator/constants"
	"github.com/gmmads/Calculator/lexer"
	"github.com/gmmads/Calculator/parser"
	"github.com/stretchr/testify/assert"
)

func TestConstant(t *testing.T) {
	interpreter := NewCalcInterpreter()
	ast := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 3}}
	result, err := interpreter.GetValue(ast)
	assert.Nil(t, err)
	assert.Equal(t, 3.0, result)
}

func TestUminus(t *testing.T) {
	interpreter := NewCalcInterpreter()
	child := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 3}}
	ast := &parser.AstNode{NodeType: constants.UMINUS, Token: lexer.Token{TokenType: constants.MINUS}, Children: []*parser.AstNode{child}}
	result, err := interpreter.GetValue(ast)
	assert.Nil(t, err)
	assert.Equal(t, -3.0, result)
}

func TestAddition(t *testing.T) {
	interpreter := NewCalcInterpreter()
	left := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 20}}
	right := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 22}}
	ast := &parser.AstNode{NodeType: constants.PLUS, Token: lexer.Token{TokenType: constants.PLUS}, Children: []*parser.AstNode{left, right}}
	result, err := interpreter.GetValue(ast)
	assert.Nil(t, err)
	assert.Equal(t, 42.0, result)
}

func TestSubtraction(t *testing.T) {
	interpreter := NewCalcInterpreter()
	left := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 20}}
	right := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 22}}
	ast := &parser.AstNode{NodeType: constants.MINUS, Token: lexer.Token{TokenType: constants.MINUS}, Children: []*parser.AstNode{left, right}}
	result, err := interpreter.GetValue(ast)
	assert.Nil(t, err)
	assert.Equal(t, -2.0, result)
}

func TestMult(t *testing.T) {
	interpreter := NewCalcInterpreter()
	left := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 20}}
	right := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 22}}
	ast := &parser.AstNode{NodeType: constants.MULT, Token: lexer.Token{TokenType: constants.MULT}, Children: []*parser.AstNode{left, right}}
	result, err := interpreter.GetValue(ast)
	assert.Nil(t, err)
	assert.Equal(t, 20.0*22.0, result)
}

func TestDivision(t *testing.T) {
	interpreter := NewCalcInterpreter()
	left := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 20}}
	right := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 22}}
	ast := &parser.AstNode{NodeType: constants.DIV, Token: lexer.Token{TokenType: constants.DIV}, Children: []*parser.AstNode{left, right}}
	result, err := interpreter.GetValue(ast)
	assert.Nil(t, err)
	assert.Equal(t, 20.0/22.0, result)
}

func TestDivisionByZero(t *testing.T) {
	interpreter := NewCalcInterpreter()
	left := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 7}}
	right := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 0}}
	ast := &parser.AstNode{NodeType: constants.DIV, Token: lexer.Token{TokenType: constants.DIV}, Children: []*parser.AstNode{left, right}}
	_, err := interpreter.GetValue(ast)
	assert.NotNil(t, err)
	assert.Equal(t, "error: Division by zero", err.Error())
}

func TestComplicatedExpression(t *testing.T) {
	interpreter := NewCalcInterpreter()
	one := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 1}}
	two := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 2}}
	three := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 3}}
	four := &parser.AstNode{NodeType: constants.NUMBER, Token: lexer.Token{TokenType: constants.NUMBER, Value: 4}}

	plus := &parser.AstNode{NodeType: constants.PLUS, Token: lexer.Token{TokenType: constants.PLUS}, Children: []*parser.AstNode{one, two}}
	uminus := &parser.AstNode{NodeType: constants.UMINUS, Token: lexer.Token{TokenType: constants.UMINUS}, Children: []*parser.AstNode{three}}
	mult := &parser.AstNode{NodeType: constants.MULT, Token: lexer.Token{TokenType: constants.MULT}, Children: []*parser.AstNode{uminus, four}}
	ast := &parser.AstNode{NodeType: constants.MINUS, Token: lexer.Token{TokenType: constants.MINUS}, Children: []*parser.AstNode{plus, mult}}

	result, err := interpreter.GetValue(ast)
	assert.Nil(t, err)
	assert.Equal(t, (1.0+2.0)-((-3.0)*4.0), result)
}
