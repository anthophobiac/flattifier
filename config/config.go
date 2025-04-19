package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("No config file found: %v", err)
	} else {
		log.Println("Config loaded from", viper.ConfigFileUsed())
	}
}
