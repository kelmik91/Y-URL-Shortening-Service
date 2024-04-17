package main

import (
	"github.com/kelmik91/Y-URL-Shortening-Service/internal/config"
	"github.com/kelmik91/Y-URL-Shortening-Service/internal/server"
)

func main() {
	config.ParseFlags()

	server.Run(config.Host)
}
