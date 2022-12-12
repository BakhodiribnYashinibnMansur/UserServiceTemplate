package service_test

import (
	"context"
	"testing"
	"users_service/config"
	"users_service/genproto/user_service"
	"users_service/package/repository"
	"users_service/package/service"
	"users_service/util/logrus_log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var logrus *logrus_log.Logger
var services *service.Service

func TestInitFunc(t *testing.T) {
	logrus = logrus_log.GetLogger()
	cfg := config.Get()

	db, err := repository.NewPostgres(context.Background(), cfg, logrus)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db, logrus)
	services = service.NewService(repos, logrus)
	server := grpc.NewServer()
	reflection.Register(server)
	user_service.RegisterUserServiceServer(server, services)
}

func TestUserSignIn(t *testing.T) {
	t.Parallel()
	type UserSignInStruct struct {
		UserName string
		Password string
	}

	type TestCase struct {
		Title      string
		Args       UserSignInStruct
		WantResult string
		WantError  error
	}

	tests := []TestCase{
		{
			Title: "Correct Return Sign In Service",
			Args: UserSignInStruct{
				UserName: "username",
				Password: "password",
			},
			WantResult: "c14e2b22-d2c7-4466-bbdb-892a0fd81a03",
		},
		{
			Title: "Return Not Found Sign In Service",
			Args: UserSignInStruct{
				UserName: "username2",
				Password: "password2",
			},
			WantError: status.Error(codes.NotFound, service.ErrNotFound.Error()),
		},
	}

	for _, test := range tests {
		t.Run(test.Title, func(t *testing.T) {
			logrus.Info("TEST CASE TITLE :  ", test.Title)
			resp, err := services.UserReaderService.SigInUser(context.Background(), &user_service.SignInUserRequest{Username: test.Args.UserName, Password: test.Args.Password})
			if err != nil {
				grpcErr, okGrpc := status.FromError(err)
				wantErr, okWant := status.FromError(test.WantError)
				if okGrpc && okWant {
					if grpcErr.Code() != wantErr.Code() && grpcErr.Message() != wantErr.Message() {
						t.Errorf("want %s, got %s", test.WantError, err.Error())
					}
				}
			}
			if resp != nil {
				if resp.Id != test.WantResult {
					t.Errorf("UserSignInFunc() = %v, want %v", resp, test.WantResult)
				}
			}
		})
	}
}
