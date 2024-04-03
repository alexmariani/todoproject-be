package controllers

import (
	"net/http"
	"todoproject-be/src/services"

	"github.com/gin-gonic/gin"
)

type TipController struct {
	TipService *services.TipService
}

// AddTip godoc
// @Summary      Aggiungi un tip
// @Description  Aggiungi un tip per un utente
// @Tags         Tip
// @Accept       json
// @Produce      json
// @Param        tip body models.Tip true "Tip data"
// @Success       200 {string} string "Tip creato con successo."
// @Failure 	  400 {string} string "Bad request"
// @Failure 	  401 {string} string "Unauthorized"
// @Failure 	  500 {string} string "Internal server error"
// @Router       /authenticated/tips [post]
// @Security Bearer
func (t TipController) AddTip(ctx *gin.Context) {
	t.TipService.AddTip(ctx)
	ctx.JSON(http.StatusOK, "Tip creato con successo.")
}

// GetAllTips godoc
// @Summary      Recupera tips
// @Description  Recupera tutti i tips per un utente
// @Tags         Tip
// @Accept		 json
// @Produce      json
// @Param        idUtente path int true "Id utente"
// @Success       200 {array} models.Tip
// @Failure 	  400 {string} string "Bad request"
// @Failure 	  401 {string} string "Unauthorized"
// @Failure 	  500 {string} string "Internal server error"
// @Router       /authenticated/tips/{idUtente} [get]
// @Security Bearer
func (t TipController) GetAllTips(ctx *gin.Context) {
	tips := t.TipService.GetAllTips(ctx)
	ctx.JSON(http.StatusOK, tips)
}

// GetTip godoc
// @Summary      Recupera tip
// @Description  Recupera un singolo tip
// @Tags         Tip
// @Accept		 json
// @Produce      json
// @Param        idTip path int true "Id tip"
// @Success       200 {object} models.Tip
// @Failure 	  400 {string} string "Bad request"
// @Failure 	  401 {string} string "Unauthorized"
// @Failure 	  500 {string} string "Internal server error"
// @Router       /authenticated/tips/dettaglio/{idTip} [get]
// @Security Bearer
func (t TipController) GetTip(ctx *gin.Context) {
	tip := t.TipService.GetTip(ctx)
	ctx.JSON(http.StatusOK, tip)
}

// GetTip godoc
// @Summary      Aggiorna tip
// @Description  Aggiorna un singolo tip
// @Tags         Tip
// @Accept		 json
// @Produce      json
// @Param        tip body models.Tip true "Tip da aggiornare"
// @Success       200 {object} models.Tip
// @Failure 	  400 {string} string "Bad request"
// @Failure 	  401 {string} string "Unauthorized"
// @Failure 	  500 {string} string "Internal server error"
// @Router       /authenticated/tips [put]
// @Security Bearer
func (t TipController) UpdateTip(ctx *gin.Context) {
	tip := t.TipService.UpdateTip(ctx)
	ctx.JSON(http.StatusOK, tip)
}
