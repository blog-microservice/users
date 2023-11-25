package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/saufiroja/blog-microservice/users/infrastructure/grpc/pb"
)

type repository struct {
}

func NewRepository() IUserRepository {
	return &repository{}
}
func (r *repository) InsertUser(trx *sql.Tx, req *pb.CreateUserRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	query := `INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := trx.ExecContext(ctx, query, req.Id, req.Name, req.Email, req.Password, req.CreatedAt, req.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}
