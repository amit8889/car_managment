package driver

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/amit8889/car_managemnt/utils/logger"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)
	var err error
	logger.Info("===db connection start======", dsn)
	db, err = sql.Open("mysql", dsn)

	if err != nil {
		logger.Error("Error opening database: %v", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(20)                  // Maximum number of open connections
	db.SetMaxIdleConns(10)                  // Maximum number of idle connections
	db.SetConnMaxLifetime(30 * time.Minute) // Maximum lifetime of a connection
	db.SetConnMaxIdleTime(10 * time.Minute) // Maximum idle time before a connection is closed

	// Test connection
	if err := db.Ping(); err != nil {
		logger.Error("Error pinging database: %v", err)
	}
	logger.Success("Successfully connected to MySQL with a configured connection pool.", "")
}

func GetDB() *sql.DB {
	return db
}

func CloseDBConn() {
	db.Close()
}
