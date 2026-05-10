package main

import (
	"log"
	"os"

	wamqp "prc-13/worker/core/amqp"
	"prc-13/worker/core/consumer"
	"prc-13/worker/core/store"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	rabbitURL := os.Getenv("RABBIT_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@localhost:5672/"
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to open channel: %v", err)
	}
	defer ch.Close()

	// Объявляем очереди
	if err := wamqp.DeclareQueues(ch); err != nil {
		log.Fatalf("queue declare error: %v", err)
	}

	// Prefetch = 1 (одно сообщение за раз)
	if err := ch.Qos(1, 0, false); err != nil {
		log.Fatalf("qos error: %v", err)
	}

	msgs, err := ch.Consume(
		wamqp.MainQueue,
		"",
		false, // auto-ack = false (ручное подтверждение)
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("consume error: %v", err)
	}

	processed := store.NewProcessedStore()
	log.Println("worker started, waiting for jobs...")

	for d := range msgs {
		consumer.ProcessMessage(d, processed, ch)
	}
}
