package split

import (
	"ledger/models"
	"maps"
	"math"
)

func SettleNaive(people models.People) []models.Transaction {
	balances := maps.Clone(people)

	lenders := []string{}
	borrowers := []string{}

	for key, value := range balances {
		if value > 0 {
			lenders = append(lenders, key)
		} else if value < 0 {
			borrowers = append(borrowers, key)
		}
	}

	l, b := 0, 0

	transactions := []models.Transaction{}

	for l < len(lenders) && b < len(borrowers) {
		lender := lenders[l]
		borrower := borrowers[b]

		lent := balances[lender]
		borrowed := balances[borrower]

		transferAmount := min(lent, -borrowed)

		transactions = append(transactions, models.Transaction{From: borrower, To: lender, Amount: transferAmount})

		balances[lender] -= transferAmount
		balances[borrower] += transferAmount

		if math.Abs(balances[lender]) < 1e-10 {
			balances[lender] = 0.0
			l += 1
		}
		if math.Abs(balances[borrower]) < 1e-10 {
			balances[borrower] = 0.0
			b += 1
		}
	}

	return transactions
}
