package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Db *mongo.Database
	Client *mongo.Client
}

func ConnectDB() *Store {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Error connecting to database", err)
		return nil
	}
	database := client.Database("personal_transactions")

	return &Store{
		Db: database,
		Client: client,
	}
}