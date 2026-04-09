package main

import (
	"net"
	"os"
	"task/db"
	"task/handlers"
	"task/logger"
	mid "task/middleware"
	taskpb "task/proto/gen"
	"task/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	// init logger
	if err := logger.Init(true); err != nil {
		panic(err)
	}
	defer logger.Sync()

	err := utils.LoadEnv()
	if err != nil {
		logger.Log.Fatal("load env", zap.Error(err))
	}

	err = db.ConnectDB("DATABASE_URL")
	if err != nil {
		logger.Log.Fatal("connect db", zap.Error(err))
	}
	defer db.CloseDB()

	// Создаем новый gRPC сервер
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(mid.ValidateMiddleware))

	// Регистрация сервиса
	taskpb.RegisterTaskServiceServer(grpcServer, &handlers.Server{})

	// Создаем сетевой слушатель
	var port string = os.Getenv("TASK_PORT")
	if port == "" {
		logger.Log.Fatal("TASK_PORT environment variable is not set")
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Log.Fatal("failed to listen", zap.Error(err))
	}

	logger.Log.Info("Task gRPC server is running", zap.String("port", port))
	if err := grpcServer.Serve(lis); err != nil {
		logger.Log.Fatal("failed to serve", zap.Error(err))
	}
}
