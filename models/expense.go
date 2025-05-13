package models

import (
	"fmt"
	"strings"
	"time"
)

type Expense struct {
	TimeStamp time.Time `bson:"updated_at" json:"updated_at"`
	ItemID    int       `bson:"item_id" json:"item_id"`
	Item      string    `bson:"item" json:"item"`
	Price     float64   `bson:"price" json:"price"`
	Lent      string    `bson:"lent" json:"lent"`
	Involved  []string  `bson:"involved" json:"involved"`
	SplitType string    `bson:"type" json:"type"`
	Splits    []float64 `bson:"splits" json:"splits"`
}

func (e *Expense) Display() {

	involved := strings.Join(e.Involved, ", ")

	fmt.Printf("Item: %s\nPrice: %f\nLent By: %s\nBorrowed by: %s\n", e.Item, e.Price, e.Lent, involved)
}
