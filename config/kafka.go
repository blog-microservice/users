package config

import "os"

func initKafka(conf *AppConfig) {
	host := os.Getenv("KAFKA_HOST")
	port := os.Getenv("KAFKA_PORT")

	conf.MessageBroker.Host = host
	conf.MessageBroker.Port = port
}
