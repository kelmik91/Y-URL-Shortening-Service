package config

import (
	"flag"
	"os"
)

var (
	Host       string
	BaseURLRes string
)

func ParseFlags() {
	flag.StringVar(&Host, "a", "localhost:8080", "адрес запуска HTTP-сервера")
	flag.StringVar(&BaseURLRes, "b", "http://localhost:8080", "базовый адрес результирующего сокращённого URL")
	flag.Parse()

	if envRunAddr := os.Getenv("SERVER_ADDRESS"); envRunAddr != "" {
		Host = envRunAddr
	}
	if envBaseAddr := os.Getenv("BASE_URL"); envBaseAddr != "" {
		BaseURLRes = envBaseAddr
	}
}
