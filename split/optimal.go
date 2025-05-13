package split

import (
	"ledger/models"
	"maps"
	"math"
)

func SettleOptimal(people models.People) []models.Transaction {
	balances := maps.Clone(people)

	getExtremes := func(balances models.People) (string, string) {
		var maxLender string
		maxLent := -math.MaxFloat64

		var maxBorrower string
		maxBorrowed := math.MaxFloat64

		for key, val := range balances {
			if val > 0 && val > maxLent {
				maxLent = val
				maxLender = key
			}
			if val < 0 && val < maxBorrowed {
				maxBorrowed = val
				maxBorrower = key
			}
		}
		return maxLender, maxBorrower
	}

	transactions := []models.Transaction{}

	for {
		lender, borrower := getExtremes(balances)

		if lender == "" || borrower == "" {
			break
		}

		transferAmount := min(balances[lender], -balances[borrower])

		balances[lender] -= transferAmount
		balances[borrower] += transferAmount

		transactions = append(transactions, models.Transaction{From: borrower, To: lender, Amount: transferAmount})

	}

	return transactions

}
