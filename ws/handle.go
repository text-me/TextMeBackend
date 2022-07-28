package ws

import (
	"encoding/json"
	"github.com/text-me/TextMeBackend/log"
	"github.com/text-me/TextMeBackend/models"
)

const (
	MessageSend     = "messageSend"
	MessageReceived = "messageReceived"
	NewGroupSend    = "newGroupSend"
	NewGroupReceive = "newGroupReceived"
)

type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type MessagePayload struct {
	Text string `json:"text"`
}

type NewGroupPayload struct {
	Title string `json:"title"`
}

func ProcessRequest(message *ClientMessage) {
	var wsMessage Message
	if err := json.Unmarshal(message.Data, &wsMessage); err != nil {
		log.Error(err)
		return
	}

	switch wsMessage.Type {
	case MessageSend:
		var newMessagePayload MessagePayload
		if err := json.Unmarshal(wsMessage.Data, &newMessagePayload); err != nil {
			log.Error(err)
			return
		}

		newMessage, err := models.AddMessage(newMessagePayload.Text)
		if err != nil {
			return
		}
		newMessageJson, err := newMessage.ToJson()
		if err != nil {
			return
		}

		responseMessage := &Message{
			Type: MessageReceived,
			Data: newMessageJson,
		}

		response, err := json.Marshal(responseMessage)
		if err != nil {
			log.Error(err)
			return
		}

		message.Client.Send <- response
	case NewGroupSend:
		var newGroupPayload NewGroupPayload
		if err := json.Unmarshal(wsMessage.Data, &newGroupPayload); err != nil {
			log.Error(err)
			return
		}

		newGroup, err := models.AddGroup(newGroupPayload.Title)
		if err != nil {
			return
		}
		newGroupJson, err := newGroup.ToJson()
		if err != nil {
			return
		}

		responseMessage := &Message{
			Type: NewGroupReceive,
			Data: newGroupJson,
		}

		response, err := json.Marshal(responseMessage)
		if err != nil {
			log.Error(err)
			return
		}

		message.Client.Send <- response
	}
}
