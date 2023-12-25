package config

import (
	"os"

	"github.com/spf13/viper"
)

type AppConfig struct {
	DatabaseURL string
}

var AppConfigInstance AppConfig

func LoadConfig() {
	viper.SetConfigFile("app/config/database.env")
	viper.ReadInConfig()

	// should panic if file is not detected, in production.

	// if err := viper.ReadInConfig(); err != nil {
	// 	panic(err)
	// }

	// fallback for development purpose
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL != "" {
		AppConfigInstance.DatabaseURL = databaseURL
		return
	}

	// default
	viper.SetDefault("DatabaseURL", "user:password@tcp(localhost:3306)/db_name")
	viper.AutomaticEnv()

	// Unmarshal the configuration into the AppConfigInstance
	if err := viper.Unmarshal(&AppConfigInstance); err != nil {
		panic(err)
	}
}
