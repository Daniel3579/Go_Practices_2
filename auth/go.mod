module auth

go 1.25.0

require (
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.12.3
	golang.org/x/crypto v0.49.0
	google.golang.org/grpc v1.80.0
	google.golang.org/protobuf v1.36.11
)

require go.uber.org/multierr v1.10.0 // indirect

require (
	go.uber.org/zap v1.27.1
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260120221211-b8f7ae30c516 // indirect
)
