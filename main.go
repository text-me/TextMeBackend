package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type NewMessagePayload struct {
	Text string `json:"text"`
}

type Message struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
}

type DbConnectionStatus struct {
	Ok bool `json:"ok"`
}

var id int

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"text\": \"Hello, world!\"}"))
	})

	r.Get("/checkConnection", func(w http.ResponseWriter, r *http.Request) {
		isOk := checkConnection()

		status := &DbConnectionStatus{
			Ok: isOk,
		}

		response, err := json.Marshal(status)
		if err != nil {
			fmt.Println(err)
			return
		}

		w.Header().Set("Content-Type", "application-json")

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	r.Post("/newMessage", func(writer http.ResponseWriter, request *http.Request) {
		decoder := json.NewDecoder(request.Body)
		var newMessagePayload NewMessagePayload
		err := decoder.Decode(&newMessagePayload)
		if err != nil {
			fmt.Println(err)
			return
		}

		var message = &Message{
			Id:   id + 1,
			Text: newMessagePayload.Text,
		}

		response, err := json.Marshal(message)
		if err != nil {
			fmt.Println(err)
			return
		}

		writer.Header().Set("Content-Type", "application-json")

		writer.WriteHeader(http.StatusCreated)
		writer.Write(response)

		id += 1
	})

	fmt.Println("Listening at localhost:80")
	err := http.ListenAndServe(":80", r)
	if err != nil {
		fmt.Println(err)
		return
	}
}
