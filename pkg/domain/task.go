package domain

import (
	"time"
)

type Task struct {
	ID        string     `json:"id"`
	URL       string     `json:"url"`
	Payload   string     `json:"payload"`
	ExecuteAt time.Time  `json:"execute_at"`
	Status    TaskStatus `json:"status"` // pending, completed, failed
	CreatedAt time.Time  `json:"created_at"`
}

type TaskStatus int

var (
	TaskStatusRunning   TaskStatus = 1
	TaskStatusCompleted TaskStatus = 2
	TaskStatusError     TaskStatus = 3
	TaskStatusPaused    TaskStatus = 4
)

type TaskCreateInput struct {
	URL     string `json:"url"`
	Payload string `json:"payload"`
	//ExecuteAt time.Time `json:"execute_at"`
	After int8 `json:"after"`
}
