package interpreter

import (
	"errors"

	"github.com/gmmads/Calculator/constants"
	"github.com/gmmads/Calculator/parser"
)

type CalcInterpreter struct{}

func NewCalcInterpreter() Interpreter {
	return &CalcInterpreter{}
}

func (interpreter *CalcInterpreter) GetValue(ast *parser.AstNode) (float64, error) {
	switch ast.NodeType {
	case constants.NUMBER:
		return float64(ast.Token.Value), nil
	case constants.UMINUS:
		childValue, err := interpreter.GetValue(ast.Children[0])
		if err != nil {
			return 0, err
		}
		return -childValue, nil
	default:
		// Binop
		left, err1 := interpreter.GetValue(ast.Children[0])
		if err1 != nil {
			return 0, err1
		}
		right, err2 := interpreter.GetValue(ast.Children[1])
		if err2 != nil {
			return 0, err2
		}
		switch ast.Token.TokenType {
		case constants.PLUS:
			return left + right, nil
		case constants.MINUS:
			return left - right, nil
		case constants.MULT:
			return left * right, nil
		default:
			// Division
			if right == 0 {
				return 0, errors.New("error: Division by zero")
			}
			return left / right, nil
		}
	}
}
