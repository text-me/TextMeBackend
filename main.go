package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"text\": \"Hello, world!\"}"))
	})

	fmt.Println("Listening at localhost:80")
	err := http.ListenAndServe(":80", r)
	if err != nil {
		fmt.Println(err)
		return
	}
}
