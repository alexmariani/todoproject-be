package controllers

import (
	"net/http"
	"todoproject-be/src/services"

	"github.com/gin-gonic/gin"
)

type TipController struct {
	TipService *services.TipService
}

func (t TipController) AddTip(ctx *gin.Context) {
	t.TipService.AddTip(ctx)
	ctx.JSON(http.StatusOK, "Tip creato con successo.")
}

func (t TipController) GetAllTips(ctx *gin.Context) {
	tips := t.TipService.GetAllTips(ctx)
	ctx.JSON(http.StatusOK, tips)
}

func (t TipController) GetTip(ctx *gin.Context) {
	tip := t.TipService.GetTip(ctx)
	ctx.JSON(http.StatusOK, tip)
}

func (t TipController) UpdateTip(ctx *gin.Context) {
	tip := t.TipService.UpdateTip(ctx)
	ctx.JSON(http.StatusOK, tip)
}
