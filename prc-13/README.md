# Коляда Даниил
## Практическая работа №13

### Цель работы

Освоить базовую работу с брокером сообщений RabbitMQ в приложении на Go, научиться публиковать сообщения в очередь, реализовывать отдельного потребителя сообщений, подтверждать успешную обработку через ack и понимать назначение очередей сообщений в микросервисной архитектуре

---

### Шаги

Запуск RabbitMQ
```
docker compose up -d
```

После запуска management UI доступен по адресу  
`http://localhost:15672`

---

Запуск worker
```
go run ./cmd/worker
```

---

Запуск tasks
```
go run ./cmd/tasks
```

---

Отправляем запрос
![Screenshot](./screenshots/Screenshot_1.png)

---

В worker появляется лог вида
```
2026/05/10 19:47:32 worker started, waiting for messages...
2026/05/10 19:49:13 received event=task.created task_id=t_1778431753498821000 ts=2026-05-10T16:49:13Z
```

---

Интерфейс RabbitMQ
![Screenshot](./screenshots/Screenshot_2.png)
![Screenshot](./screenshots/Screenshot_3.png)

---

### Выводы

Освоили базовую работу с брокером сообщений RabbitMQ в приложении на Go, научились публиковать сообщения в очередь, реализовывать отдельного потребителя сообщений, подтверждать успешную обработку через ack и понимать назначение очередей сообщений в микросервисной архитектуре

---

### Дерево проекта

```
├── README.md
├── deploy
│   └── rabbit
│       └── docker-compose.yml
├── go.mod
├── go.sum
├── screenshots
│   └── ...
├── services
│   └── tasks
│       ├── amqpclient
│       │   └── amqpclient.go
│       ├── cmd
│       │   └── tasks
│       │       └── main.go
│       ├── core
│       │   ├── amqp
│       │   │   └── publisher.go
│       │   ├── network
│       │   │   └── handler.go
│       │   └── service
│       │       └── task.go
│       ├── events
│       │   └── events.go
│       └── publisher
│           └── publisher.go
└── worker
    └── cmd
        └── worker
            └── main.go

18 directories, 15 files
```