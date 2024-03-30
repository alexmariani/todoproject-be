package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		now := time.Now()
		log.Printf("[%s] %s %s %s", now.Format(time.RFC3339), c.Request.Method, c.Request.URL.Path, clientIp)
		c.Next()
	}
}
