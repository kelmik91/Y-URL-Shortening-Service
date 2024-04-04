package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/kelmik91/Y-URL-Shortening-Service/internal/handlers"
	"net/http"
)

func Run(host string) {
	r := chi.NewRouter()

	r.Post("/", handlers.MainHandler)
	r.Get("/{id}", handlers.MainHandler)

	err := http.ListenAndServe(host, r)
	if err != nil {
		panic(err)
	}
}
