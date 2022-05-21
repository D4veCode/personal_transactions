package routes

import (
	"github.com/D4vecode/personal_transactions/handlers"
	"github.com/gin-gonic/gin"
)


func AccountsRoutes (r *gin.Engine) {
	r.POST("/create-account", handlers.CreateAccount())
}