package entity

type Calculation struct {
	Expr   string `json:"expr"`
	Result int64  `json:"result"`
}
