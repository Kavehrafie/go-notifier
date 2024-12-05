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
