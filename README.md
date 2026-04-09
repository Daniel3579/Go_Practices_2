# Коляда Даниил
## Практическая работа №3

### Цель работы

Освоить организацию структурированного логирования в backend-приложении на Go с использованием библиотеки zap для записи, анализа и сопровождения событий приложения

---

### Логи auth сервиса

```cmd
Daniel@Mac auth % go run cmd/main.go
{"level":"info","ts":1775755183.228881,"caller":"cmd/main.go:33","msg":"environment variables loaded successfully"}
{"level":"info","ts":1775755183.2578518,"caller":"db/db.go:44","msg":"Successfully connected to database"}
{"level":"info","ts":1775755183.25791,"caller":"cmd/main.go:43","msg":"Database connection established"}
{"level":"info","ts":1775755183.257993,"caller":"cmd/main.go:63","msg":"gRPC server started","port":":8080"}
{"level":"warn","ts":1775755211.388063,"caller":"utils/token.go:50","msg":"token parsing failed","token_type":"access","error":"token has invalid claims: token is expired"}
{"level":"warn","ts":1775755211.388128,"caller":"handlers/handlers.go:71","msg":"invalid access token","error":"invalid token: token has invalid claims: token is expired"}
{"level":"error","ts":1775755211.388144,"caller":"middleware/middleware.go:26","msg":"RPC call failed","method":"/auth.AuthService/Validate","duration":0.000194625,"error":"rpc error: code = Unauthenticated desc = invalid token","stacktrace":"main.main.UnaryLoggingInterceptor.func3\n\t/Users/Daniel/МИРЭЯ/MAGA/Go2/auth/middleware/middleware.go:26\nauth/proto/gen._AuthService_Validate_Handler\n\t/Users/Daniel/МИРЭЯ/MAGA/Go2/auth/proto/gen/auth_service_grpc.pb.go:187\ngoogle.golang.org/grpc.(*Server).processUnaryRPC\n\t/Users/Daniel/go/pkg/mod/google.golang.org/grpc@v1.80.0/server.go:1430\ngoogle.golang.org/grpc.(*Server).handleStream\n\t/Users/Daniel/go/pkg/mod/google.golang.org/grpc@v1.80.0/server.go:1856\ngoogle.golang.org/grpc.(*Server).serveStreams.func2.1\n\t/Users/Daniel/go/pkg/mod/google.golang.org/grpc@v1.80.0/server.go:1065"}
{"level":"warn","ts":1775755211.4472,"caller":"utils/token.go:50","msg":"token parsing failed","token_type":"refresh","error":"token has invalid claims: token is expired"}
{"level":"warn","ts":1775755211.44733,"caller":"handlers/handlers.go:89","msg":"invalid refresh token","error":"invalid token: token has invalid claims: token is expired"}
{"level":"error","ts":1775755211.4473479,"caller":"middleware/middleware.go:26","msg":"RPC call failed","method":"/auth.AuthService/RefreshToken","duration":0.000256958,"error":"rpc error: code = Unauthenticated desc = invalid token","stacktrace":"main.main.UnaryLoggingInterceptor.func3\n\t/Users/Daniel/МИРЭЯ/MAGA/Go2/auth/middleware/middleware.go:26\nauth/proto/gen._AuthService_RefreshToken_Handler\n\t/Users/Daniel/МИРЭЯ/MAGA/Go2/auth/proto/gen/auth_service_grpc.pb.go:205\ngoogle.golang.org/grpc.(*Server).processUnaryRPC\n\t/Users/Daniel/go/pkg/mod/google.golang.org/grpc@v1.80.0/server.go:1430\ngoogle.golang.org/grpc.(*Server).handleStream\n\t/Users/Daniel/go/pkg/mod/google.golang.org/grpc@v1.80.0/server.go:1856\ngoogle.golang.org/grpc.(*Server).serveStreams.func2.1\n\t/Users/Daniel/go/pkg/mod/google.golang.org/grpc@v1.80.0/server.go:1065"}
{"level":"info","ts":1775755246.90044,"caller":"handlers/handlers.go:141","msg":"user logged in successfully","username":"Daniel"}
{"level":"info","ts":1775755246.900476,"caller":"middleware/middleware.go:32","msg":"RPC call completed","method":"/auth.AuthService/Login","duration":0.073354542}
{"level":"info","ts":1775755255.053501,"caller":"middleware/middleware.go:32","msg":"RPC call completed","method":"/auth.AuthService/Validate","duration":0.000094417}
{"level":"info","ts":1775755261.291154,"caller":"middleware/middleware.go:32","msg":"RPC call completed","method":"/auth.AuthService/Validate","duration":0.000074667}
{"level":"info","ts":1775755267.449643,"caller":"middleware/middleware.go:32","msg":"RPC call completed","method":"/auth.AuthService/Validate","duration":0.000089167}
{"level":"info","ts":1775755270.3899379,"caller":"middleware/middleware.go:32","msg":"RPC call completed","method":"/auth.AuthService/Validate","duration":0.000069375}
{"level":"info","ts":1775755272.810761,"caller":"middleware/middleware.go:32","msg":"RPC call completed","method":"/auth.AuthService/Validate","duration":0.000090459}
{"level":"info","ts":1775755276.3137128,"caller":"middleware/middleware.go:32","msg":"RPC call completed","method":"/auth.AuthService/Validate","duration":0.000069958}
{"level":"info","ts":1775755286.582499,"caller":"middleware/middleware.go:32","msg":"RPC call completed","method":"/auth.AuthService/Validate","duration":0.000062459}
{"level":"info","ts":1775755293.680792,"caller":"middleware/middleware.go:32","msg":"RPC call completed","method":"/auth.AuthService/Validate","duration":0.000063708}
{"level":"info","ts":1775755302.1046379,"caller":"middleware/middleware.go:32","msg":"RPC call completed","method":"/auth.AuthService/Validate","duration":0.000072792}
{"level":"info","ts":1775755310.0684102,"caller":"middleware/middleware.go:32","msg":"RPC call completed","method":"/auth.AuthService/Validate","duration":0.000093625}
{"level":"info","ts":1775755314.924938,"caller":"middleware/middleware.go:32","msg":"RPC call completed","method":"/auth.AuthService/Validate","duration":0.00003225}
^C{"level":"info","ts":1775755335.632031,"caller":"cmd/main.go:71","msg":"Shutting down gRPC server"}
```

