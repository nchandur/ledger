package main

import (
	"fmt"
	"ledger/models"
	"ledger/split"
	"ledger/utils"
	"log"
)

func main() {

	expenses, err := utils.ReadExpenseReport("data/expense.json")

	if err != nil {
		log.Fatal(err)
	}

	people := make(models.People)

	err = split.CalculateNetAmount(expenses, people)

	if err != nil {
		log.Fatal(err)
	}

	utils.WriteLedger(people)

	fmt.Printf("%v\n\n", people)

	transactions := split.SimplifyDebts(people)

	utils.WriteTransactions(transactions)

	for _, t := range transactions {
		fmt.Printf("%s pays %s: $%f\n", t.From, t.To, t.Amount)
	}

	fmt.Printf("\n\n%v\n", people)

}
