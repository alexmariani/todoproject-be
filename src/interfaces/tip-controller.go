package interfaces

import (
	"github.com/gin-gonic/gin"
)

type TipController interface {
	AddTip(ctx *gin.Context)
	GetAllTips(ctx *gin.Context)
	GetTip(ctx *gin.Context)
	UpdateTip(ctx *gin.Context)
}
