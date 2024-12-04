package repository

import (
	"errors"
	"github.com/kavehrafie/go-scheduler/pkg/database"
)

type Repository interface {
}

func NewTaskRepository(db *database.Store) (Repository, error) {
	switch (*db).GetDriverName() {
	case "sqlite":
		return NewSQLiteRepository(db)
	}

	return nil, errors.New("driver not supported")
}
