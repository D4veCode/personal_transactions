package routes

import (
	"github.com/D4vecode/personal_transactions/handlers"
	"github.com/gin-gonic/gin"
)
func TransactionRoutes (r *gin.Engine) {
	r.POST("/upload-transactions", handlers.AddTransactions())
	r.POST("/send-email", handlers.SendTransactionsEmail())
}