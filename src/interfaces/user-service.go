package interfaces

import (
	"todoproject-be/src/models"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	AddUser(ctx *gin.Context)
	GetUser(ctx *gin.Context) *models.User
	ResetPassword(ctx *gin.Context)
}
