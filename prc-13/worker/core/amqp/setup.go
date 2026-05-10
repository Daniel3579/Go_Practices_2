package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	MainQueue = "task_jobs"
	DLQ       = "task_jobs_dlq"
)

func DeclareQueues(ch *amqp.Channel) error {
	// Сначала объявляем DLQ (обычная durable очередь)
	_, err := ch.QueueDeclare(
		DLQ,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		return err
	}

	// Основная очередь с dead-letter настройкой
	args := amqp.Table{
		"x-dead-letter-exchange":    "",  // используем default exchange
		"x-dead-letter-routing-key": DLQ, // маршрутизация в DLQ
	}
	_, err = ch.QueueDeclare(
		MainQueue,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		args,
	)
	return err
}
