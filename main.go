package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"harryd.com/tools/app/config"
	"harryd.com/tools/app/handlers"
	"harryd.com/tools/app/middleware"
	"harryd.com/tools/app/routes"
)

var (
	homeHandlerMu  sync.Mutex
	itemsHandlerMu sync.Mutex
)

func main() {
	config.LoadConfig()

	go config.WatchConfigChanges()

	db := config.GetDB()
	defer db.Close()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	initScriptPath := filepath.Join(wd, "scripts/init.sql")

	if err := runInitScript(db, initScriptPath); err != nil {
		log.Fatal(err)
	}

	router := gin.New()

	homeHandler := handlers.NewHomeHandler()
	itemsHandler := handlers.NewItemHandler(db)

	routerInitializer := &routes.RouterInitializer{
		MiddlewareInterface: &middleware.MiddlewareImpl{},
	}

	routerInitializer.InitializeRoutes(router, homeHandler, itemsHandler)

	if err := router.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}

func createItemHandler(db *sql.DB) *handlers.ItemHandlerImpl {
	itemsHandlerMu.Lock()
	defer itemsHandlerMu.Unlock()
	return handlers.NewItemHandler(db)
}

func runInitScript(db *sql.DB, filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	statements := strings.Split(string(content), ";")

	for _, statement := range statements {
		if strings.TrimSpace(statement) == "" {
			continue
		}
		if _, err := db.Exec(statement); err != nil {
			return err
		}
	}

	return nil
}
