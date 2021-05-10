package history

import "github.com/gmmads/Calculator/entity"

type History interface {
	Save(calculation *entity.Calculation) (*entity.Calculation, error)
	FindAll() ([]entity.Calculation, error)
}
