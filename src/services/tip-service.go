package services

import (
	"strconv"
	"todoproject-be/src/models"
	"todoproject-be/src/repository"
	"todoproject-be/src/utility"

	"github.com/gin-gonic/gin"
)

type TipService struct {
	TipRepository *repository.TipRepository
}

func (ts TipService) AddTip(ctx *gin.Context) {
	body := &models.Tip{}
	utility.ValidateBody(ctx, body)
	err := ts.TipRepository.AddTip(body)
	utility.CheckErrore(ctx, err)
}

func (ts TipService) UpdateTip(ctx *gin.Context) models.Tip {
	body := &models.Tip{}
	utility.ValidateBody(ctx, body)
	tip, err := ts.TipRepository.UpdateTip(body)
	utility.CheckErrore(ctx, err)
	return tip
}

func (ts TipService) GetAllTips(ctx *gin.Context) []models.Tip {
	idUtente, err := strconv.Atoi(ctx.Param("idUtente"))
	utility.CheckErrore(ctx, err)
	tips, err := ts.TipRepository.GetAllTips(idUtente)
	utility.CheckErrore(ctx, err)
	return tips
}

func (ts TipService) GetTip(ctx *gin.Context) models.Tip {
	idTip, err := strconv.Atoi(ctx.Param("idTip"))
	utility.CheckErrore(ctx, err)
	tip, err := ts.TipRepository.GetTip(idTip)
	utility.CheckErrore(ctx, err)
	return tip
}
