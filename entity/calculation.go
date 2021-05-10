package entity

type Calculation struct {
	Expr   string  `json:"expr"`
	Result float64 `json:"result"`
}
