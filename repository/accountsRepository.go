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

var accountCollection *mongo.Collection = database.DBStore.Db.Collection("accounts")

func CreateAccount(account models.Account) error {
	_, err := accountCollection.InsertOne(context.TODO(), account)

	if err != nil {
		log.Fatal("Error creating account", err)
		return err
	}

	return nil
}

func FindAccountById(id string) (models.Account, error) {
	var account models.Account
	objId, _ := primitive.ObjectIDFromHex(id)
	err := accountCollection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&account)

	if err != nil {
		log.Fatal("Error finding account", err)
		return models.Account{}, err
	}

	return account, nil
}

func FindAccountByEmail(email string) (models.Account, error) {
	var account models.Account
	err := accountCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&account)

	if err != nil {
		log.Fatal("Error finding account", err)
		return models.Account{}, err
	}

	return account, nil
}