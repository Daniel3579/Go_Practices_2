package amqp

import (
	"context"
	"encoding/json"
	"log"
	"prc-13/services/tasks/core/jobs"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type TaskEvent struct {
	Event  string `json:"event"`
	TaskID string `json:"task_id"`
	TS     string `json:"ts"`
}

type Publisher struct {
	channel   *amqp.Channel
	queueName string
}

func NewPublisher(ch *amqp.Channel, queueName string) *Publisher {
	return &Publisher{
		channel:   ch,
		queueName: queueName,
	}
}

// PublishTaskCreated публикует событие task.created.
// В случае ошибки только логирует (best effort), не возвращает ошибку вызывающему коду.
func (p *Publisher) PublishTaskCreated(taskID string) {
	msg := TaskEvent{
		Event:  "task.created",
		TaskID: taskID,
		TS:     time.Now().UTC().Format(time.RFC3339),
	}

	body, err := json.Marshal(msg)
	if err != nil {
		log.Printf("failed to marshal event: %v", err)
		return
	}

	err = p.channel.PublishWithContext(
		context.Background(),
		"",          // exchange
		p.queueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent, // persistent message
			Body:         body,
		},
	)
	if err != nil {
		log.Printf("failed to publish message: %v", err)
		// best effort: ошибка только логируется, задача уже создана
		return
	}
	log.Printf("event published: task_id=%s", taskID)
}

// Добавьте в существующий файл (рядом с PublishTaskCreated)
func (p *Publisher) PublishJob(queueName string, job *jobs.TaskJob) error {
	body, err := json.Marshal(job)
	if err != nil {
		return err
	}
	return p.channel.PublishWithContext(
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
