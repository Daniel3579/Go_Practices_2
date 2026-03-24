package middleware

import (
	"context"
	"separation/task/auth"
	"separation/task/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// func ValidateMiddleware(UnaryServerInterceptor grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
// 	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

func ValidateMiddleware(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	token, err := utils.GetTokenMetadata(ctx, "authorization_access")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	username, err, errCode := auth.RequestValidate(token)
	if err != nil {
		if errCode == codes.Unauthenticated {
			refreshToken, err := utils.GetTokenMetadata(ctx, "authorization_refresh")
			if err != nil {
				return nil, status.Error(codes.Unauthenticated, err.Error())
			}

			accessToken, err, errCode := auth.RequestRefreshToken(refreshToken)
			if err != nil {
				return nil, status.Error(errCode, err.Error())
			}

			username, err, errCode = auth.RequestValidate(accessToken)
			if err != nil {
				return nil, status.Error(errCode, err.Error())
			}

		} else {
			return nil, status.Error(errCode, err.Error())
		}
	}

	// Установите имя пользователя в контекст для использования в следующем обработчике
	ctx = context.WithValue(ctx, "username", username)

	// Вызовите следующий обработчик в цепочке
	return handler(ctx, req)
}
