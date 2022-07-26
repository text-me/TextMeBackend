package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/text-me/TextMeBackend/log"
	"github.com/text-me/TextMeBackend/models"
	"github.com/text-me/TextMeBackend/ws"
	"net/http"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Error("Can't load .env file")
	}

	models.InitDb()

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
			ws.ProcessRequest(clientWsMessage)
		}
	}()

	log.Info("Listening at localhost:80")
	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Error(err)
		return
	}
}
