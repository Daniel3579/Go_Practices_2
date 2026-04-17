package handlers

import (
	"context"
	"task/db"
	"task/dtos"
	"task/logger"
	"task/utils"

	taskpb "github.com/Daniel3579/Go_Practices_2/task-sdk/gen"

	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	taskpb.UnimplementedTaskServiceServer
}

func (s *Server) Insert(ctx context.Context, req *taskpb.InsertRequest) (*taskpb.SelectResponse, error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		logger.Log.Warn("username not found in context")
		return nil, status.Error(codes.Unauthenticated, "username not found in context")
	}

	if req.GetTitle() == "" || req.GetDescription() == "" {
		logger.Log.Warn("bad request", zap.String("username", username))
		return nil, status.Error(codes.InvalidArgument, "Bad request")
	}

	res, err := db.InsertIntoTask(username, &dtos.InsertRequest{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Due_date:    req.GetDueDate().AsTime(),
	})
	if err != nil {
		logger.Log.Error("insert task db error", zap.Error(err), zap.String("username", username))
		return nil, status.Error(codes.Internal, err.Error())
	}

	logger.Log.Info("task created", zap.Int32("id", int32(res.Id)), zap.String("username", username))

	return &taskpb.SelectResponse{
		Id:          int32(res.Id),
		Username:    res.Username,
		Title:       res.Title,
		Description: res.Description,
		DueDate:     timestamppb.New(res.Due_date),
		Done:        res.Done,
	}, nil
}

func (s *Server) Select(ctx context.Context, req *taskpb.IdRequest) (*taskpb.SelectResponse, error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "username not found in context")
	}

	res, err := db.SelectCurrentTask(username, int(req.GetId()))
	if err != nil {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return &taskpb.SelectResponse{
		Id:          int32(res.Id),
		Username:    res.Username,
		Title:       res.Title,
		Description: res.Description,
		DueDate:     timestamppb.New(res.Due_date),
		Done:        res.Done,
	}, nil
}

func (s *Server) SelectAll(ctx context.Context, _ *emptypb.Empty) (*taskpb.SelectAllResponse, error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "username not found in context")
	}

	res, err := db.SelectAllTasks(username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	selectAllRes, err := utils.SliceResponseToRepeatedResponse(res)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return selectAllRes, nil
}

func (s *Server) Update(ctx context.Context, req *taskpb.UpdateRequest) (*taskpb.SelectResponse, error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "username not found in context")
	}

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "\"Id\" field id required")
	}

	var dueDate *time.Time
	if req.DueDate != nil {
		t := req.GetDueDate().AsTime()
		dueDate = &t
	}

	res, err := db.UpdateTask(username, int(req.GetId()), &dtos.UpdateTaskRequest{
		Title:       req.Title,
		Description: req.Description,
		Due_date:    dueDate,
		Done:        req.Done,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &taskpb.SelectResponse{
		Id:          int32(res.Id),
		Username:    res.Username,
		Title:       res.Title,
		Description: res.Description,
		DueDate:     timestamppb.New(res.Due_date),
		Done:        res.Done,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *taskpb.IdRequest) (*emptypb.Empty, error) {
	username, ok := ctx.Value("username").(string)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "username not found in context")
	}

	err := db.DeleteTask(username, int(req.GetId()))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return nil, nil
}
