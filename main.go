package main

import (
	db "github.com/D4vecode/personal_transactions/database"
	"github.com/D4vecode/personal_transactions/routes"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	_ "github.com/D4vecode/personal_transactions/docs"
)

// @title Personal Transactions API
// @version version(1.0)
// @description Simple transactions API to handle CSV files and Emails.

// @contact.url http://www.swagger.io/support
// @contact.email davidasb.developer@gmail.com
// @license.name license(Mandatory)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {

	
	router := gin.Default()

	db.ConnectDB()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	routes.AccountsRoutes(router)
	routes.TransactionRoutes(router)

	router.Run(":8080")

}