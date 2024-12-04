package repository

import (
	"context"
	"github.com/kavehrafie/go-scheduler/pkg/domain"
)

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task) error
	ListPendingTasks(ctx context.Context) ([]*domain.Task, error)
	UpdateStatus(ctx context.Context, id string, status int) error
}
