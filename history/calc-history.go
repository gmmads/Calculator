package history

import (
	"github.com/gmmads/Calculator/entity"
)

type CalcHistory struct{}

var (
	calculations []entity.Calculation
)

func NewCalcHistory() History {
	calculations = make([]entity.Calculation, 0)
	return &CalcHistory{}
}

func (*CalcHistory) Save(calculation *entity.Calculation) (*entity.Calculation, error) {
	calculations = append(calculations, *calculation)
	return calculation, nil
}

func (*CalcHistory) FindAll() ([]entity.Calculation, error) {
	return calculations, nil
}
