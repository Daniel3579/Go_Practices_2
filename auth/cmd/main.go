package main

import (
	"auth/db"
	"auth/handlers"
	authpb "auth/proto/gen"
	"auth/utils"
	"net"
	"os"
	"os/signal"
	"syscall"

	mid "auth/middleware"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	// Инициализируем логгер
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()

	err = utils.LoadEnv()
	if err != nil {
		sugar.Fatalw("Failed to load environment variables", "error", err)
	}
	sugar.Info("environment variables loaded successfully")

	if err = db.ConnectDB(logger); err != nil {
		sugar.Fatalw("Failed to connect to database", "error", err)
	}
	defer func() {
		if err := db.CloseDB(); err != nil {
			sugar.Errorw("failed to close database connection", "error", err)
		}
	}()
	sugar.Info("Database connection established")

	// Создаем новый gRPC сервер
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(mid.UnaryLoggingInterceptor(logger)),
	)

	// Регистрация сервиса
	authpb.RegisterAuthServiceServer(grpcServer, handlers.NewServer(logger))

	// Создаем сетевой слушатель
	var port string = os.Getenv("AUTH_PORT")
	if port == "" {
		sugar.Fatal("AUTH_PORT environment variable is not set")
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		sugar.Fatalw("Failed to listen on port", "port", port, "error", err)
	}
	sugar.Infow("gRPC server started", "port", port)

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		sugar.Info("Shutting down gRPC server")
		grpcServer.GracefulStop()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		sugar.Fatalw("gRPC server failed", "error", err)
	}
}
