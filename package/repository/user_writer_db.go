package repository

import (
	"users_service/util/logrus_log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserWriterDB struct {
	db     *pgxpool.Pool
	logrus *logrus_log.Logger
}

func NewUserWriterDB(db *pgxpool.Pool, logrus *logrus_log.Logger) *UserWriterDB {
	return &UserWriterDB{db: db, logrus: logrus}
}
