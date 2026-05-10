package main

import (
	"log"
	"net/http"
	"os"
	"prc-13/services/tasks/core/amqp"
	"prc-13/services/tasks/core/network"
	"prc-13/services/tasks/core/service"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Читаем переменные окружения
	rabbitURL := os.Getenv("RABBIT_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@localhost:5672/"
	}
	queueName := os.Getenv("QUEUE_NAME")
	if queueName == "" {
		queueName = "task_events"
	}

	// Подключаемся к RabbitMQ
	conn, err := amqp091.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel: %v", err)
	}
	defer ch.Close()

	// Объявляем очередь (durable = true)
	_, err = ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("failed to declare queue: %v", err)
	}

	// Инициализируем издателя
	publisher := amqp.NewPublisher(ch, queueName)

	// Инициализируем сервис задач и HTTP-обработчик
	taskService := service.NewTaskService()
	taskHandler := network.NewTaskHandler(taskService, publisher)

	// Настраиваем маршруты
	mux := http.NewServeMux()
	mux.HandleFunc("POST /v1/tasks", taskHandler.CreateTask)

	// Внутри main() после создания publisher и taskHandler:
	jobHandler := network.NewJobHandler(publisher, "task_jobs") // имя очереди

	mux.HandleFunc("POST /v1/jobs/process-task", jobHandler.ProcessTask)

	// Запускаем сервер
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	log.Printf("tasks service listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
