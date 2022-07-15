package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type NewMessagePayload struct {
	Text string
}

func helloWorldRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("{\"Text\": \"Hello, world!\"}"))
	if err != nil {
		fmt.Println(err)
	}
}

func getMessagesRoute(w http.ResponseWriter, r *http.Request) {
	messages := getMessages()
	jsonResponse, err := json.Marshal(messages)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application-json")

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println(err)
	}
}

func newMessageRoute(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var newMessagePayload NewMessagePayload
	err := decoder.Decode(&newMessagePayload)
	if err != nil {
		fmt.Println(err)
		return
	}

	message := addMessage(newMessagePayload.Text)
	jsonResponse, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application-json")

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(jsonResponse)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Can't load .env file")
	}

	initDb()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", helloWorldRoute)
	r.Post("/newMessage", newMessageRoute)
	r.Get("/getMessages", getMessagesRoute)

	fmt.Println("Listening at localhost:80 ")
	err := http.ListenAndServe(":80", r)
	if err != nil {
		fmt.Println(err)
		return
	}
}
