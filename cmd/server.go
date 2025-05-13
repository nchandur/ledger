package main

import (
	"ledger/api"
	"ledger/db"
	"log"
)

func main() {

	err := db.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.DisconnectDB()

	r := api.SetUpRouter()

	log.Println("Server running at port :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
