package model

import (
	"database/sql"
	"github.com/google/uuid"
)

type NotificationStatus int

const (
	StatusPending NotificationStatus = 0
	StatusSent    NotificationStatus = 1
	StatusFailed  NotificationStatus = 2
	StatusDeleted NotificationStatus = 3
)

type Notification struct {
	ID          string       `json:"id" db:"id"`
	Content     string       `json:"content" db:"content"`
	ScheduledAt sql.NullTime `json:"schedule_at" db:"schedule_at"`
	SentAt      sql.NullTime `json:"sent_at" db:"sent_at"`
	CreatedAt   sql.NullTime `json:"created_at" db:"created_at"`
	UpdateAt    sql.NullTime `json:"updated_at" db:"updated_at"`
	DeletedAt   sql.NullTime `json:"deleted_at" db:"deleted_at"`
}

type NotificationCreateRequest struct {
	Content     string       `json:"content" binding:"required"`
	ScheduledAt sql.NullTime `json:"schedule_at" binding:"required"`
}

func NewNotification(req NotificationCreateRequest) *Notification {
	return &Notification{
		ID:          uuid.NewString(),
		Content:     req.Content,
		ScheduledAt: req.ScheduledAt,
	}
}
