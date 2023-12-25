package routes

import (
	"github.com/gin-gonic/gin"
	"harryd.com/tools/app/interfaces"
)

type RouterInitializer struct {
	MiddlewareInterface interfaces.MiddlewareInterface
}

func (r *RouterInitializer) InitializeRoutes(router *gin.Engine, homeHandler interfaces.HomeHandlerInterface, itemsHandler interfaces.ItemHandlerInterface) {
	apiV1 := router.Group("/api/v1")
	apiV1.GET("/", homeHandler.Home)

	itemsGroup := apiV1.Group("/items")
	r.MiddlewareInterface.SetMiddleware(itemsGroup)
	{
		itemsGroup.GET("/", itemsHandler.GetItems)
		itemsGroup.GET("/:itemID", itemsHandler.GetItem)
		itemsGroup.POST("/", itemsHandler.CreateItem)
		itemsGroup.DELETE("/:itemID", itemsHandler.DeleteItem)
		itemsGroup.PUT("/:itemID", itemsHandler.UpdateItem)
	}
}
