package main

import (
	"todoproject-be/docs"
	"todoproject-be/src/database"
	"todoproject-be/src/middlewares"
	"todoproject-be/src/migration"
	"todoproject-be/src/routes"

	"github.com/gin-gonic/gin"
)

// @title           Swagger Example API
// @version         1.0
// @description     Backend todoproject.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	server := gin.Default()
	database.ConnectDatabase()
	migration.DoMigration()
	server.Use(middlewares.Logger())
	routes.AddRoutes(server)
	docs.SwaggerInfo.Title = "Swagger Example APi"

	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}
