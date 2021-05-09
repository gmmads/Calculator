package calculator

import (
	"github.com/gmmads/Calculator/entity"
	"github.com/gmmads/Calculator/interpreter"
	"github.com/gmmads/Calculator/lexer"
	"github.com/gmmads/Calculator/parser"
	"github.com/gmmads/Calculator/repository"
)

type BasicCalculator struct{}

var (
	repo      repository.CalculationRepository
	lex       lexer.Lexer
	parse     parser.Parser
	interpret interpreter.Interpreter
)

// A calculator that uses a Lexer, a Parser, and an Interpreter to evaluate expressions.
// Uses a repository to keep a history of users' calculations.
func NewBasicCalculator(repository repository.CalculationRepository) Calculator {
	repo = repository
	lex = lexer.NewCalcLexer()
	parse = parser.NewCalcParser()
	interpret = interpreter.NewCalcInterpreter()
	return &BasicCalculator{}
}

func (calc *BasicCalculator) Validate(expr string) error {
	_, err := calc.Evaluate(expr)
	return err
}

func (*BasicCalculator) Evaluate(expr string) (*entity.Calculation, error) {
	tokens, err1 := lex.Lex(expr)
	if err1 != nil {
		return nil, err1
	}
	ast, err2 := parse.Parse(tokens)
	if err2 != nil {
		return nil, err2
	}
	result, err3 := interpret.GetValue(ast)
	if err3 != nil {
		return nil, err3
	}
	calculation := entity.Calculation{Expr: expr, Result: int64(result)}
	return repo.Save(&calculation)
}

func (*BasicCalculator) GetHistory() ([]entity.Calculation, error) {
	return repo.FindAll()
}
