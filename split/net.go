package split

import (
	"fmt"
	"ledger/models"
	"math"
)

func CalculateNetAmount(expenses []models.Expense, people models.People) error {
	var err error

	for _, expense := range expenses {
		switch expense.SplitType {
		case "equal":
			err = SplitEqual(expense, people)
		case "manual":
			err = SplitManual(expense, people)
		case "percentage":
			err = SplitPercentages(expense, people)
		}

		if err != nil {
			return fmt.Errorf("item: %s, %v", expense.Item, err)
		}
	}

	for key := range people {
		people[key] = math.Round(people[key]*1000000) / 1000000
	}

	return nil
}
