package controllers

import (
	"net/http"
	"todoproject-be/src/models"
	"todoproject-be/src/security"
	"todoproject-be/src/services"
	"todoproject-be/src/utility"

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

func (u UserController) Login(ctx *gin.Context) {
	user := &models.User{}
	utility.ValidateBody(ctx, user)
	userDb := u.UserService.GetUser(ctx)
	err := utility.VerifyPassword(user.Password, userDb.Password)
	utility.CheckErrore(ctx, err, http.StatusInternalServerError)
	token, err := security.GenerateToken(uint(user.ID))
	utility.CheckErrore(ctx, err, http.StatusInternalServerError)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
