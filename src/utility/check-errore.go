package utility

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckErrore(ctx *gin.Context, err error, httpError int) {
	fmt.Println(httpError)
	if err != nil {
		ctx.JSON(httpError, gin.H{"message": err.Error()})
		panic(err)
	}
}
