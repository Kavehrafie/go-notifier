package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/kavehrafie/go-scheduler/internal/model"
	"github.com/kavehrafie/go-scheduler/pkg/database"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrInvalidDriver = errors.New("invalid database drive")
	ErrAlreadyExists = errors.New("already exists")
)

type Store interface {
	Create(ctx context.Context, sa *model.ScheduledRequest) error
	//Get(ctx context.Context, id string) (*model.ScheduledAction, error)
	//Update(ctx context.Context, sa *model.ScheduledAction) error
	//Delete(ctx context.Context, id string) error
	//
	//List(ctx context.Context, offset, limit int) ([]*model.ScheduledAction, error)
	//ListByStatus(ctx context.Context, status model.ScheduledActionStatus) ([]*model.ScheduledAction, error)
	//ListPending(ctx context.Context, before time.Time) ([]*model.ScheduledAction, error)

	Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)

	Close() error
	Ping(ctx context.Context) error
}

type StoreFactory interface {
	NewStore(config database.Config) (Store, error)
}
