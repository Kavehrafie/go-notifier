package repository

import (
	"context"
	"github.com/kavehrafie/go-scheduler/pkg/domain"
	"github.com/sirupsen/logrus"
)

type TaskRepository interface {
	Create(ctx context.Context, task *domain.Task, log *logrus.Logger) error
	ListPendingTasks(ctx context.Context) ([]*domain.Task, error)
	UpdateStatus(ctx context.Context, id string, status int) error
}
