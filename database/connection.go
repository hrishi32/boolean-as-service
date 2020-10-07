package database

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

var (
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string
)

// GetConnection is a function to provide new or existing database connection.
func GetConnection() (*gorm.DB, error) {
	if db == nil {
		connection, err := createNewConnection()
		if err != nil {
			return nil, err
		}

		db = connection
	}

	return db, nil
}

// createNewConnection is a helper function for GetConnection
func createNewConnection() (*gorm.DB, error) {
	docker := os.Getenv("DOCKER")

	var dsn string
	if docker == "yes" {
		dbPort = os.Getenv("DB_PORT")
		dbHost = os.Getenv("DB_HOST")
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASSWORD")
		dbName = os.Getenv("DB_NAME")

	} else {
		dbPort = "8084"
		dbHost = "127.0.0.1"
		dbUser = "root"
		dbPass = "m"
		dbName = "boolean"

	}
	dsn = dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return connection, err
}
