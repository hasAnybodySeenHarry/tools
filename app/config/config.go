package config

import (
	"database/sql"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type AppConfig struct {
	DatabaseURL string
}

var AppConfigInstance AppConfig
var db *sql.DB

func WatchConfigChanges() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&AppConfigInstance); err != nil {
			panic(err)
		}
		connectToDB()
	})
}

func LoadConfig() {
	viper.SetConfigFile("app/config/database.env")
	viper.ReadInConfig()

	// should panic if file is not detected, in production.

	// if err := viper.ReadInConfig(); err != nil {
	// 	panic(err)
	// }

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" {
		AppConfigInstance.DatabaseURL = databaseURL
		return
	}

	viper.SetDefault("DatabaseURL", "user:password@tcp(localhost:3306)/db_name")
	viper.AutomaticEnv()

	// Unmarshal the configuration into the AppConfigInstance
	if err := viper.Unmarshal(&AppConfigInstance); err != nil {
		panic(err)
	}

	connectToDB()
}

func connectToDB() {
	newDB, err := sql.Open("mysql", AppConfigInstance.DatabaseURL)
	if err != nil {
		return
	}
	if err := newDB.Ping(); err != nil {
		return
	}
	db = newDB
}

func GetDB() *sql.DB {
	return db
}
