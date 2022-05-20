package models

import "time"

type Transaction struct {
	ID        int  `json:"id" bson:"id"`
	Date      time.Time `json:"date" bson:"date"`
	Amount    float64  `json:"amount" bson:"amount"`
  Type 		  string `json:"type" bson:"type"`
}