package server

import "github.com/saufiroja/blog-microservice/users/infrastructure/grpc/pb"

func (rpc *GrpcServer) defineRoute(provider gRPCProvider) {
	pb.RegisterUsersServer(rpc.grpcServer, &provider.handlers.user)
}
