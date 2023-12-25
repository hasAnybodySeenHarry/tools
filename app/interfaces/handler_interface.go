package interfaces

import "github.com/gin-gonic/gin"

type HomeHandlerInterface interface {
	Home(c *gin.Context)
}

type ItemHandlerInterface interface {
	GetItems(c *gin.Context)
	GetItem(c *gin.Context)
	CreateItem(c *gin.Context)
	UpdateItem(c *gin.Context)
	DeleteItem(c *gin.Context)
}
