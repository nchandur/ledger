package utils

import (
	"encoding/json"
	"io"
	"ledger/models"
	"os"
)

func ReadExpenseReport(filepath string) ([]models.Expense, error) {

	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	byteValue, err := io.ReadAll(file)

	if err != nil {
		return nil, err
	}

	var expenses []models.Expense

	if err := json.Unmarshal(byteValue, &expenses); err != nil {
		return nil, err
	}

	return expenses, nil
}
