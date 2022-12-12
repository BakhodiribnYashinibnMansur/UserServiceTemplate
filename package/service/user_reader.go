package service

import (
	"context"
	"users_service/genproto/user_service"
	"users_service/package/repository"
	"users_service/util/logrus_log"
)

type UserReader struct {
	repo   repository.UserReaderRepository
	logrus *logrus_log.Logger
}

func NewUserReaderService(repo repository.UserReaderRepository, logrus *logrus_log.Logger) *UserReader {
	return &UserReader{repo: repo, logrus: logrus}
}

func (ur *UserReader) SigInUser(ctx context.Context, req *user_service.SignInUserRequest) (*user_service.SignInUserResponse, error) {
	id, err := ur.repo.SigInUser(ctx, req)
	if err != nil {
		return nil, ErrorHandler(err, "")
	}
	return id, nil
}
