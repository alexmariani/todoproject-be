package utility

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckErrore(ctx *gin.Context, err error) {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		panic(err)
	}
}
