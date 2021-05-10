package interpreter

import (
	"github.com/gmmads/Calculator/parser"
)

type Interpreter interface {
	GetValue(ast *parser.AstNode) (float64, error)
}
