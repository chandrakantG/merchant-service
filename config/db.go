package config

import "strings"

type DatabaseConfig struct {
	Host               string
	User               string
	Password           string
	DBName             string
	MaxPoolSize        int
	MaxIdleConnections int
}

var dbConfig DatabaseConfig

func initDatabaseConfig(environment string) {
	switch strings.ToUpper(environment) {
	case "TEST":
		dbConfig = DatabaseConfig{
			Host:               getString("MYSQL_HOST"),
			User:               getString("MYSQL_USER"),
			Password:           getString("MYSQL_PASSWORD"),
			DBName:             getString("MYSQL_DATABASE"),
			MaxPoolSize:        getInt("MYSQL_MAX_POOL_Size"),
			MaxIdleConnections: getInt("MYSQL_MAX_IDLE_CONNECTIONS"),
		}

	case "APP":
		dbConfig = DatabaseConfig{
			Host:               getString("MYSQL_HOST"),
			User:               getString("MYSQL_USER"),
			Password:           getString("MYSQL_PASSWORD"),
			DBName:             getString("MYSQL_DATABASE"),
			MaxPoolSize:        getInt("MYSQL_MAX_POOL_Size"),
			MaxIdleConnections: getInt("MYSQL_MAX_IDLE_CONNECTIONS"),
		}
	}
}

func GetDBConfig() DatabaseConfig {
	return dbConfig
}
