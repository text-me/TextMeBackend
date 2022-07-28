package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	Title    string
	Messages []Message
}

type Message struct {
	gorm.Model
	Text    string
	GroupID uint
	Group   Group
}
