package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Text string
}

type Group struct {
	gorm.Model
	Title string
}
