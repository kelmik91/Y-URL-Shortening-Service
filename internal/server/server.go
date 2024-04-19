package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/kelmik91/Y-URL-Shortening-Service/internal/handlers"
	"github.com/kelmik91/Y-URL-Shortening-Service/internal/logger"
	"net/http"
)

func Run(host string) {
	r := chi.NewRouter()

	r.Post("/", logger.WithLoggingFunc(handlers.MainHandlerSet))
	r.Get("/{id}", logger.WithLoggingFunc(handlers.MainHandlerGetByID))

	logger.Sugar.Infow(
		"Start server",
		"host", host,
	)

	err := http.ListenAndServe(host, r)
	if err != nil {
		logger.Sugar.Fatalw(err.Error(), "event", "start server")
		panic(err)
	}
}
