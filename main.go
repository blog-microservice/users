package main

import (
	"github.com/saufiroja/blog-microservice/users/config"
	"github.com/saufiroja/blog-microservice/users/infrastructure/grpc/server"
)

func main() {
	port := config.NewAppConfig().Grpc.Port
	host := config.NewAppConfig().Grpc.Host
	grpcServer := server.NewGrpcServer(host, port)
	grpcServer.Start()
}
