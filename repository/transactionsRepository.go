package repository

import (
	"context"
	"log"
	"fmt"
	"github.com/D4vecode/personal_transactions/database"
	"github.com/D4vecode/personal_transactions/models"
)

func SaveTransactions(store *database.Store, transactions []*models.Transaction) error {
	collection := store.Db.Collection("transactions")
	transactionDocuments := make([]interface{}, len(transactions))

	for i, transaction := range transactions {
		fmt.Println(transaction)
		transactionDocuments[i] = transaction
	}

	_, err := collection.InsertMany(context.TODO(), transactionDocuments)

	if err != nil {
		log.Fatal("Error inserting documents", err)
		return err
	}

	return nil
}