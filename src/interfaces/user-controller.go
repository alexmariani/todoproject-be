package interfaces

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	AddUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
}
