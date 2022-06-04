package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func SetupConfig() error {

	var configFile string
	environment := os.Getenv("ENVIRONMENT")
	if strings.ToUpper(environment) == "TEST" {
		configFile = "application-test"
		viper.AutomaticEnv() //need to load os env variable
	}

	if strings.ToUpper(environment) == "APP" {
		configFile = "application"
		viper.AutomaticEnv() //need to load os env variable
	}

	viper.SetConfigName(configFile)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./../")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("viper err:", err.Error())
	}

	return nil
}

func LoadConfig() {
	environment := os.Getenv("ENVIRONMENT")
	initDatabaseConfig(environment)
	initServiceConfig()
}
