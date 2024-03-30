package controllers

import (
	"net/http"
	"todoproject-be/src/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func (u UserController) AddUser(ctx *gin.Context) {
	u.UserService.AddUser(ctx)
	ctx.JSON(http.StatusOK, "Utente creato con successo.")
}

func (u UserController) GetUser(ctx *gin.Context) {
	user := u.UserService.GetUser(ctx)
	ctx.JSON(http.StatusOK, user)
}
