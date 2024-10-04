package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig(fileName string) *viper.Viper {
	config := viper.New()

	config.SetConfigName(fileName)

	config.AddConfigPath(".")
	config.AddConfigPath("$HOME")

	err := config.ReadInConfig()//it reads config file based on name and path set earlier
	if err != nil {
		log.Fatal("Error while parsing configuration file", err)
	}

	return config
}
