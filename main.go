package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/text-me/TextMeBackend/log"
	"github.com/text-me/TextMeBackend/ws"
	"net/http"
)

type WsMessage struct {
	Type string `json:"type"`
}

type NewMessagePayload struct {
	Text string `json:"text"`
}

func helloWorldRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("{\"Text\": \"Hello, world!\"}"))
	if err != nil {
		log.Error(err)
	}
}

func getMessagesRoute(w http.ResponseWriter, r *http.Request) {
	messages := getMessages()
	jsonResponse, err := json.Marshal(messages)
	if err != nil {
		log.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application-json")

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonResponse)
	if err != nil {
		log.Error(err)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Error("Can't load .env file")
	}

	initDb()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", helloWorldRoute)
	r.Get("/getMessages", getMessagesRoute)

	// Run WebSocket listener
	hub := ws.InitHub()
	go hub.Run()
	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r)
	})

	go func() {
		for {
			clientWsMessage := <-hub.ReceiveMessage

			var wsMessage WsMessage
			if err := json.Unmarshal(clientWsMessage.Data, &wsMessage); err != nil {
				log.Error(err)
				return
			}

			switch wsMessage.Type {
			case "newMessage":
				var newMessagePayload NewMessagePayload
				if err := json.Unmarshal(clientWsMessage.Data, &newMessagePayload); err != nil {
					log.Error(err)
					return
				}

				newMessage := addMessage(newMessagePayload.Text)

				response, err := json.Marshal(newMessage)
				if err != nil {
					log.Error(err)
					return
				}

				clientWsMessage.Client.Send <- response
			}
		}
	}()

	log.Info("Listening at localhost:80")
	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Error(err)
		return
	}
}
