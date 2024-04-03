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

// AddUser godoc
// @Summary      Registra utente
// @Description  Recupera un singolo utente
// @Tags         User
// @Accept		 json
// @Produce      json
// @Param        user body models.User true "Utente da registrare"
// @Success       200 {string} string "Utente creato con successo"
// @Failure 	  400 {string} string "Bad request"
// @Failure 	  401 {string} string "Unauthorized"
// @Failure 	  500 {string} string "Internal server error"
// @Router       /users [post]
// @Security Bearer
func (u UserController) AddUser(ctx *gin.Context) {
	u.UserService.AddUser(ctx)
	ctx.JSON(http.StatusOK, "Utente creato con successo.")
}

// AddUser godoc
// @Summary      Ottieni utente
// @Description  Ottieni un singolo utente
// @Tags         User
// @Accept		 json
// @Produce      json
// @Param        username path string true "Username utente"
// @Success       200 {object} models.User
// @Failure 	  400 {string} string "Bad request"
// @Failure 	  401 {string} string "Unauthorized"
// @Failure 	  500 {string} string "Internal server error"
// @Router       /authenticated/users/{username} [get]
// @Security Bearer
func (u UserController) GetUser(ctx *gin.Context) {
	user := u.UserService.GetUser(ctx)
	ctx.JSON(http.StatusOK, user)
}

// AddUser godoc
// @Summary      Login utente
// @Description   Login utente
// @Tags         User
// @Accept		 json
// @Produce      json
// @Param        user body models.User true "Utente logjn"
// @Success       200 {object} string "token:fjdfjdfljjdfljfd"
// @Failure 	  400 {string} string "Bad request"
// @Failure 	  401 {string} string "Unauthorized"
// @Failure 	  500 {string} string "Internal server error"
// @Router       /users/login [post]
// @Security Bearer
func (u UserController) Login(ctx *gin.Context) {
	user := &models.User{}
	utility.ValidateBody(ctx, user)
	userDb := u.UserService.GetUser(ctx)
	errPass := utility.VerifyPassword(user.Password, userDb.Password)
	utility.CheckErrore(ctx, errPass, http.StatusInternalServerError)
	token, err := security.GenerateToken(uint(user.ID))
	utility.CheckErrore(ctx, err, http.StatusInternalServerError)
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (u UserController) ResetPassword(ctx *gin.Context) {
	u.UserService.ResetPassword(ctx)
	ctx.JSON(http.StatusOK, "Password cambiata con successo")
}
