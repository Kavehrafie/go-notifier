package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/kavehrafie/go-scheduler/pkg/database"
	"github.com/kavehrafie/go-scheduler/pkg/domain"
	"time"
)

func newSQLiteTaskRepository(store *database.Store) (*SQLiteTaskRepository, error) {
	db := (*store).GetDB()
	if db == nil {
		return nil, errors.New("database pointer is empty")
	}
	_, err := db.Exec(`
CREATE TABLE IF NOT EXISTS tasks (
    id TEXT PRIMARY KEY,
    url TEXT NOT NULL,
    payload TEXT NOT NULL,
    execute_at DATETIME NOT NULL,
    status INTEGER NOT NULL,
    create_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP);
	
	CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status);
	`)
	if err != nil {
		return nil, err
	}

	return &SQLiteTaskRepository{
		db: db,
	}, nil
}

type SQLiteTaskRepository struct {
	db *sql.DB
}

func (s *SQLiteTaskRepository) Create(ctx context.Context, input *domain.TaskCreateInput) error {
	query := `INSERT INTO tasks (id, url, payload, execute_at, status) 
				VALUES (?, ?, ?, ?, ?)`

	task := domain.Task{
		ID:        uuid.NewString(),
		Status:    domain.TaskStatusRunning,
		CreatedAt: time.Now(),
		Payload:   input.Payload,
		ExecuteAt: time.Now().Add(time.Second * time.Duration(input.After)),
		URL:       input.URL,
	}

	result, err := s.db.ExecContext(ctx, query,
		task.ID,
		task.URL,
		task.Payload,
		task.ExecuteAt,
		task.Status,
	)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (s *SQLiteTaskRepository) ListPendingTasks(ctx context.Context) ([]domain.Task, error) {
	query := `SELECT id, url, payload, execute_at, status FROM tasks WHERE status = ? AND execute_at <= ?`

	rows, err := s.db.QueryContext(ctx, query, domain.TaskStatusRunning, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ts []domain.Task
	for rows.Next() {
		var t domain.Task
		err = rows.Scan(
			&t.ID,
			&t.URL,
			&t.Payload,
			&t.ExecuteAt,
			&t.Status,
		)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}

	return ts, nil
}

func (s *SQLiteTaskRepository) UpdateStatus(ctx context.Context, id string, status domain.TaskStatus) error {
	query := `UPDATE tasks SET status = ? WHERE id = ?`
	res, err := s.db.ExecContext(ctx, query, status, id)
	if err != nil {
		return err
	}

	num, _ := res.RowsAffected()
	if num < 1 {
		return errors.New("task not found")
	}
	return nil
}
