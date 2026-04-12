# Коляда Даниил
## Практическая работа №4

### Цель работы

Освоить базовую организацию мониторинга backend-приложения на Go с использованием Prometheus для сбора метрик и Grafana для их визуализации

---

## Развертываение Docker контейнеров

Создали docker-compose файл

Описали контейнеризацию Prometheus и Grafana

Произвели развертывание

![Screenshot](./screenshots/Screenshot_2.png)

---

## Проверка стасута сервиса

Приложение поднято. Статус UP

![Screenshot](./screenshots/Screenshot_3.png)

---

## Мониторинг и графики

Визуализация метрик

![Screenshot](./screenshots/Screenshot_1.png)

---

### Выводы

Освоили базовую организацию мониторинга backend-приложения на Go с использованием Prometheus для сбора метрик и Grafana для их визуализации

---

### Дерево проекта
```
├── README.md
├── auth
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

22 directories, 40 files
```