package services

import (
	"todoproject-be/src/models"
	"todoproject-be/src/repository"
	"todoproject-be/src/utility"

	"github.com/gin-gonic/gin"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (us UserService) AddUser(ctx *gin.Context) {
	body := &models.User{}
	utility.ValidateBody(ctx, body)
	err := us.UserRepository.AddUser(body)
	utility.CheckErrore(ctx, err)
}

func (us UserService) GetUser(ctx *gin.Context) *models.User {
	user := &models.User{}
	username := ctx.Param("username")
	user, err := us.UserRepository.GetUser(username)
	utility.CheckErrore(ctx, err)
	return user
}
