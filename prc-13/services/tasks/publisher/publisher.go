package publisher

import (
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type TaskEvent struct {
	Event  string `json:"event"`
	TaskID string `json:"task_id"`
	TS     string `json:"ts"`
}

func PublishTaskCreated(ch *amqp.Channel, queueName, taskID string) error {
	_, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msg := TaskEvent{
		Event:  "task.created",
		TaskID: taskID,
		TS:     time.Now().UTC().Format(time.RFC3339),
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return ch.PublishWithContext(
		context.Background(),
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
}
