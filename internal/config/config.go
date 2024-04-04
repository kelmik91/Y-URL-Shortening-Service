package config

import "flag"

var (
	Host       = flag.String("a", "localhost:8080", "адрес запуска HTTP-сервера")
	BaseUrlRes = flag.String("b", "http://localhost:8080/", "базовый адрес результирующего сокращённого URL")
)

func ParseFlags() {
	flag.Parse()
}
