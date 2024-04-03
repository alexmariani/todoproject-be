package routes

import (
	"net/http"
	"todoproject-be/src/controllers"
	"todoproject-be/src/middlewares"
	"todoproject-be/src/repository"
	"todoproject-be/src/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AddRoutes(server *gin.Engine) {

	//Recupero instanze dei controller
	tipController, userController := ControllerIstance()

	server.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	api := server.Group("/api/v1")
	{
		// Api per l'utente
		api.POST("/users", userController.AddUser)
		api.POST("/users/login", userController.Login)
	}

	//Api per gli utenti autenticati
	auth := server.Group("/api/v1/authenticated")
	{
		//JWT middleware
		auth.Use(middlewares.JwtAuthMiddleware())

		// Api per l'utente
		api.GET("/users/:username", userController.GetUser)
		api.POST("/users/reset", userController.ResetPassword)

		//Api per il tip
		auth.GET("/tips/:idUtente", tipController.GetAllTips)
		auth.GET("/tips/dettaglio/:idTip", tipController.GetTip)
		auth.POST("/tips", tipController.AddTip)
		auth.PUT("/tips", tipController.UpdateTip)
	}
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//Api route errata
	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
}

func ControllerIstance() (controllers.TipController, controllers.UserController) {
	tipRepository := &repository.TipRepository{}
	userRepository := &repository.UserRepository{}

	tipService := &services.TipService{TipRepository: tipRepository}
	userService := &services.UserService{UserRepository: userRepository}

	tipController := &controllers.TipController{TipService: tipService}
	userController := &controllers.UserController{UserService: userService}

	return *tipController, *userController
}
