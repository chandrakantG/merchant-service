package database

import (
	"database/sql"
	"log"
	"merchant-service/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Config config.DatabaseConfig
}

type ConnectionInterface interface {
	GetConnection() *sql.DB
}

func DBConnection() ConnectionInterface {
	return &Connection{Config: config.GetDBConfig()}
}

func (db *Connection) GetConnection() *sql.DB {

	dbConfig := db.Config

	dbConn, errConn := sql.Open("mysql", dbConfig.User+":"+dbConfig.Password+"@tcp("+dbConfig.Host+")/"+dbConfig.DBName+"?parseTime=true")

	if errConn != nil {
		log.Fatal("Failed to connect to db : ", "database.getConnection", errConn, nil)
		return nil
	}
	errPing := dbConn.Ping()
	if errPing != nil {
		log.Fatal("Error while connecting database", "database.getConnection", errPing, nil)
		return nil
	}

	dbConn.SetMaxOpenConns(dbConfig.MaxPoolSize)
	dbConn.SetMaxIdleConns(dbConfig.MaxIdleConnections)
	dbConn.SetConnMaxLifetime(time.Minute * 5)
	return dbConn
}
