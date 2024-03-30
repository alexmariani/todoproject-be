package utility

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateBody(ctx *gin.Context, body interface{}) {
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		panic(err)
	}
}
