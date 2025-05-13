package models

import "fmt"

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func (t *Transaction) Display() {
	fmt.Printf("%s pays %s $%f\n", t.From, t.To, t.Amount)
}
