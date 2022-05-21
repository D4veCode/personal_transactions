package main

import (
	//"context"
	//"fmt"

	db "github.com/D4vecode/personal_transactions/database"
	"github.com/D4vecode/personal_transactions/routes"

	//"github.com/D4vecode/personal_transactions/repository"
	//u "github.com/D4vecode/personal_transactions/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	db.ConnectDB()

	routes.AccountsRoutes(router)
	routes.TransactionRoutes(router)

	router.Run(":8080")
	
	/*transactions, err := u.ReadCSV("./data/transactions.csv", ",")
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
	}*/

}