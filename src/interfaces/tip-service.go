package interfaces

import (
	"todoproject-be/src/models"

	"github.com/gin-gonic/gin"
)

type TipService interface {
	AddTip(ctx *gin.Context)
	UpdateTip(ctx *gin.Context) models.Tip
	GetAllTips(ctx *gin.Context) []models.Tip
	GetTip(ctx *gin.Context) models.Tip
}
