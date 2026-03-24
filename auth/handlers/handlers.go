package handlers

import (
	"context"
	"separation/auth/db"
	authpb "separation/auth/proto/gen"
	"separation/auth/utils"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	authpb.UnimplementedAuthServiceServer
}

func (s *Server) SignUp(ctx context.Context, req *authpb.AuthRequest) (*authpb.SignUpResponse, error) {
	if req.GetUsername() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}

	hash, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Произошла ошибка при хэшировании пароля: %v", err)
	}

	err = db.InsertIntoAuth(
		&db.InsertRequest{
			Username: req.GetUsername(),
			Hash:     hash,
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authpb.SignUpResponse{
		Username: req.GetUsername(),
		Hash:     hash,
	}, nil
}

func (s *Server) Validate(ctx context.Context, _ *emptypb.Empty) (*authpb.ValidateResponse, error) {
	token, err := utils.GetTokenMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	username, err := utils.IsValid(token, "access")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &authpb.ValidateResponse{
		Username: username,
	}, nil
}

func (s *Server) RefreshToken(ctx context.Context, _ *emptypb.Empty) (*authpb.RefreshResponse, error) {
	refreshToken, err := utils.GetTokenMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	username, err := utils.IsValid(refreshToken, "refresh")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	accessToken, err := utils.GenerateToken(username, "access", time.Minute*15)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authpb.RefreshResponse{
		AccessToken: accessToken,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *authpb.AuthRequest) (*authpb.LoginResponse, error) {
	hash, err := db.SelectHash(req.GetUsername())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	err = utils.CheckPassword(hash, req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	refreshToken, err := utils.GenerateToken(req.GetUsername(), "refresh", time.Hour*24*7)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	accessToken, err := utils.GenerateToken(req.GetUsername(), "access", time.Minute*15)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authpb.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *authpb.DeleteRequest) (*emptypb.Empty, error) {
	token, err := utils.GetTokenMetadata(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	tokenUsername, err := utils.IsValid(token, "access")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if tokenUsername != req.GetUsername() {
		return nil, status.Error(codes.PermissionDenied, "Удалить можно только себя")
	}

	err = db.DeleteFromAuth(req.GetUsername())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
