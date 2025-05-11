package utils

import (
	"encoding/json"
	"ledger/models"
	"os"
)

func WriteTransactions(transactions []models.Transaction) error {
	file, err := os.Create("data/transactions.json")

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "\t")
	if err := encoder.Encode(transactions); err != nil {
		return err
	}

	return nil

}
