package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddlware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		log.Printf("[%s] %s %s %v", ctx.ClientIP(), ctx.Request.Method, ctx.FullPath(), time.Since(start))
	}
}
