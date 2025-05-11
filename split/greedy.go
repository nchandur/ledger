package split

import (
	"ledger/models"
	"math"
	"sort"
)

func SimplifyDebts(balances map[string]float64) []models.Transaction {
	const epsilon = 1e-9
	var transactions []models.Transaction

	type person struct {
		Name    string
		Balance float64
	}

	for {
		var creditors []person
		var debtors []person

		for name, balance := range balances {
			if math.Abs(balance) < epsilon {
				continue
			}
			if balance > 0 {
				creditors = append(creditors, person{name, balance})
			} else {
				debtors = append(debtors, person{name, balance})
			}
		}

		if len(creditors) == 0 || len(debtors) == 0 {
			break
		}

		sort.Slice(creditors, func(i, j int) bool {
			return creditors[i].Balance > creditors[j].Balance
		})
		sort.Slice(debtors, func(i, j int) bool {
			return debtors[i].Balance < debtors[j].Balance
		})

		creditor := creditors[0]
		debtor := debtors[0]

		amount := math.Min(creditor.Balance, -debtor.Balance)
		amount = math.Round(amount*1000000) / 1000000

		transactions = append(transactions, models.Transaction{
			From:   debtor.Name,
			To:     creditor.Name,
			Amount: amount,
		})

		balances[creditor.Name] -= amount
		balances[debtor.Name] += amount
	}

	return transactions
}
