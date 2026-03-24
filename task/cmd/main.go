package main

import (
	"log"
	"net"
	"os"
	"separation/task/db"
	"separation/task/handlers"
	mid "separation/task/middleware"
	taskpb "separation/task/proto/gen"
	"separation/task/utils"

	"google.golang.org/grpc"
)

func main() {
	err := utils.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = db.ConnectDB("DATABASE_URL")
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	// Создаем новый gRPC сервер
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(mid.ValidateMiddleware))

	// Регистрация сервиса
	taskpb.RegisterTaskServiceServer(grpcServer, &handlers.Server{})

	// Создаем сетевой слушатель
	var port string = os.Getenv("TASK_PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Task gRPC server is running at " + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
