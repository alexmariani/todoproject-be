package services

import (
	"net/http"
	"todoproject-be/src/models"
	"todoproject-be/src/repository"
	"todoproject-be/src/utility"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (us UserService) AddUser(ctx *gin.Context) {
	body := &models.User{}
	utility.ValidateBody(ctx, body)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	utility.CheckErrore(ctx, err, http.StatusInternalServerError)
	body.Password = string(hashedPassword)
	errIns := us.UserRepository.AddUser(body)
	utility.CheckErrore(ctx, errIns, http.StatusInternalServerError)
}

func (us UserService) GetUser(ctx *gin.Context) *models.User {
	user := &models.User{}
	username := ctx.Param("username")
	user, err := us.UserRepository.GetUser(username)
	utility.CheckErrore(ctx, err, http.StatusInternalServerError)
	return user
}

func (us UserService) ResetPassword(ctx *gin.Context) {
	body := &models.User{}
	utility.ValidateBody(ctx, body)
	err := us.UserRepository.ResetPassword(body.Username, body.Password)
	utility.CheckErrore(ctx, err, http.StatusInternalServerError)
}
