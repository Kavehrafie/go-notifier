package repository

import (
	"github.com/kavehrafie/go-scheduler/pkg/database"
	"github.com/sirupsen/logrus"
)

type SQLiteRepository struct {
	TaskRepo *SQLiteTaskRepository
}

func NewSQLiteRepository(store *database.Store, log *logrus.Logger) (*SQLiteRepository, error) {
	task, err := newSQLiteTaskRepository(store)
	log.Info(task)
	if err != nil {
		return nil, err
	}
	return &SQLiteRepository{
		TaskRepo: task,
	}, nil
}

func (sp *SQLiteRepository) GetTaskRepository() TaskRepository {
	return sp.TaskRepo
}

//func (sp *SQLiteRepository) Create(ctx context.Context, task *domain.Task) error { return nil }
//func (sp *SQLiteRepository) ListPendingTasks(ctx context.Context) ([]*domain.Task, error) {
//	return nil, nil
//}
//func (sp *SQLiteRepository) UpdateStatus(ctx context.Context, id string, status int) error {
//	return nil
//}
