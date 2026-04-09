package middleware

import (
	"context"
	"task/auth"
	"task/logger"
	"task/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateMiddleware(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	logger.Log.Debug("incoming request", zap.String("method", info.FullMethod))

	token, err := utils.GetTokenMetadata(ctx, "authorization_access")
	if err != nil {
		logger.Log.Warn("missing access token", zap.Error(err), zap.String("method", info.FullMethod))
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	username, err, errCode := auth.RequestValidate(token)
	if err != nil {
		if errCode == codes.Unauthenticated {
			logger.Log.Info("access token invalid, attempting refresh", zap.String("method", info.FullMethod))
			refreshToken, err := utils.GetTokenMetadata(ctx, "authorization_refresh")
			if err != nil {
				logger.Log.Warn("missing refresh token", zap.Error(err))
				return nil, status.Error(codes.Unauthenticated, err.Error())
			}

			accessToken, err, errCode := auth.RequestRefreshToken(refreshToken)
			if err != nil {
				logger.Log.Error("refresh token failed", zap.Error(err))
				return nil, status.Error(errCode, err.Error())
			}

			username, err, errCode = auth.RequestValidate(accessToken)
			if err != nil {
				logger.Log.Error("validate after refresh failed", zap.Error(err))
				return nil, status.Error(errCode, err.Error())
			}

		} else {
			logger.Log.Error("validate token error", zap.Error(err))
			return nil, status.Error(errCode, err.Error())
		}
	}

	// Установите имя пользователя в контекст для использования в следующем обработчике
	ctx = context.WithValue(ctx, "username", username)
	logger.Log.Debug("authenticated request", zap.String("username", username), zap.String("method", info.FullMethod))

	// Вызовите следующий обработчик в цепочке
	return handler(ctx, req)
}
