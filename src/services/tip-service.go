package services

import (
	"net/http"
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
	utility.CheckErrore(ctx, err, http.StatusInternalServerError)
}

func (ts TipService) UpdateTip(ctx *gin.Context) models.Tip {
	body := &models.Tip{}
	utility.ValidateBody(ctx, body)
	tip, err := ts.TipRepository.UpdateTip(body)
	utility.CheckErrore(ctx, err, http.StatusInternalServerError)
	return tip
}

func (ts TipService) GetAllTips(ctx *gin.Context) []models.Tip {
	idUtente, err := strconv.Atoi(ctx.Param("idUtente"))
	utility.CheckErrore(ctx, err, http.StatusInternalServerError)
	tips, err := ts.TipRepository.GetAllTips(idUtente)
	utility.CheckErrore(ctx, err, http.StatusInternalServerError)
	return tips
}

func (ts TipService) GetTip(ctx *gin.Context) models.Tip {
	idTip, err := strconv.Atoi(ctx.Param("idTip"))
	utility.CheckErrore(ctx, err, http.StatusInternalServerError)
	tip, err := ts.TipRepository.GetTip(idTip)
	utility.CheckErrore(ctx, err, http.StatusInternalServerError)
	return tip
}
