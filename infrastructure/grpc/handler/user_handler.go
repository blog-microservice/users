package handler

import (
	"context"

	"github.com/saufiroja/blog-microservice/users/infrastructure/grpc/pb"
	"github.com/saufiroja/blog-microservice/users/infrastructure/producer"
	"github.com/saufiroja/blog-microservice/users/services"
	"github.com/saufiroja/blog-microservice/users/utility"
)

type UserHandler struct {
	pb.UnimplementedUsersServer
	userService services.IUserService
	publish     producer.IProducer
}

func NewUserHandler(userService services.IUserService, publish producer.IProducer) *UserHandler {
	return &UserHandler{
		userService: userService,
		publish:     publish,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.Response, error) {
	err := h.userService.InsertUser(req)
	if err != nil {
		return nil, err
	}

	err = h.publish.Publish(utility.SendEmail, []byte("send email to user"))
	if err != nil {
		return nil, err
	}

	return &pb.Response{
		Message: "success",
	}, nil
}
