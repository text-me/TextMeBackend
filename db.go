package main

import (
	"fmt"
	"github.com/text-me/TextMeBackend/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

type Message struct {
	gorm.Model
	Text string
}

func initDb() {
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

	err := db.AutoMigrate(&Message{})
	if err != nil {
		log.Error(err)
	}
}

func addMessage(text string) *Message {
	insert := &Message{Text: text}
	db.Create(insert)
	return insert
}

func getMessages() []Message {
	var messages []Message
	db.Find(&messages)

	return messages
}
