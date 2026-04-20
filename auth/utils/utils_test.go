package utils

import (
	"context"
	"os"
	"testing"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

func TestHashAndCheckPassword(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	password := "StrongP@ssw0rd!"
	hash, err := HashPassword(password, logger)
	if err != nil {
		t.Fatalf("HashPassword returned error: %v", err)
	}
	if hash == "" {
		t.Fatal("expected non-empty hash")
	}

	// Correct password should pass
	if err := CheckPassword(hash, password, logger); err != nil {
		t.Fatalf("CheckPassword failed for correct password: %v", err)
	}

	// Incorrect password should fail
	if err := CheckPassword(hash, "wrongpassword", logger); err == nil {
		t.Fatal("CheckPassword succeeded for wrong password; expected failure")
	}
}

func TestGenerateAndValidateToken_Success(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// Ensure SECRET_KEY set for tests
	os.Setenv("SECRET_KEY", "test-secret-key")
	defer os.Unsetenv("SECRET_KEY")

	username := "alice"
	// generate access token
	token, err := GenerateToken(username, "access", time.Minute*5, logger)
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}
	if token == "" {
		t.Fatal("expected non-empty token")
	}

	// validate token
	gotUsername, err := IsValid(token, "access", logger)
	if err != nil {
		t.Fatalf("IsValid returned error for valid token: %v", err)
	}
	if gotUsername != username {
		t.Fatalf("IsValid returned username %q, want %q", gotUsername, username)
	}
}

func TestGenerateAndValidateToken_WrongType(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	os.Setenv("SECRET_KEY", "test-secret-key")
	defer os.Unsetenv("SECRET_KEY")

	username := "bob"
	// generate refresh token
	token, err := GenerateToken(username, "refresh", time.Hour, logger)
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	// Try validate as access token -> should fail
	_, err = IsValid(token, "access", logger)
	if err == nil {
		t.Fatal("IsValid succeeded for token with wrong type; expected error")
	}
}

func TestGenerateAndValidateToken_Expired(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	os.Setenv("SECRET_KEY", "test-secret-key")
	defer os.Unsetenv("SECRET_KEY")

	username := "carol"
	// generate token with negative duration (already expired)
	token, err := GenerateToken(username, "access", -time.Minute, logger)
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	_, err = IsValid(token, "access", logger)
	if err == nil {
		t.Fatal("IsValid succeeded for expired token; expected error")
	}
}

func TestGetTokenMetadata_FromContext(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	ctxWithMeta := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "Bearer abc.xyz"))
	token, err := GetTokenMetadata(ctxWithMeta, logger)
	if err != nil {
		t.Fatalf("GetTokenMetadata returned error: %v", err)
	}
	if token != "Bearer abc.xyz" {
		t.Fatalf("unexpected token: got %q want %q", token, "Bearer abc.xyz")
	}
}

func TestGetTokenMetadata_Missing(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	_, err := GetTokenMetadata(context.Background(), logger)
	if err == nil {
		t.Fatal("GetTokenMetadata succeeded with missing metadata; expected error")
	}
}