### Логи task сервиса
```cmd
Daniel@Mac task % go run cmd/main.go
2026-04-09T20:19:53.165+0300    INFO    db/db.go:38     db connected
2026-04-09T20:19:53.165+0300    INFO    cmd/main.go:52  Task gRPC server is running     {"port": ":8081"}
2026-04-09T20:20:11.324+0300    DEBUG   middleware/middleware.go:21     incoming request        {"method": "/task.TaskService/Insert"}
2026-04-09T20:20:11.388+0300    WARN    auth/auth.go:46 validate failed {"error": "rpc error: code = Unauthenticated desc = invalid token"}
2026-04-09T20:20:11.388+0300    INFO    middleware/middleware.go:32     access token invalid, attempting refresh        {"method": "/task.TaskService/Insert"}
2026-04-09T20:20:11.447+0300    ERROR   middleware/middleware.go:41     refresh token failed    {"error": "rpc error: code = Unauthenticated desc = invalid token"}
task/middleware.ValidateMiddleware
        /Users/Daniel/МИРЭЯ/MAGA/Go2/task/middleware/middleware.go:41
task/proto/gen._TaskService_Insert_Handler
        /Users/Daniel/МИРЭЯ/MAGA/Go2/task/proto/gen/task_service_grpc.pb.go:169
google.golang.org/grpc.(*Server).processUnaryRPC
        /Users/Daniel/go/pkg/mod/google.golang.org/grpc@v1.80.0/server.go:1430
google.golang.org/grpc.(*Server).handleStream
        /Users/Daniel/go/pkg/mod/google.golang.org/grpc@v1.80.0/server.go:1856
google.golang.org/grpc.(*Server).serveStreams.func2.1
        /Users/Daniel/go/pkg/mod/google.golang.org/grpc@v1.80.0/server.go:1065
2026-04-09T20:20:54.993+0300    DEBUG   middleware/middleware.go:21     incoming request        {"method": "/task.TaskService/Insert"}
2026-04-09T20:20:55.054+0300    DEBUG   auth/auth.go:50 validate success        {"username": "Daniel"}
2026-04-09T20:20:55.054+0300    DEBUG   middleware/middleware.go:59     authenticated request   {"username": "Daniel", "method": "/task.TaskService/Insert"}
2026-04-09T20:20:55.065+0300    DEBUG   db/db.go:63     task inserted   {"id": 57, "username": "Daniel"}
2026-04-09T20:20:55.065+0300    INFO    handlers/handlers.go:46 task created    {"id": 57, "username": "Daniel"}
2026-04-09T20:21:01.223+0300    DEBUG   middleware/middleware.go:21     incoming request        {"method": "/task.TaskService/Select"}
2026-04-09T20:21:01.291+0300    DEBUG   auth/auth.go:50 validate success        {"username": "Daniel"}
2026-04-09T20:21:01.291+0300    DEBUG   middleware/middleware.go:59     authenticated request   {"username": "Daniel", "method": "/task.TaskService/Select"}
2026-04-09T20:21:07.365+0300    DEBUG   middleware/middleware.go:21     incoming request        {"method": "/task.TaskService/Select"}
2026-04-09T20:21:07.449+0300    DEBUG   auth/auth.go:50 validate success        {"username": "Daniel"}
2026-04-09T20:21:07.450+0300    DEBUG   middleware/middleware.go:59     authenticated request   {"username": "Daniel", "method": "/task.TaskService/Select"}
2026-04-09T20:21:10.331+0300    DEBUG   middleware/middleware.go:21     incoming request        {"method": "/task.TaskService/Select"}
2026-04-09T20:21:10.390+0300    DEBUG   auth/auth.go:50 validate success        {"username": "Daniel"}
2026-04-09T20:21:10.390+0300    DEBUG   middleware/middleware.go:59     authenticated request   {"username": "Daniel", "method": "/task.TaskService/Select"}
2026-04-09T20:21:12.698+0300    DEBUG   middleware/middleware.go:21     incoming request        {"method": "/task.TaskService/Select"}
2026-04-09T20:21:12.811+0300    DEBUG   auth/auth.go:50 validate success        {"username": "Daniel"}
2026-04-09T20:21:12.811+0300    DEBUG   middleware/middleware.go:59     authenticated request   {"username": "Daniel", "method": "/task.TaskService/Select"}
2026-04-09T20:21:16.252+0300    DEBUG   middleware/middleware.go:21     incoming request        {"method": "/task.TaskService/Select"}
2026-04-09T20:21:16.313+0300    DEBUG   auth/auth.go:50 validate success        {"username": "Daniel"}
2026-04-09T20:21:16.314+0300    DEBUG   middleware/middleware.go:59     authenticated request   {"username": "Daniel", "method": "/task.TaskService/Select"}
2026-04-09T20:21:26.525+0300    DEBUG   middleware/middleware.go:21     incoming request        {"method": "/task.TaskService/SelectAll"}
2026-04-09T20:21:26.583+0300    DEBUG   auth/auth.go:50 validate success        {"username": "Daniel"}
2026-04-09T20:21:26.583+0300    DEBUG   middleware/middleware.go:59     authenticated request   {"username": "Daniel", "method": "/task.TaskService/SelectAll"}
2026-04-09T20:21:33.625+0300    DEBUG   middleware/middleware.go:21     incoming request        {"method": "/task.TaskService/Update"}
2026-04-09T20:21:33.681+0300    DEBUG   auth/auth.go:50 validate success        {"username": "Daniel"}
2026-04-09T20:21:33.681+0300    DEBUG   middleware/middleware.go:59     authenticated request   {"username": "Daniel", "method": "/task.TaskService/Update"}
2026-04-09T20:21:42.044+0300    DEBUG   middleware/middleware.go:21     incoming request        {"method": "/task.TaskService/Update"}
2026-04-09T20:21:42.104+0300    DEBUG   auth/auth.go:50 validate success        {"username": "Daniel"}
2026-04-09T20:21:42.105+0300    DEBUG   middleware/middleware.go:59     authenticated request   {"username": "Daniel", "method": "/task.TaskService/Update"}
2026-04-09T20:21:50.009+0300    DEBUG   middleware/middleware.go:21     incoming request        {"method": "/task.TaskService/Delete"}
2026-04-09T20:21:50.068+0300    DEBUG   auth/auth.go:50 validate success        {"username": "Daniel"}
2026-04-09T20:21:50.069+0300    DEBUG   middleware/middleware.go:59     authenticated request   {"username": "Daniel", "method": "/task.TaskService/Delete"}
2026-04-09T20:21:54.864+0300    DEBUG   middleware/middleware.go:21     incoming request        {"method": "/task.TaskService/Delete"}
2026-04-09T20:21:54.925+0300    DEBUG   auth/auth.go:50 validate success        {"username": "Daniel"}
2026-04-09T20:21:54.925+0300    DEBUG   middleware/middleware.go:59     authenticated request   {"username": "Daniel", "method": "/task.TaskService/Delete"}
^Csignal: interrupt
```

