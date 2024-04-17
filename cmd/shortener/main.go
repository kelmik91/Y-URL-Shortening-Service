package main

import "github.com/kelmik91/Y-URL-Shortening-Service/internal/server"

func main() {
	host := "localhost"
	port := "8080"

	server.Run(host, port)
}
