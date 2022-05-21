package handlers

import (
	"github.com/D4vecode/personal_transactions/models"
	"github.com/D4vecode/personal_transactions/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// @description creates an account based in email and full name
// @version 1.0
// @accept multipart/form-data
// @Param email body string true "Email"
// @Param name body string true "Full name"
// @Success 200 {object} gin.H {"status": 200, message: "Account created", data: models.Account{}}
// @router /create-account/ [post]
func CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {

		var account models.Account

		if err := c.ShouldBindJSON(&account); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid fields, please check again",
			})
			return
		}

		newAccount := models.Account{
			ID: primitive.NewObjectID(),
			Name: account.Name,
			Email: account.Email,
		}

		err := repository.CreateAccount(newAccount)

		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error creating account",
			})
			return
		}


		c.JSON(201, gin.H{
			"status": 201,
			"message": "Account created",
			"data": newAccount,
		})
	}
}
