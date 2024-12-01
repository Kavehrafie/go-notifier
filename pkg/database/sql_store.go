package database

import "database/sql"

type SQLStore struct {
	db *sql.DB
}

func NewSQLStore(driver, dsn string) (*SQLStore, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &SQLStore{db: db}, nil
}

func (s *SQLStore) Close() error {
	return s.db.Close()
}
