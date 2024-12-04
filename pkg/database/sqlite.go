package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

type SQliteStore struct {
	db     *sql.DB
	driver string
}

func (db *SQliteStore) GetDB() *sql.DB {
	return db.db
}

func (db *SQliteStore) GetDriverName() string {
	return db.driver
}

func NewSQLiteStore(log *logrus.Logger) (*SQliteStore, error) {
	db, err := sql.Open("sqlite", "./db/sqlite.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &SQliteStore{db: db, driver: "sqlite"}, err
}
