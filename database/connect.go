package database

import (
	"context"
	"fmt"
	"log"

	"github.com/D4vecode/personal_transactions/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	Db *mongo.Database
	Client *mongo.Client
}


func ConnectDB() *Store {

	DB_URI, err := config.GetEnv("DB_URI")

	if err != nil {
		log.Fatal(err)
		return nil
	}


	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DB_URI).SetServerAPIOptions(serverAPIOptions))
	
	if err != nil {
		log.Fatal("Error connecting to database", err)
		return nil
	}
	database := client.Database("personal_transactions")

	fmt.Println("Database connected")

	return &Store{
		Db: database,
		Client: client,
	}
}

var DBStore *Store = ConnectDB()