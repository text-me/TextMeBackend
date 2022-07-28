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
	list := make([]MessageJson, 0)
	for _, msg := range messages {
		list = append(list, MessageJson{
			Id:   msg.ID,
			Text: msg.Text,
		})
	}

	messagesJson, err := json.Marshal(list)
	if err != nil {
		log.Error(err)
		return make([]byte, 0), err
	}

	return messagesJson, nil
}

func (m Message) ToJson() ([]byte, error) {
	newMessageJson := &MessageJson{
		Id:   m.ID,
		Text: m.Text,
	}

	messageJson, err := json.Marshal(newMessageJson)
	if err != nil {
		log.Error(err)
		return make([]byte, 0), err
	}

	return messageJson, nil
}
