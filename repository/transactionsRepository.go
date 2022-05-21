package repository

import (
	"context"
	"log"

	"github.com/D4vecode/personal_transactions/database"
	"github.com/D4vecode/personal_transactions/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var transactionsCollection *mongo.Collection = database.DBStore.Db.Collection("transactions")


func SaveTransactions(transactions []*models.Transaction) error {
	transactionDocuments := make([]interface{}, len(transactions))
	for i, transaction := range transactions {
		transactionDocuments[i] = transaction
	}

	_, err := transactionsCollection.InsertMany(context.TODO(), transactionDocuments)

	if err != nil {
		log.Fatal("Error inserting documents", err)
		return err
	}

	return nil
}

func AllTransactionsByAccount (accountId primitive.ObjectID) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	cursor, err := transactionsCollection.Find(context.TODO(), bson.M{"account_id": accountId})

	if err != nil {
		log.Fatal("Error finding transactions", err)
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var transaction models.Transaction
		err := cursor.Decode(&transaction)

		if err != nil {
			log.Fatal("Error decoding transaction", err)
			return nil, err
		}

		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}