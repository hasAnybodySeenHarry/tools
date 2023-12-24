package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
    DatabaseURL string
}

var AppConfigInstance AppConfig

func LoadConfig() {
	viper.SetConfigFile("app/.env")
    viper.ReadInConfig()

    viper.SetDefault("DatabaseURL", "user:password@tcp(localhost:3306)/db_name")
    viper.AutomaticEnv()

    // Unmarshal the configuration into the AppConfigInstance
    if err := viper.Unmarshal(&AppConfigInstance); err != nil {
        panic(err)
    }
}
