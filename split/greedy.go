package split

import (
	"ledger/models"
	"maps"
	"math"
)

func SettleGreedy(people models.People) []models.Transaction {
	balances := maps.Clone(people)

	var lenders []struct {
		name  string
		value float64
	}

	var borrowers []struct {
		name  string
		value float64
	}

	for key, val := range balances {
		if val > 0 {
			lenders = append(lenders, struct {
				name  string
				value float64
			}{key, val})
		} else if val < 0 {
			borrowers = append(borrowers, struct {
				name  string
				value float64
			}{key, val})
		}
	}

	l, b := 0, 0

	transactions := []models.Transaction{}

	for l < len(lenders) && b < len(borrowers) {
		lender, lent := lenders[l].name, lenders[l].value
		borrower, borrowed := borrowers[b].name, borrowers[b].value

		transferAmount := min(lent, -borrowed)

		transactions = append(transactions, models.Transaction{From: borrower, To: lender, Amount: transferAmount})

		lent -= transferAmount
		borrowed += transferAmount

		if math.Abs(lent) < 1e-10 {
			l++
		} else {
			lenders[l] = struct {
				name  string
				value float64
			}{lender, lent}
		}

		if math.Abs(borrowed) < 1e-10 {
			b++
		} else {
			borrowers[b] = struct {
				name  string
				value float64
			}{borrower, borrowed}
		}
	}

	return transactions

}
