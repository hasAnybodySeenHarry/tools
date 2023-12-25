package interfaces

import "github.com/gin-gonic/gin"

type MiddlewareInterface interface {
    SetMiddleware(group *gin.RouterGroup)
}