package main

import (
	"fmt"
	u "github.com/D4vecode/personal_transactions/utils"
)

func main() {
	transactions, err := u.ReadCSV("./data/transactions.csv", ",")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(transactions)
	fmt.Println("Hello, world!")
}