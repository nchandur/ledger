package utils

import (
	"encoding/json"
	"ledger/models"
	"os"
)

func WriteLedger(people models.People) error {

	file, err := os.Create("data/net.json")

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	encoder.SetIndent("", "\t")
	if err := encoder.Encode(people); err != nil {
		return err
	}

	return nil

}
