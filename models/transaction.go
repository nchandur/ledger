package models

import "fmt"

type Transaction struct {
	From   string  `bson:"from" json:"from"`
	To     string  `bson:"to" json:"to"`
	Amount float64 `bson:"amount" json:"amount"`
}

func (t *Transaction) Display() {
	fmt.Printf("%s pays %s $%f\n", t.From, t.To, t.Amount)
}
