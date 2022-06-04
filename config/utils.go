package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func getString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	if value == "" {
		panic(fmt.Sprintf("%s config is not set", key))
	}

	return value
}

func getInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err == nil {
		return value
	}

	value, err = strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("%s invalid value set", key))

	}
	return value
}
