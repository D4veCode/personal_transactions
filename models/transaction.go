package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID        int  `json:"id" bson:"id"`
	Date      time.Time `json:"date" bson:"date"`
	Amount    float64  `json:"amount" bson:"amount"`
  Type 		  string `json:"type" bson:"type"`
	AccountID primitive.ObjectID `json:"account_id" bson:"account_id"`
}