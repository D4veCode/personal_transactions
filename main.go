package main

import (
	"fmt"
	"context"
	db "github.com/D4vecode/personal_transactions/database"
	"github.com/D4vecode/personal_transactions/repository"
	u "github.com/D4vecode/personal_transactions/utils"
)

func main() {
	transactions, err := u.ReadCSV("./data/transactions.csv", ",")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	processedData := u.ProcessData(transactions)

	store := db.ConnectDB()
	fmt.Println("Database connected")

	err = repository.SaveTransactions(store, processedData)

	if err != nil {
		panic(err)
	}

	defer store.Client.Disconnect(context.TODO())

	data := u.CalculateData(processedData)
	err = u.SendEmail(data)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Hello, world!")
}