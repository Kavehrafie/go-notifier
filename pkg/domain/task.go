package domain

import (
	"database/sql"
	"time"
)

type Task struct {
	ID        string       `json:"id"`
	URL       string       `json:"url"`
	Payload   string       `json:"payload"`
	ExecuteAt sql.NullTime `json:"execute_at"`
	Status    TaskStatus   `json:"status"` // pending, completed, failed
	CreatedAt time.Time    `json:"created_at"`
}

type TaskStatus int

var (
	TaskStatusRunning   TaskStatus = 1
	TaskStatusCompleted TaskStatus = 2
	TaskStatusError     TaskStatus = 3
	TaskStatusPaused    TaskStatus = 4
)
