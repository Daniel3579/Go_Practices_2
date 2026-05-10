package jobs

type TaskJob struct {
	Job       string `json:"job"`        // тип задачи, например "process_task"
	TaskID    string `json:"task_id"`    // идентификатор бизнес-объекта
	Attempt   int    `json:"attempt"`    // номер попытки (начинается с 1)
	MessageID string `json:"message_id"` // уникальный ID сообщения
}
