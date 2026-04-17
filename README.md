# Коляда Даниил
## Практическая работа №7

### Цель работы

Освоить контейнеризацию backend-приложения на Go с помощью Docker, научиться писать Dockerfile, собирать Docker-образ и запускать контейнеризированный сервис в воспроизводимой среде

---

### Команды сборки образов и запуска контейнеров

Команда сборки и запуска для сервиса **Task**

```bash
docker build -t task-image .
docker run --rm -p 8081:443 -v ./certs:/app/certs:ro -v ./.env:/app/.env:ro task-image
```

---

Команда сборки и запуска для сервиса **Auth**

```bash
docker build -t auth-image .
docker run --rm -p 8080:443 -v ./certs:/app/certs:ro -v ./.env:/app/.env:ro auth-image
```

---

Команда сборки и запуска базы данных **Postgres**

```bash
docker build -t postgres-image .
docker run --rm -p 5433:5432 --env-file ./.env -v ./data/postgres:/var/lib/postgresql/data:rw postgres-image
```

---

Команда сборки и запуска для **docker-compose**

```bash
docker compose -p practice up -d --build
```

---

### Результаты

![Screenshot](./screenshots/Screenshot_1.png)
![Screenshot](./screenshots/Screenshot_2.png)

---


### Выводы

Освоили контейнеризацию backend-приложения на Go с помощью Docker, научились писать Dockerfile, собирать Docker-образ и запускать контейнеризированный сервис в воспроизводимой среде

---

### Дерево проекта
```
├── .vscode
│   └── launch.json
├── auth
│   ├── certs
│   │   ├── ca.crt
│   │   ├── server.crt
│   │   └── server.key.ex
│   ├── cmd
│   │   └── main.go
│   ├── db
│   │   └── db.go
│   ├── handlers
│   │   └── handlers.go
│   ├── middleware
│   │   └── middleware.go
│   ├── monitoring
│   │   └── prometheus.yml
│   ├── utils
│   │   ├── env.go
│   │   ├── password.go
│   │   └── token.go
│   ├── .dockerignore
│   ├── .env.ex
│   ├── Dockerfile
│   ├── docker-compose.yml
│   ├── go.mod
│   └── go.sum
├── auth-sdk
│   ├── gen
│   │   ├── auth_requests.pb.go
│   │   ├── auth_responses.pb.go
│   │   ├── auth_service.pb.go
│   │   └── auth_service_grpc.pb.go
│   ├── proto
│   │   ├── auth_requests.proto
│   │   ├── auth_responses.proto
│   │   └── auth_service.proto
│   ├── go.mod
│   └── go.sum
├── db
│   ├── data
│   │   └── ...
│   ├── .dockerignore
│   ├── .env.ex
│   ├── Dockerfile
│   └── init.sql
├── screenshots
│   └── ...
├── task
│   ├── auth
│   │   └── auth.go
│   ├── certs
│   │   ├── ca.crt
│   │   ├── server.crt
│   │   └── server.key.ex
│   ├── cmd
│   │   └── main.go
│   ├── db
│   │   └── db.go
│   ├── dtos
│   │   ├── requests.go
│   │   └── responses.go
│   ├── handlers
│   │   └── handlers.go
│   ├── logger
│   │   └── logger.go
│   ├── middleware
│   │   └── middleware.go
│   ├── utils
│   │   └── utils.go
│   ├── .dockerignore
│   ├── .env.ex
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
├── task-sdk
│   ├── gen
│   │   ├── task_requests.pb.go
│   │   ├── task_responses.pb.go
│   │   ├── task_service.pb.go
│   │   └── task_service_grpc.pb.go
│   ├── proto
│   │   ├── task_requests.proto
│   │   ├── task_responses.proto
│   │   └── task_service.proto
│   ├── go.mod
│   └── go.sum
├── .gitignore
├── README.md
└── docker-compose.yml

29 directories, 62 files
```