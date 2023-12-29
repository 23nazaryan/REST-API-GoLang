package config

import (
	"fmt"
	"gin/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Filed to load .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsm := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsm), &gorm.Config{})
	if err != nil {
		panic("Failed to create connection from DB")
	}

	migrateErr := db.AutoMigrate(&entities.User{}, &entities.Product{}, &entities.Stock{}, &entities.Transaction{})
	if migrateErr != nil {
		panic("Failed to migrate tables to DB")
	}

	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to get connection from DB")
	}

	connectionCloseErr := dbSQL.Close()
	if connectionCloseErr != nil {
		panic("Failed to close connection from DB")
	}
}
