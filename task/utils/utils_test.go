package utils

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"

	"task/dtos"
	"task/logger"
)

// --- LoadEnv tests ---
// Для теста создаём временный .env файл и удаляем после.
func TestLoadEnv_Success(t *testing.T) {
	filename := ".env.test"
	_ = os.WriteFile(filename, []byte("KEY=VALUE\n"), 0o644)
	// ensure cleanup
	defer os.Remove(filename)

	// temporarily change working file name used by godotenv.Load by renaming
	// create an env file with expected name, then call LoadEnv; godotenv.Load() by default
	// loads .env in cwd, so we swap names and restore after test.
	_ = os.Rename(filename, ".env")
	defer func() {
		_ = os.Rename(".env", filename)
	}()

	if err := LoadEnv(); err != nil {
		t.Fatalf("LoadEnv returned unexpected error: %v", err)
	}

	// verify env var was loaded
	if got := os.Getenv("KEY"); got != "VALUE" {
		t.Fatalf("expected KEY=VALUE, got %q", got)
	}
}

func TestLoadEnv_FileMissing(t *testing.T) {
	// ensure .env doesn't exist
	_ = os.Remove(".env")
	err := LoadEnv()
	if err == nil {
		t.Fatal("expected error when .env missing, got nil")
	}
}

// --- GetTokenMetadata tests ---
func TestGetTokenMetadata_Success(t *testing.T) {
	const key = "authorization"
	const token = "bearer abc123"

	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs(key, token))
	got, err := GetTokenMetadata(ctx, key)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != token {
		t.Fatalf("expected token %q, got %q", token, got)
	}
}

func TestGetTokenMetadata_MissingMetadata(t *testing.T) {
	// ensure logger.Log is non-nil to avoid nil-pointer in GetTokenMetadata
	prev := logger.Log
	tmp, _ := zap.NewDevelopment()
	logger.Log = tmp.Sugar().Desugar()
	defer func() { logger.Log = prev }()

	ctx := context.Background() // no metadata
	_, err := GetTokenMetadata(ctx, "authorization")
	if err == nil {
		t.Fatal("expected error when metadata missing, got nil")
	}
}

func TestGetTokenMetadata_MissingToken(t *testing.T) {
	// ensure logger.Log is non-nil
	prev := logger.Log
	tmp, _ := zap.NewDevelopment()
	logger.Log = tmp.Sugar().Desugar()
	defer func() { logger.Log = prev }()

	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("other", "x"))
	_, err := GetTokenMetadata(ctx, "authorization")
	if err == nil {
		t.Fatal("expected error when token missing, got nil")
	}
}

// --- SliceResponseToRepeatedResponse tests ---
func TestSliceResponseToRepeatedResponse_EmptySlice(t *testing.T) {
	in := []dtos.SelectResponse{}
	got, err := SliceResponseToRepeatedResponse(&in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil {
		t.Fatal("expected non-nil response")
	}
	if len(got.Responses) != 0 {
		t.Fatalf("expected 0 responses, got %d", len(got.Responses))
	}
}

func TestSliceResponseToRepeatedResponse_Populated(t *testing.T) {
	d1 := dtos.SelectResponse{
		Id:          1,
		Username:    "u1",
		Title:       "t1",
		Description: "d1",
		Due_date:    time.Date(2026, 4, 20, 12, 0, 0, 0, time.UTC),
		Done:        true,
	}
	d2 := dtos.SelectResponse{
		Id:          2,
		Username:    "u2",
		Title:       "t2",
		Description: "d2",
		Due_date:    time.Date(2026, 4, 21, 13, 30, 0, 0, time.UTC),
		Done:        false,
	}
	in := []dtos.SelectResponse{d1, d2}

	got, err := SliceResponseToRepeatedResponse(&in)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got.Responses) != 2 {
		t.Fatalf("expected 2 responses, got %d", len(got.Responses))
	}

	// manual field checks
	if got.Responses[0].Id != int32(d1.Id) ||
		got.Responses[0].Username != d1.Username ||
		got.Responses[0].Title != d1.Title ||
		got.Responses[0].Description != d1.Description ||
		got.Responses[0].Done != d1.Done {
		t.Fatalf("first response fields mismatch: got %+v", got.Responses[0])
	}

	if got.Responses[1].Id != int32(d2.Id) ||
		got.Responses[1].Username != d2.Username ||
		got.Responses[1].Title != d2.Title ||
		got.Responses[1].Description != d2.Description ||
		got.Responses[1].Done != d2.Done {
		t.Fatalf("second response fields mismatch: got %+v", got.Responses[1])
	}

	// check timestamppb equality by converting back to time
	if !got.Responses[0].DueDate.AsTime().Equal(d1.Due_date) {
		t.Fatalf("due date mismatch for first item: got %v want %v", got.Responses[0].DueDate.AsTime(), d1.Due_date)
	}
	if !got.Responses[1].DueDate.AsTime().Equal(d2.Due_date) {
		t.Fatalf("due date mismatch for second item: got %v want %v", got.Responses[1].DueDate.AsTime(), d2.Due_date)
	}
}

// --- Ptr tests ---
func TestPtr_Primitive(t *testing.T) {
	v := 123
	p := Ptr(v)
	if p == nil {
		t.Fatal("Ptr returned nil")
	}
	if *p != v {
		t.Fatalf("expected %d, got %d", v, *p)
	}
}

func TestPtr_Struct(t *testing.T) {
	type S struct {
		A string
		B int
	}
	s := S{A: "x", B: 2}
	ps := Ptr(s)
	if ps == nil {
		t.Fatal("Ptr returned nil for struct")
	}
	if !reflect.DeepEqual(*ps, s) {
		t.Fatalf("expected %+v, got %+v", s, *ps)
	}
}
