package history

import (
	"testing"

	"github.com/gmmads/Calculator/entity"
	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	history := NewCalcHistory()
	calculation := &entity.Calculation{Expr: "2+2", Result: 4.0}

	result, err := history.Save(calculation)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, calculation, result)
}

func TestFindAll(t *testing.T) {
	history := NewCalcHistory()
	result, err := history.FindAll()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, make([]entity.Calculation, 0), result)
}

func TestFindAllWithElements(t *testing.T) {
	history := NewCalcHistory()
	calc := entity.Calculation{Expr: "1+2*3", Result: 7.0}
	calculations = []entity.Calculation{calc}

	result, err := history.FindAll()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, calculations, result)
}
