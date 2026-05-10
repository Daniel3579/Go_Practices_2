# Коляда Даниил
## Практическая работа №11

### Цель работы

Освоить разработку GraphQL API на языке Go с использованием библиотеки gqlgen, научиться описывать GraphQL-схему, генерировать серверный каркас приложения, реализовывать резолверы для запросов и мутаций, а также тестировать работу API

---

### Шаги

Запрос списка задач
![Screenshot](./screenshots/Screenshot_1.png)

---

Запрос одной задачи по идентификатору
![Screenshot](./screenshots/Screenshot_2.png)

---

Создание задачи
![Screenshot](./screenshots/Screenshot_3.png)

---

Обновление задачи
![Screenshot](./screenshots/Screenshot_4.png)

---

Удаление Задачи
![Screenshot](./screenshots/Screenshot_5.png)

---

### Выводы

Освоили разработку GraphQL API на языке Go с использованием библиотеки gqlgen, научились описывать GraphQL-схему, генерировать серверный каркас приложения, реализовывать резолверы для запросов и мутаций, а также тестировать работу API

---

### Дерево проекта

```
├── README.md
├── go.mod
├── go.sum
├── gqlgen.yml
├── graph
│   ├── generated.go
│   ├── model
│   │   └── models_gen.go
│   ├── resolver.go
│   ├── schema.graphql
│   └── schema.resolvers.go
├── screenshots
│   └── ...
├── server.go
└── services
    └── graphql
        ├── cmd
        │   └── graphql
        │       └── main.go
        ├── handlers
        │   └── handlers.go
        └── store
            └── store.go

10 directories, 18 files
```