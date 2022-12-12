package service

import (
	"context"

	"users_service/genproto/user_service"
	"users_service/util/logrus_log"

	"users_service/package/repository"
)

type Service struct {
	UserReaderService
	UserWriterService
}

// UserReaderService interface for selecting data
type UserReaderService interface {
	SigInUser(ctx context.Context, req *user_service.SignInUserRequest) (*user_service.SignInUserResponse, error)
}

// Writer interface for inserting data
type UserWriterService interface {
}

func NewService(repos *repository.Repository, logrus *logrus_log.Logger) *Service {
	return &Service{UserReaderService: NewUserReaderService(repos.UserReaderRepository, logrus), UserWriterService: NewUserWriterService(repos.UserWriterRepository, logrus)}
}
