package models

import (
	"encoding/json"
	"github.com/text-me/TextMeBackend/log"
)

type MessageJson struct {
	Id   uint   `json:"id"`
	Text string `json:"text"`
}

func AddMessage(text string) *Message {
	insert := &Message{Text: text}
	db.Create(insert)
	return insert
}

func SelectMessages() []Message {
	var messages []Message
	db.Find(&messages)

	return messages
}

func MessagesToJson(messages []Message) ([]byte, error) {
	var list []MessageJson
	for _, msg := range messages {
		list = append(list, MessageJson{
			Id:   msg.ID,
			Text: msg.Text,
		})
	}

	messageRawJson, err := json.Marshal(list)
	if err != nil {
		log.Error(err)
		return make([]byte, 0), err
	}

	return messageRawJson, nil
}

func (m Message) ToJson() ([]byte, error) {
	newMessageJson := &MessageJson{
		Id:   m.ID,
		Text: m.Text,
	}

	messageRawJson, err := json.Marshal(newMessageJson)
	if err != nil {
		log.Error(err)
		return make([]byte, 0), err
	}

	return messageRawJson, nil
}
