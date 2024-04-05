package config

import (
	"flag"
	"os"
)

var (
	Host       string
	BaseUrlRes string
)

func ParseFlags() {
	flag.StringVar(&Host, "a", "localhost:8080", "адрес запуска HTTP-сервера")
	flag.StringVar(&BaseUrlRes, "b", "http://localhost:8080", "базовый адрес результирующего сокращённого URL")
	flag.Parse()

	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		Host = envRunAddr
	}
	if envBaseAddr := os.Getenv("BASE_URL"); envBaseAddr != "" {
		BaseUrlRes = envBaseAddr
	}
}
