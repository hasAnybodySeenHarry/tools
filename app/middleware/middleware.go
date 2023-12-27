package middleware

import (
	"github.com/gin-gonic/gin"
)

type MiddlewareImpl struct{}

func (m *MiddlewareImpl) SetMiddleware(group *gin.RouterGroup) {
	group.Use(LoggerMiddleware(), AuthMiddleware())
}
