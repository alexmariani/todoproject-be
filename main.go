package main

import (
	"todoproject-be/src/database"
	"todoproject-be/src/middlewares"
	"todoproject-be/src/migration"
	"todoproject-be/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	database.ConnectDatabase()
	migration.DoMigration()
	server.Use(middlewares.Logger())
	routes.AddRoutes(server)
	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}
