package routes

import (
	"github.com/gin-gonic/gin"
	"harryd.com/tools/app/interfaces"
)

func InitializeRoutes(router *gin.Engine, homeHandler interfaces.HomeHandlerInterface, itemsHandler interfaces.ItemHandlerInterface) {
	router.GET("/", homeHandler.Home)
	router.GET("/items", itemsHandler.GetItems)
	router.GET("/items/:itemID", itemsHandler.GetItem)
}
