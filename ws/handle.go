package ws

import (
	"encoding/json"
	"github.com/text-me/TextMeBackend/log"
	"github.com/text-me/TextMeBackend/models"
)

const (
	MessageSend     = "messageSend"
	MessageReceived = "messageReceived"
)

type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type MessagePayload struct {
	Text string `json:"text"`
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

		newMessage := models.AddMessage(newMessagePayload.Text)
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
	}
}
