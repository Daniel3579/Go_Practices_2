# Коляда Даниил
## Практическая работа №5

### Цель работы

Освоить базовые практические подходы к защите backend-приложения на Go за счёт включения HTTPS с использованием TLS-сертификата и предотвращения SQL-инъекций при работе с базой данных

---

### Создание CA сертификата и ключа
```bash
# Создаем приватный ключ CA
openssl genrsa -out certs/ca.key 2048

# Создаем самоподписанный CA сертификат
openssl req -new -x509 -key certs/ca.key -out certs/ca.crt -days 365 \
  -subj "/CN=MyCA"
```

---

### Создание сертификатов для Auth сервиса
```bash
# Создаем приватный ключ для Auth
openssl genrsa -out certs/auth/server.key 2048

# Создаем CSR (Certificate Signing Request)
openssl req -new -key certs/auth/server.key -out certs/auth/server.csr \
  -subj "/CN=auth-service"

# Подписываем CA сертификатом
openssl x509 -req -in certs/auth/server.csr -CA certs/ca.crt -CAkey certs/ca.key \
  -CAcreateserial -out certs/auth/server.crt -days 365 \
  -extfile <(printf "subjectAltName=DNS:auth-service,DNS:localhost,IP:127.0.0.1")

# Удаляем CSR
rm certs/auth/server.csr
```

![Screenshot](./screenshots/Screenshot_3.png)

---

### Создание сертификатов для Task сервиса
```bash
# Создаем приватный ключ для Task
openssl genrsa -out certs/task/server.key 2048

# Создаем CSR
openssl req -new -key certs/task/server.key -out certs/task/server.csr \
  -subj "/CN=task-service"

# Подписываем CA сертификатом
openssl x509 -req -in certs/task/server.csr -CA certs/ca.crt -CAkey certs/ca.key \
  -CAcreateserial -out certs/task/server.crt -days 365 \
  -extfile <(printf "subjectAltName=DNS:task-service,DNS:localhost,IP:127.0.0.1")

# Удаляем CSR
rm certs/task/server.csr
```

---

### Тестирование

Попытка вызова эндпоинта через HTTP

![Screenshot](./screenshots/Screenshot_2.png)

Вызов эндпоинта через HTTPS

![Screenshot](./screenshots/Screenshot_1.png)

---

### Предотвращение SQL-инъекций

Для предотващений SQL-инъекций была использована передача значений через параметризацию

```go
err := db.QueryRow("INSERT INTO task (username, title, description, due_date) VALUES ($1, $2, $3, $4) RETURNING *;",
    username,
    req.Title,
    req.Description,
    req.Due_date,
).Scan(&res.Id, &res.Username, &res.Title, &res.Description, &res.Due_date, &res.Done)
```

---

### Выводы

Освоили базовые практические подходы к защите backend-приложения на Go за счёт включения HTTPS с использованием TLS-сертификата и предотвращения SQL-инъекций при работе с базой данных

---

### Дерево проекта
```
├── .vscode
│   └── launch.json
├── README.md
├── auth
│   ├── .env
│   ├── certs
│   │   ├── ca.crt
│   │   ├── server.crt
│   │   └── server.key
│   ├── cmd
│   │   └── main.go
│   ├── db
│   │   └── db.go
│   ├── docker-compose.yml
│   ├── go.mod
│   ├── go.sum
│   ├── handlers
│   │   └── handlers.go
│   ├── middleware
│   │   └── middleware.go
│   ├── monitoring
│   │   └── prometheus.yml
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
├── ca.key
├── screenshots
│   └── ...
└── task
    ├── .env
    ├── auth
    │   └── auth.go
    ├── certs
    │   ├── ca.crt
    │   ├── server.crt
    │   └── server.key
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
```