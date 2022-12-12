package repository

import (
	"context"
	"users_service/genproto/user_service"
	"users_service/util/helper"
	"users_service/util/logrus_log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserReaderDB struct {
	db     *pgxpool.Pool
	logrus *logrus_log.Logger
}

func NewUserReaderDB(db *pgxpool.Pool, logrus *logrus_log.Logger) *UserReaderDB {
	return &UserReaderDB{db: db, logrus: logrus}
}

func (ur *UserReaderDB) SigInUser(ctx context.Context, req *user_service.SignInUserRequest) (*user_service.SignInUserResponse, error) {
	logrus := ur.logrus
	id := &user_service.SignInUserResponse{}

	query := `SELECT
	id
	 FROM account
	 WHERE
	  username=:username
	   AND
	    password=:password`
	params := map[string]interface{}{
		"username": req.Username,
		"password": req.Password,
	}

	query, arr := helper.ReplaceQueryParams(query, params)
	row := ur.db.QueryRow(ctx, query, arr...)

	err := row.Scan(&id.Id)
	if err != nil {
		logrus.Error("!!!Sign In User--->Query  --- ", err)
		return id, err
	}

	return id, nil
}
