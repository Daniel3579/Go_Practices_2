package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"task/db"
	"task/handlers"
	"task/logger"
	mid "task/middleware"
	"task/utils"

	taskpb "github.com/Daniel3579/Go_Practices_2/task-sdk/gen"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// init logger
	if err := logger.Init(true); err != nil {
		panic(err)
	}
	defer logger.Sync()

	// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

	err := utils.LoadEnv()
	if err != nil {
		logger.Log.Fatal("load env", zap.Error(err))
	}
	logger.Log.Info("environment variables loaded successfully")

	// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

	err = db.ConnectDB("DATABASE_URL")
	if err != nil {
		logger.Log.Fatal("connect db", zap.Error(err))
	}
	defer db.CloseDB()
	logger.Log.Info("Database connection established")

	// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

	certFile := os.Getenv("TASK_CERT_FILE")
	if certFile == "" {
		logger.Log.Fatal("TASK_CERT_FILE environment variable is not set")
	}

	keyFile := os.Getenv("TASK_KEY_FILE")
	if keyFile == "" {
		logger.Log.Fatal("TASK_KEY_FILE environment variable is not set")
	}

	// Загружаем TLS credentials
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		logger.Log.Fatal("failed to load TLS credentials: %v", zap.Error(err))
	}

	// Создаем новый gRPC сервер
	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(mid.ValidateMiddleware),
	)

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
	logger.Log.Info("gRPC server started", zap.String("port", port))

	// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		logger.Log.Info("Shutting down gRPC server")
		grpcServer.GracefulStop()
	}()

	if err := grpcServer.Serve(lis); err != nil {
		logger.Log.Fatal("gRPC server failed", zap.Error(err))
	}
}
