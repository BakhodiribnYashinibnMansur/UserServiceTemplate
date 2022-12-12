package service

import (
	"users_service/package/repository"
	"users_service/util/logrus_log"
)

type UserWriter struct {
	repo   repository.UserWriterRepository
	logrus *logrus_log.Logger
}

func NewUserWriterService(repo repository.UserWriterRepository, logrus *logrus_log.Logger) *UserWriter {
	return &UserWriter{repo: repo, logrus: logrus}
}

// func (aw *UserWriter) UpdateOperator(ctx context.Context, in *admin_service.UpdateOperatorRequest) (*emptypb.Empty, error) {
// 	return aw.repo.UpdateOperator(ctx, in)
// }
