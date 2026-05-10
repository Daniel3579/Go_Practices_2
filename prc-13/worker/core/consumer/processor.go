package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	wamqp "prc-13/worker/core/amqp"
	"prc-13/worker/core/store" // замените на ваш путь

	amqp "github.com/rabbitmq/amqp091-go"
)

type TaskJob struct {
	Job       string `json:"job"`
	TaskID    string `json:"task_id"`
	Attempt   int    `json:"attempt"`
	MessageID string `json:"message_id"`
}

const maxAttempts = 3

func ProcessMessage(d amqp.Delivery, processed *store.ProcessedStore, ch *amqp.Channel) {
	var job TaskJob
	if err := json.Unmarshal(d.Body, &job); err != nil {
		log.Printf("invalid message: %v", err)
		_ = d.Ack(false)
		return
	}

	if processed.Exists(job.MessageID) {
		log.Printf("duplicate %s, ack", job.MessageID)
		_ = d.Ack(false)
		return
	}

	err := doBusinessLogic(job)
	if err == nil {
		processed.MarkDone(job.MessageID)
		_ = d.Ack(false)
		log.Printf("success: task_id=%s attempt=%d", job.TaskID, job.Attempt)
		return
	}

	log.Printf("error: task_id=%s attempt=%d: %v", job.TaskID, job.Attempt, err)
	job.Attempt++

	if job.Attempt <= maxAttempts {
		body, _ := json.Marshal(job)
		_ = ch.PublishWithContext(
			context.Background(), // заменили d.GetContext()
			"",
			wamqp.MainQueue,
			false,
			false,
			amqp.Publishing{
				ContentType:  "application/json",
				DeliveryMode: amqp.Persistent,
				Body:         body,
			},
		)
		_ = d.Ack(false)
		return
	}

	body, _ := json.Marshal(job)
	_ = ch.PublishWithContext(
		context.Background(), // заменили d.GetContext()
		"",
		wamqp.DLQ,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
	_ = d.Ack(false)
	log.Printf("moved to DLQ: task_id=%s", job.TaskID)
}

func doBusinessLogic(job TaskJob) error {
	time.Sleep(2 * time.Second)
	if job.TaskID == "t_fail" {
		return fmt.Errorf("simulated failure")
	}
	return nil
}
