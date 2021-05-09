package calculator

import "github.com/gmmads/Calculator/entity"

type Calculator interface {
	Validate(expr string) error
	Evaluate(expr string) (*entity.Calculation, error)
	GetHistory() ([]entity.Calculation, error)
}
