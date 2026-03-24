package main

import (
	"log"
	"net"
	"os"
	"separation/auth/db"
	"separation/auth/handlers"
	"separation/auth/utils"

	authpb "separation/auth/proto/gen"

	"google.golang.org/grpc"
)

func main() {
	err := utils.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	err = db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB()

	// Создаем новый gRPC сервер
	grpcServer := grpc.NewServer()

	// Регистрация сервиса
	authpb.RegisterAuthServiceServer(grpcServer, &handlers.Server{})

	// Создаем сетевой слушатель
	var port string = os.Getenv("AUTH_PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Auth gRPC server is running at " + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
