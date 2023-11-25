package services

import (
	"github.com/oklog/ulid/v2"
	"github.com/saufiroja/blog-microservice/users/infrastructure/database"
	"github.com/saufiroja/blog-microservice/users/infrastructure/grpc/pb"
	"github.com/saufiroja/blog-microservice/users/repositories"
)

type service struct {
	userRepo repositories.IUserRepository
	db       database.IApplicationDatabase
}

func NewService(
	userRepo repositories.IUserRepository,
	db database.IApplicationDatabase,
) IUserService {
	return &service{
		userRepo: userRepo,
		db:       db,
	}
}

func (s *service) InsertUser(req *pb.CreateUserRequest) error {
	trx, err := s.db.StartTransaction()
	if err != nil {
		return err
	}

	req.Id = ulid.MustNew(ulid.Now(), nil).String()

	err = s.userRepo.InsertUser(trx, req)
	if err != nil {
		_ = s.db.RollbackTransaction(trx)
		return err
	}

	return s.db.CommitTransaction(trx)
}
