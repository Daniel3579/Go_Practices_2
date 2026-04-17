package main

import (
	"auth/db"
	"auth/handlers"
	"auth/utils"
	"net"
	"os"
	"os/signal"
	"syscall"

	authpb "github.com/Daniel3579/Go_Practices_2/auth-sdk/gen"

	mid "auth/middleware"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	// Инициализируем логгер
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// –––––––––––––––––––––––––––––––––––

	sugar := logger.Sugar()

	// ——————————————————————————————————————————————————————————————————————————

	err = utils.LoadEnv()
	if err != nil {
		sugar.Fatalw("Failed to load environment variables", "error", err)
	}
	sugar.Info("environment variables loaded successfully")

	// ——————————————————————————————————————————————————————————————————————————

	if err = db.ConnectDB(logger); err != nil {
		sugar.Fatalw("Failed to connect to database", "error", err)
	}
	defer func() {
		if err := db.CloseDB(); err != nil {
			sugar.Errorw("failed to close database connection", "error", err)
		}
	}()
	sugar.Info("Database connection established")

	// ——————————————————————————————————————————————————————————————————————————

	// Start metrics HTTP server
	metricsPort := os.Getenv("AUTH_METRICS_PORT")
	if metricsPort == "" {
		metricsPort = ":9090"
	}
	mid.StartMetricsServer(metricsPort, logger)

	// ——————————————————————————————————————————————————————————————————————————

	certFile := os.Getenv("AUTH_CERT_FILE")
	if certFile == "" {
		sugar.Fatal("AUTH_CERT_FILE environment variable is not set")
	}

	keyFile := os.Getenv("AUTH_KEY_FILE")
	if keyFile == "" {
		sugar.Fatal("AUTH_KEY_FILE environment variable is not set")
	}

	// Загружаем TLS credentials
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		sugar.Fatalf("failed to load TLS credentials: %v", err)
	}

	// Создаем новый gRPC сервер
	grpcServer := grpc.NewServer(
		grpc.Creds(creds),
		grpc.UnaryInterceptor(mid.UnaryMetricsInterceptor(logger)),
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

	// ——————————————————————————————————————————————————————————————————————————

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
