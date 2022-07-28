package models

import (
	"errors"
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

	// Migrations
	err := db.AutoMigrate(&Message{}, &Group{})
	log.Info("migration is done")
	if err != nil {
		log.Error(err)
	}

	// Seed data
	if errors.Is(db.First(&Group{}).Error, gorm.ErrRecordNotFound) {
		err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Create(&Group{Title: "Main"}).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			log.Info("Seed data was inserted")
		}
	}
}
