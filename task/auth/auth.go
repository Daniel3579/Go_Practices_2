package auth

import (
	"context"
	"fmt"
	"os"
	"task/logger"
	"time"

	auth "auth/proto/gen"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func RequestValidate(accessToken string) (string, error, codes.Code) {
	// Подключаемся к gRPC серверу
	var address string = os.Getenv("AUTH_SERVER")
	if address == "" {
		logger.Log.Error("AUTH_SERVER not set")
		return "", fmt.Errorf("AUTH_SERVER environment variable is not set"), codes.InvalidArgument
	}

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Log.Error("grpc dial auth server failed", zap.Error(err), zap.String("addr", address))
		return "", fmt.Errorf("Ошибка подключения к gRPC серверу: %w", err), codes.Internal
	}
	defer conn.Close()

	client := auth.NewAuthServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	md := metadata.Pairs("authorization", accessToken)
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	// Вызов метода Validate без передачи параметров
	resp, err := client.Validate(ctx, &emptypb.Empty{})
	if err != nil {
		logger.Log.Warn("validate failed", zap.Error(err))
		return "", err, codes.Unauthenticated
	}

	logger.Log.Debug("validate success", zap.String("username", resp.GetUsername()))
	return resp.GetUsername(), nil, codes.OK
}

func RequestRefreshToken(refreshToken string) (string, error, codes.Code) {
	var address string = os.Getenv("AUTH_SERVER")
	if address == "" {
		return "", fmt.Errorf("AUTH_SERVER environment variable is not set"), codes.InvalidArgument
	}

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", fmt.Errorf("Ошибка подключения к gRPC серверу: %w", err), codes.Internal
	}
	defer conn.Close()

	client := auth.NewAuthServiceClient(conn)

	// Создание метаданных с токеном
	md := metadata.New(map[string]string{"Authorization": refreshToken})

	// Добавление метаданных к контексту
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// Вызов метода RefreshToken без передачи параметров
	resp, err := client.RefreshToken(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err, codes.Unauthenticated
	}

	return resp.GetAccessToken(), nil, codes.OK
}
