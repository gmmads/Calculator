package history

import (
	"errors"

	"github.com/gmmads/Calculator/entity"
)

type CalcHistory struct{}

var (
	calculations []entity.Calculation
)

func NewCalcHistory() History {
	return &CalcHistory{}
}

func (*CalcHistory) Save(calculation *entity.Calculation) (*entity.Calculation, error) {
	calculations = append(calculations, *calculation)
	return calculation, nil
}

func (*CalcHistory) FindAll() ([]entity.Calculation, error) {
	if len(calculations) == 0 {
		return nil, errors.New("no calculations have been made yet.")
	}
	return calculations, nil
}
