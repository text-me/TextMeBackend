package main

import (
	"github.com/text-me/TextMeBackend/log"
	"github.com/text-me/TextMeBackend/models"
	"net/http"
)

func helloWorldRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("{\"Text\": \"Hello, world!\"}"))
	if err != nil {
		log.Error(err)
	}
}

func getMessagesRoute(w http.ResponseWriter, r *http.Request) {
	messages := models.SelectMessages()

	messagesJson, err := models.MessagesToJson(messages)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application-json")

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(messagesJson)
	if err != nil {
		log.Error(err)
	}
}
