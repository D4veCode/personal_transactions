package handlers

import (
	"fmt"

	"github.com/D4vecode/personal_transactions/repository"
	"github.com/D4vecode/personal_transactions/utils"
	"github.com/gin-gonic/gin"
)

type sendEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}


// @description Adds transactions to the database from a csv file
// @version 1.0
// @accept multipart/form-data
// @Param transactions formData file true "transactions"
// @Param account_id formData string true "account_id"
// @Success 200 {object} gin.H {"message": "Transactions successfully uploaded and saved."}
// @router /upload-transactions/ [post]
func AddTransactions() gin.HandlerFunc {

	return func(c *gin.Context) {
		accoundId := c.PostForm("account_id")
		csvFile, csvHeaders, err := c.Request.FormFile("transactions")

		if err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid file",
			})
			return
		}

		if csvHeaders.Header.Values("Content-Type")[0] != "text/csv" {
			c.JSON(400, gin.H{
				"message": "Invalid file, must be a csv file",
			})
			return
		}

		account, err := repository.FindAccountById(accoundId)

		if err != nil {
			c.JSON(404, gin.H{
				"message": "Account not found",
			})
			return
		}

		transactions, err := utils.ReadCSV(csvFile, ",")

		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"message": "Error reading file",
			})
			return 
		}

		processedData := utils.ProcessData(transactions, account.ID)

		repository.SaveTransactions(processedData)

		c.JSON(200, gin.H{
			"message": "Transactions successfully uploaded and saved.",
		})
	}

}

// @description Sends an email from an account
// @version 1.0
// @accept application/json
// @Param email query string true "Email"
// @Success 200 {object} gin.H {"message": "Email sent"}
// @router /send-email/ [post]
func SendTransactionsEmail () gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody sendEmailRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid fields, please check again",
			})
			return
		}

		data, err := repository.FindAccountByEmail(requestBody.Email)

		if err != nil {
			c.JSON(404, gin.H{
				"message": "Account not found",
			})
			return
		}
		
		transactions, err := repository.AllTransactionsByAccount(data.ID)

		calculatedData := utils.CalculateData(transactions)

		utils.SendEmail(calculatedData, data.Email)

		c.JSON(200, gin.H{
			"message": "Email sent",
		})
		return
	}
}