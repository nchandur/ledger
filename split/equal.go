package split

import (
	"fmt"
	"ledger/models"
	"math"
)

func SplitEqual(expense models.Expense, people models.People) error {

	numPeople := len(expense.Involved)

	if numPeople == 0 {
		return fmt.Errorf("no people invovled")
	}

	if expense.Price < 0 {
		return fmt.Errorf("invalid price")
	}

	cents := int(math.Round(expense.Price * 100))
	baseShare := cents / numPeople
	remainder := cents % numPeople

	for idx, invovled := range expense.Involved {
		share := baseShare

		if idx < remainder {
			share += 1
		}

		if _, ok := people[invovled]; ok {
			people[invovled] -= float64(share) / 100
		} else {
			people[invovled] = -float64(share) / 100
		}
	}

	if _, ok := people[expense.Lent]; ok {
		people[expense.Lent] += float64(cents) / 100
	} else {
		people[expense.Lent] = float64(cents) / 100
	}

	return nil
}
