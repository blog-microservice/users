package config

import (
	"os"
)

func initGrpc(conf *AppConfig) {
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	conf.Grpc.Port = port
	conf.Grpc.Host = host
}
