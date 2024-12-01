package repository

import (
	"context"
	"errors"
	"github.com/kavehrafie/go-notify/pkg/database"
)

var ErrNotFound = errors.New("notification not found")

type NotificationRepository struct {
	store database.Store
}

func NewNotificationRepository(store database.Store) *NotificationRepository {
	return &NotificationRepository{store}
}

func (r NotificationRepository) Create(ctx context.Context) {

}
