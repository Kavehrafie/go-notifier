package repository

import (
	"errors"
	"github.com/kavehrafie/go-scheduler/pkg/database"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	GetTaskRepository() TaskRepository
}

func NewRepository(db *database.Store, log *logrus.Logger) (Repository, error) {
	switch (*db).GetDriverName() {
	case "sqlite":
		return NewSQLiteRepository(db, log)
	}

	return nil, errors.New("driver not supported")
}
