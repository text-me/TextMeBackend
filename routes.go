package main

import (
	"github.com/text-me/TextMeBackend/log"
	"github.com/text-me/TextMeBackend/models"
	"net/http"
)

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

func getGroupsRoute(w http.ResponseWriter, r *http.Request) {
	groups := models.SelectGroups()

	messagesJson, err := models.GroupsToJson(groups)
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
