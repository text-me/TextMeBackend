package models

import (
	"fmt"
	"github.com/text-me/TextMeBackend/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func InitDb() {
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=postgres password=%s dbname=postgres", dbHost, dbPassword)

	if _db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		log.Error(err)
		return
	} else {
		db = _db
	}

	log.Info("DB connection is established")

	err := db.AutoMigrate(&Message{}, &Group{})
	if err != nil {
		log.Error(err)
	}
}
