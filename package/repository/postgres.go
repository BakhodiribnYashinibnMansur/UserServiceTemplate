package repository

import (
	"context"
	"fmt"
	"users_service/config"
	"users_service/util/logrus_log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgres(ctx context.Context, cfg *config.Config, logrus *logrus_log.Logger) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return pool, nil
}
