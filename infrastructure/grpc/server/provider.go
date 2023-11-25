package server

import (
	"github.com/saufiroja/blog-microservice/users/infrastructure/database"
	"github.com/saufiroja/blog-microservice/users/infrastructure/grpc/handler"
	"github.com/saufiroja/blog-microservice/users/infrastructure/producer"
	"github.com/saufiroja/blog-microservice/users/repositories"
	"github.com/saufiroja/blog-microservice/users/services"
)

type gRPCProvider struct {
	handlers struct {
		user handler.UserHandler
	}
}

func (rpc *GrpcServer) provide(db database.IApplicationDatabase) gRPCProvider {
	provider := gRPCProvider{}

	publish := producer.NewProducer(rpc.conf)
	userRepo := repositories.NewRepository()
	userService := services.NewService(userRepo, db)
	provider.handlers.user = *handler.NewUserHandler(userService, publish)

	return provider
}
