package repository

import (
	"github.com/kavehrafie/go-scheduler/pkg/database"
)

type SQLiteRepository struct {
	Task TaskRepository
}

func NewSQLiteRepository(store *database.Store) (Repository, error) {
	task, err := newSQLiteTaskRepository(store)
	if err != nil {
		return nil, err
	}
	return &SQLiteRepository{
		Task: task,
	}, nil
}
