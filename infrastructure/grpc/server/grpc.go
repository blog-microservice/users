package server

import (
	"fmt"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/saufiroja/blog-microservice/users/config"
	"github.com/saufiroja/blog-microservice/users/infrastructure/database"
	"github.com/saufiroja/blog-microservice/users/utility"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	host       string
	port       string
	grpcServer *grpc.Server
	conf       *config.AppConfig
}

func NewGrpcServer(host, port string) *GrpcServer {
	return &GrpcServer{
		host: host,
		port: port,
		conf: config.NewAppConfig(),
	}
}

func (s *GrpcServer) Start() {
	logger := utility.InterceptorLogger()
	db := database.NewApplicationDatabase(s.conf, *logger)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.host, s.port))
	if err != nil {
		panic(err)
	}

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	// Register your service here
	provider := s.provide(db)

	s.grpcServer = grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(logger, opts...),
		),

		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(logger, opts...),
		),
	)

	s.defineRoute(provider)

	reflection.Register(s.grpcServer)

	logger.Info("--------------------")
	logger.Info(fmt.Sprintf("grpc server running on %s:%s", s.host, s.port))
	logger.Info("--------------------")

	if err := s.grpcServer.Serve(lis); err != nil {
		logger.Error(fmt.Sprintf("error serving grpc server: %s", err.Error()))
		panic(err)
	}
}
