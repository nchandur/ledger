package models

type Expense struct {
	Item      string    `json:"item"`
	Price     float64   `json:"price"`
	Lent      string    `json:"lent"`
	Involved  []string  `json:"involved"`
	SplitType string    `json:"type"`
	Splits    []float64 `json:"splits"`
}
