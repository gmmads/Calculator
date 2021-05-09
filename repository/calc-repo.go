package repository

import (
	"github.com/gmmads/Calculator/entity"
)

type CalculationRepository interface {
	Save(calculation *entity.Calculation) (*entity.Calculation, error)
	FindAll() ([]entity.Calculation, error)
}
