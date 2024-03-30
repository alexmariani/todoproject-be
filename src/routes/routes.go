package routes

import (
	"net/http"
	"todoproject-be/src/controllers"
	"todoproject-be/src/repository"
	"todoproject-be/src/services"

	"github.com/gin-gonic/gin"
)

func AddRoutes(server *gin.Engine) {

	v1 := server.Group("/v1")
	{

		//Api di test
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		tipController, userController := ControllerIstance()

		// Api per l'utente
		v1.GET("/users/:username", userController.GetUser)
		v1.POST("/users", userController.AddUser)

		//Api per il tip
		v1.POST("/tips", tipController.AddTip)
		v1.GET("/tips/:idUtente", tipController.GetAllTips)
		v1.GET("/tips/dettaglio/:idTip", tipController.GetTip)
		v1.PUT("/tips", tipController.UpdateTip)
	}

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
