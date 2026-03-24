package utils

import (
	"context"
	"fmt"
	"separation/task/dtos"
	taskpb "separation/task/proto/gen"

	"github.com/joho/godotenv"
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
		return "", fmt.Errorf("Missing metadata: %v", ok)
	}

	var tokens []string = md[metaType]
	if len(tokens) == 0 {
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
