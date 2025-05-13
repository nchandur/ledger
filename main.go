package main

import (
	"fmt"
	"ledger/crud"
	"ledger/db"
	"ledger/utils"
	"log"
)

func main() {
	db.ConnectDB()
	defer db.DisconnectDB()

	expenses, err := utils.ReadExpenseReport("data/costco.json")

	if err != nil {
		log.Fatal(err)
	}

	groupName := "costco"
	err = crud.CreateGroup(groupName)

	group, err := crud.AccessGroup(groupName)

	if err != nil {
		log.Fatal(err)
	}

	group.AddItems(expenses)

	fmt.Println("added items")

	err = crud.CalculateNetAndTransactions(groupName)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("calculated balances and transactions")

}
