package split

import (
	"fmt"
	"ledger/models"
)

func checkPerc(percs []float64) bool {
	var sum float64

	for _, perc := range percs {
		sum += perc
	}

	return sum == 1

}

func SplitPercentages(expense models.Expense, people models.People) error {
	if len(expense.Involved) != len(expense.Splits) {
		return fmt.Errorf("length of involved and percentages do not match")
	}

	if !checkPerc(expense.Splits) {
		return fmt.Errorf("percentages do not add up to 1.00")
	}

	for idx, involved := range expense.Involved {
		borrowed := expense.Splits[idx] * expense.Price

		if _, ok := people[involved]; ok {
			people[involved] -= borrowed
		} else {
			people[involved] = -borrowed
		}

	}

	if _, ok := people[expense.Lent]; ok {
		people[expense.Lent] += expense.Price
	} else {
		people[expense.Lent] = expense.Price
	}
	return nil
}
