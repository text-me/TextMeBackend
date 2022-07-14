package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func checkConnection() bool {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Can't load .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=postgres password=%s dbname=%s", dbHost, dbPassword, dbName)
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Println("DB connection is established")
	return true
}
