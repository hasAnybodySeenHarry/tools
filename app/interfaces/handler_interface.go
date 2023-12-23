package interfaces

import "github.com/gin-gonic/gin"

type HomeHandlerInterface interface {
    Home(c *gin.Context)
}

type ItemHandlerInterface interface {
	GetItems(c *gin.Context)
    GetItem(c *gin.Context)
}