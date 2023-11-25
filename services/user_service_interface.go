package services

import "github.com/saufiroja/blog-microservice/users/infrastructure/grpc/pb"

type IUserService interface {
	InsertUser(req *pb.CreateUserRequest) error
}
