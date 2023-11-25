package repositories

import (
	"database/sql"

	"github.com/saufiroja/blog-microservice/users/infrastructure/grpc/pb"
)

type IUserRepository interface {
	InsertUser(trx *sql.Tx, req *pb.CreateUserRequest) error
}
