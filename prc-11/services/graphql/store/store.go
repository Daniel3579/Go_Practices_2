package store

type Task struct {
	ID          string
	Title       string
	Description *string
	Done        bool
}

var Tasks = []*Task{
	{ID: "t_001", Title: "Изучить REST", Description: strPtr("Прочитать про REST API"), Done: false},
	{ID: "t_002", Title: "Изучить GraphQL", Description: strPtr("Разобраться со схемами"), Done: true},
}

func strPtr(s string) *string {
	return &s
}
