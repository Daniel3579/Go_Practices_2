package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"encoding/json"
	"log"
	"net/http"
	"time"
)

// HealthResponse структура ответа health-эндпоинта
type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	// Можно добавить поле Details map[string]string для проверок зависимостей
}

// healthHandler обрабатывает GET /health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Здесь можно выполнить реальные проверки:
	// - соединение с БД
	// - доступность кеша
	// - место на диске и т.д.
	// Если хотя бы одна проверка не пройдена, вернуть HTTP 503.

	// Пример всегда успешного ответа
	resp := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func health() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	// Дополнительно можно добавить простой эндпоинт для проверки готовности (readiness)
	mux.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		// Логика проверки готовности (например, загружены ли данные)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ready"))
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Запуск сервера в горутине
	go func() {
		log.Printf("Starting server on port %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
	log.Println("Server exited gracefully")
}

func main() {
	health()

	// // init logger
	// if err := logger.Init(true); err != nil {
	// 	panic(err)
	// }
	// defer logger.Sync()

	// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

	// err := utils.LoadEnv()
	// if err != nil {
	// 	logger.Log.Fatal("load env", zap.Error(err))
	// }
	// logger.Log.Info("environment variables loaded successfully")

	// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

	// err = db.ConnectDB("DATABASE_URL")
	// if err != nil {
	// 	logger.Log.Fatal("connect db", zap.Error(err))
	// }
	// defer db.CloseDB()
	// logger.Log.Info("Database connection established")

	// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

	// certFile := os.Getenv("TASK_CERT_FILE")
	// if certFile == "" {
	// 	logger.Log.Fatal("TASK_CERT_FILE environment variable is not set")
	// }

	// keyFile := os.Getenv("TASK_KEY_FILE")
	// if keyFile == "" {
	// 	logger.Log.Fatal("TASK_KEY_FILE environment variable is not set")
	// }

	// Загружаем TLS credentials
	// creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	// if err != nil {
	// 	logger.Log.Fatal("failed to load TLS credentials: %v", zap.Error(err))
	// }

	// // Создаем новый gRPC сервер
	// grpcServer := grpc.NewServer(
	// 	grpc.Creds(creds),
	// 	grpc.UnaryInterceptor(mid.ValidateMiddleware),
	// )

	// // Регистрация сервиса
	// taskpb.RegisterTaskServiceServer(grpcServer, &handlers.Server{})

	// // Создаем сетевой слушатель
	// var port string = os.Getenv("TASK_PORT")
	// if port == "" {
	// 	logger.Log.Fatal("TASK_PORT environment variable is not set")
	// }

	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	logger.Log.Fatal("failed to listen", zap.Error(err))
	// }

	// logger.Log.Info("Task gRPC server is running", zap.String("port", port))
	// if err := grpcServer.Serve(lis); err != nil {
	// 	logger.Log.Fatal("failed to serve", zap.Error(err))
	// }
	// logger.Log.Info("gRPC server started", zap.String("port", port))

	// ––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––––

	// Graceful shutdown
	// go func() {
	// 	sigChan := make(chan os.Signal, 1)
	// 	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// 	<-sigChan

	// 	logger.Log.Info("Shutting down gRPC server")
	// 	grpcServer.GracefulStop()
	// }()

	// if err := grpcServer.Serve(lis); err != nil {
	// 	logger.Log.Fatal("gRPC server failed", zap.Error(err))
	// }
}
