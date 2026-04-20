# Коляда Даниил
## Практическая работа №8

### Цель работы

Освоить основы CI/CD для backend-проекта на Go, научиться настраивать автоматический pipeline для проверки, сборки, упаковки Docker-образа и подготовки приложения к доставке

---

### Описание

CI – это Continuous Integration, то есть непрерывная интеграция<br>
Смысл в том, что после каждого изменения кода система автоматически проверяет проект:
- устанавливает зависомости
- запускает тесты
- выполняет сборку
- проверяет, что код не сломал проект

CD – это Continuous Deployment, то есть непрерывное развертывание<br>
Отвечает за упаковку и доставку результата – например, публикацию Docker-образа и деплой на сервер

---

### Структура pipeline

- [main.yml](.github/workflows/main.yml)
    - [test-and-build.yml](.github/workflows/test-and-build.yml)
    - [build-and-push-docker.yml](.github/workflows/build-and-push-docker.yml)

---

### Создали токен

![Screenshot](./screenshots/Screenshot_1.png)
![Screenshot](./screenshots/Screenshot_4.png)

---

### Cоздали секрет

![Screenshot](./screenshots/Screenshot_2.png)

---

### Успешная сборка

![Screenshot](./screenshots/Screenshot_3.png)

---

### Образы опубликованы в registry

![Screenshot](./screenshots/Screenshot_5.png)

---

### Выводы

Освоили основы CI/CD для backend-проекта на Go, научились настраивать автоматический pipeline для проверки, сборки, упаковки Docker-образа и подготовки приложения к доставке

---

### Дерево проекта

```
├── .github
│   └── workflows
│       ├── build-and-push-docker.yml
│       ├── main.yml
│       └── test-and-build.yml
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
│   │   ├── token.go
│   │   └── utils_test.go
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
├── postgres
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
│   │   ├── utils.go
│   │   └── utils_test.go
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

31 directories, 70 files
```