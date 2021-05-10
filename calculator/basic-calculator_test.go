package calculator

import (
	"testing"

	"github.com/gmmads/Calculator/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(calculation *entity.Calculation) (*entity.Calculation, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Calculation), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Calculation, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Calculation), args.Error(1)
}

func TestGetHistory(t *testing.T) {
	mockRepo := new(MockRepository)
	calc := NewBasicCalculator(mockRepo)

	// Mock setup of repository
	calculation := entity.Calculation{Expr: "2+2", Result: 4.0}
	mockRepo.On("FindAll").Return([]entity.Calculation{calculation}, nil)

	result, _ := calc.GetHistory()

	mockRepo.AssertExpectations(t)
	assert.Equal(t, "2+2", result[0].Expr)
	assert.Equal(t, 4.0, result[0].Result)
}

func TestEvaluateLexerError(t *testing.T) {
	mockRepo := new(MockRepository)
	calc := NewBasicCalculator(mockRepo)

	_, err := calc.Evaluate("042 + 3")
	assert.NotNil(t, err)
}

func TestEvaluateParserError(t *testing.T) {
	mockRepo := new(MockRepository)
	calc := NewBasicCalculator(mockRepo)

	_, err := calc.Evaluate("(1 + 3")

	assert.NotNil(t, err)
}

func TestEvaluateInterpreterError(t *testing.T) {
	mockRepo := new(MockRepository)
	calc := NewBasicCalculator(mockRepo)

	_, err := calc.Evaluate("7/0")

	assert.NotNil(t, err)
}

func TestEvaluate(t *testing.T) {
	mockRepo := new(MockRepository)
	calc := NewBasicCalculator(mockRepo)

	calculation := entity.Calculation{Expr: "21*2", Result: 42.0}

	mockRepo.On("Save").Return(&calculation, nil)

	result, err := calc.Evaluate("21*2")

	mockRepo.AssertExpectations(t)

	assert.Nil(t, err)
	assert.Equal(t, "21*2", result.Expr)
	assert.Equal(t, 42.0, result.Result)
}
