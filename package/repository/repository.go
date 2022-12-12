package repository

import (
	"context"
	"users_service/genproto/user_service"
	"users_service/util/logrus_log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	UserReaderRepository
	UserWriterRepository
}

// Reader interface for selecting data
type UserReaderRepository interface {
	SigInUser(ctx context.Context, req *user_service.SignInUserRequest) (*user_service.SignInUserResponse, error)
}

// Writer interface for inserting data
type UserWriterRepository interface {
}

func NewRepository(db *pgxpool.Pool, logrus *logrus_log.Logger) *Repository {
	return &Repository{UserReaderRepository: NewUserReaderDB(db, logrus), UserWriterRepository: NewUserWriterDB(db, logrus)}
}
