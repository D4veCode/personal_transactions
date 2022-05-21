package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID        primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string  `json:"name" bson:"name" binding:"required"`
	Email     string  `json:"email" bson:"email" binding:"required,email"`
}