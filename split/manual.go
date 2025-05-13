package split

import (
	"fmt"
	"ledger/models"
	"math"
)

func roundToCent(val float64) float64 {
	val = math.Round(val * 100)

	return val / 100
}

func checkManual(splits []float64, total float64) bool {
	var sum float64

	for _, split := range splits {
		sum += split
	}
	return roundToCent(sum) == roundToCent(total)
}

func SplitManual(expense models.Expense, people models.People) error {
	if len(expense.Involved) != len(expense.Splits) {
		return fmt.Errorf("length of involved and splits do not match")
	}

	if !checkManual(expense.Splits, expense.Price) {
		return fmt.Errorf("price splits do not add up to total.")
	}

	for idx, involved := range expense.Involved {

		if _, ok := people[involved]; ok {
			people[involved] -= expense.Splits[idx]
		} else {
			people[involved] = -expense.Splits[idx]
		}
	}

	if _, ok := people[expense.Lent]; ok {
		people[expense.Lent] += expense.Price
	} else {
		people[expense.Lent] = expense.Price
	}

	return nil
}
