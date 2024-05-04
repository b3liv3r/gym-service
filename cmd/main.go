package main

import (
	"github.com/b3liv3r/gym-service/config"
	"github.com/b3liv3r/gym-service/modules/db"
	server "github.com/b3liv3r/gym-service/modules/gym/grpc"
	"github.com/b3liv3r/gym-service/modules/gym/repository"
	"github.com/b3liv3r/gym-service/modules/gym/service"
	loggerx "github.com/b3liv3r/logger"
	gymv1 "github.com/b3liv3r/protos-for-gym/gen/go/gym"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	appConf := config.MustLoadConfig()

	logger := loggerx.InitLogger(appConf.Name, appConf.Production)

	sqlDB, err := db.NewSqlDB(logger, appConf.Db)
	if err != nil {
		logger.Fatal("failed to connect to db", zap.Error(err))
	}

	repo := repository.NewGymRepositoryDB(sqlDB)
	service := service.NewGymService(repo, logger)
	s := InitRPC(service)
	lis, err := net.Listen("tcp", appConf.GrpcServerPort)
	if err != nil {
		logger.Error("failed to listen:", zap.Error(err))
	}
	logger.Info("grpc server listening at", zap.Stringer("address", lis.Addr()))
	if err = s.Serve(lis); err != nil {
		logger.Fatal("failed to serve:", zap.Error(err))
	}
}

func InitRPC(gservice service.Gymer) *grpc.Server {
	s := grpc.NewServer()
	gymv1.RegisterGymServer(s, server.NewGymRPCServer(gservice))

	return s
}
