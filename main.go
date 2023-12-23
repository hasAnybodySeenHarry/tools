package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"harryd.com/shop/app/handlers"
	"harryd.com/shop/app/middleware"
	"harryd.com/shop/app/routes"
)

func main() {
	router := gin.Default();
	router.Use(middleware.LoggerMiddlware())

	homeHandler := &handlers.HomeHandler{}
    itemsHandler := handlers.NewItemHandler()

	routes.InitializeRoutes(router, homeHandler, itemsHandler)

	if err := router.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}