package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"todo-list/cmd/app/models"
)

var DB *gorm.DB

// ConnectDB Connect to MySQL database
func ConnectDB() (*gorm.DB, error) {
	mysqlDataBaseName := os.Getenv("MYSQL_DATABASE")
	mysqlUserName := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")

	// Database configuration
	dsn := mysqlUserName + ":" + mysqlPassword + "@tcp(mysql:3306)/" + mysqlDataBaseName
	fmt.Println(dsn)
	// Connect to MySQL database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Task{})

	return db, nil
}
