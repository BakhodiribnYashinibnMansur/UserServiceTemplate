package main

import (
	"context"
	"net"
	"users_service/config"
	"users_service/genproto/user_service"
	"users_service/package/repository"
	"users_service/package/service"
	"users_service/util/logrus_log"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logrus := logrus_log.GetLogger()
	var cfg = config.Get()
	// Migations Up
	repository.MigrateUP(logrus)

	db, err := repository.NewPostgres(context.Background(), cfg, logrus)

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	listen, err := net.Listen("tcp", cfg.ServerPort)
	if err != nil {
		logrus.Fatal("error while listening tcp port: ", err)
	}

	repos := repository.NewRepository(db, logrus)
	services := service.NewService(repos, logrus)
	server := grpc.NewServer()
	reflection.Register(server)

	user_service.RegisterUserServiceServer(server, services)
	if err := server.Serve(listen); err != nil {
		logrus.Fatal("error listening: %v", (err))
	}
	logrus.Info("start service in port " + cfg.ServerPort)

}
