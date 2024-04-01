package main

import (
	"github.com/kelmik91/Y-URL-Shortening-Service/internal/handlers"
	"net/http"
)

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", handlers.MainHandler)

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		panic(err)
	}
}
