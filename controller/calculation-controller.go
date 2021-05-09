package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gmmads/Calculator/calculator"
	"github.com/gmmads/Calculator/entity"
	"github.com/gmmads/Calculator/errors"
)

type CalculationController interface {
	GetHistory(resp http.ResponseWriter, req *http.Request)
	AddCalculation(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

var (
	calculate calculator.Calculator
)

func NewCalculationController(calc calculator.Calculator) CalculationController {
	calculate = calc
	return &controller{}
}

func (*controller) GetHistory(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	calculations, err := calculate.GetHistory()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the history"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(calculations)
}

func (*controller) AddCalculation(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var calculation entity.Calculation
	err := json.NewDecoder(req.Body).Decode(&calculation)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling the data"})
		return
	}
	result, err2 := calculate.Evaluate(calculation.Expr)
	if err2 != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err2.Error()})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
