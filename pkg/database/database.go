package database

import (
	"database/sql"
	"fmt"
	"github.com/kavehrafie/go-scheduler/pkg/config"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

type Store interface {
	GetDB() *sql.DB
	GetDriverName() string
}

func NewStore(cfg *config.Config, log *logrus.Logger) (Store, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config is nil")
	}

	if cfg.DB.Driver == "sqlite3" || cfg.DB.Driver == "sqlite" {
		return NewSQLiteStore(log)
	}

	return nil, fmt.Errorf("unknown driver %s", cfg.DB.Driver)
}