---

### Выводы

Освоили организацию структурированного логирования в backend-приложении на Go с использованием библиотеки zap для записи, анализа и сопровождения событий приложения

---

### Дерево проекта
```
├── README.md
├── auth
│   ├── cmd
│   │   └── main.go
│   ├── db
│   │   └── db.go
│   ├── go.mod
│   ├── go.sum
│   ├── handlers
│   │   └── handlers.go
│   ├── middleware
│   │   └── middleware.go
│   ├── proto
│   │   ├── auth_requests.proto
│   │   ├── auth_responses.proto
│   │   ├── auth_service.proto
│   │   └── gen
│   │       ├── auth_requests.pb.go
│   │       ├── auth_responses.pb.go
│   │       ├── auth_service.pb.go
│   │       └── auth_service_grpc.pb.go
│   └── utils
│       ├── env.go
│       ├── password.go
│       └── token.go
├── screenshots
│   ├── ...
└── task
    ├── auth
    │   └── auth.go
    ├── cmd
    │   └── main.go
    ├── db
    │   └── db.go
    ├── dtos
    │   ├── requests.go
    │   └── responses.go
    ├── go.mod
    ├── go.sum
    ├── handlers
    │   └── handlers.go
    ├── logger
    │   └── logger.go
    ├── middleware
    │   └── middleware.go
    ├── proto
    │   ├── gen
    │   │   ├── task_requests.pb.go
    │   │   ├── task_responses.pb.go
    │   │   ├── task_service.pb.go
    │   │   └── task_service_grpc.pb.go
    │   ├── task_requests.proto
    │   ├── task_responses.proto
    │   └── task_service.proto
    └── utils
        └── utils.go

21 directories, 45 files
```