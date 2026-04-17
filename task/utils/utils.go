package utils

import (
	"context"
	"fmt"
	"task/dtos"
	"task/logger"

	taskpb "github.com/Daniel3579/Go_Practices_2/task-sdk/gen"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Ошибка загрузки файла .env: %w", err)
	}
	return nil
}

func GetTokenMetadata(ctx context.Context, metaType string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Log.Warn("missing metadata in context")
		return "", fmt.Errorf("Missing metadata: %v", ok)
	}

	var tokens []string = md[metaType]
	if len(tokens) == 0 {
		logger.Log.Warn("missing token", zap.String("type", metaType))
		return "", fmt.Errorf("Missing %s token", metaType)
	}

	return tokens[0], nil
}

func SliceResponseToRepeatedResponse(res *[]dtos.SelectResponse) (*taskpb.SelectAllResponse, error) {
	selectAllResponse := &taskpb.SelectAllResponse{
		Responses: make([]*taskpb.SelectResponse, len(*res)),
	}

	for i, r := range *res {
		selectAllResponse.Responses[i] = &taskpb.SelectResponse{
			Id:          int32(r.Id),
			Username:    r.Username,
			Title:       r.Title,
			Description: r.Description,
			DueDate:     timestamppb.New(r.Due_date),
			Done:        r.Done,
		}
	}

	return selectAllResponse, nil
}

func Ptr[T any](value T) *T {
	return &value
}
