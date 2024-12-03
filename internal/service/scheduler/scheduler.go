package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kavehrafie/go-scheduler/internal/model"
	"github.com/kavehrafie/go-scheduler/internal/store"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"time"
)

type Scheduler struct {
	store  store.Store
	log    *logrus.Logger
	client *http.Client
	cron   *cron.Cron
}

func NewScheduler(store store.Store, logger *logrus.Logger) *Scheduler {
	return &Scheduler{
		store:  store,
		log:    logger,
		client: &http.Client{Timeout: 30 * time.Second},
		cron:   cron.New(),
	}
}

func (s *Scheduler) Start(ctx context.Context) error {
	_, err := s.cron.AddFunc("* * * * *", func() {
		// execute this action on schedule
	})
	if err != nil {
		return fmt.Errorf("failed to schedule action: %v", err)
	}

	s.cron.Start()

	go func() {
		<-ctx.Done()
		s.Stop()
	}()

	return nil
}

func (s *Scheduler) Stop() {
	ctx := s.cron.Stop()
	<-ctx.Done()
}

func (s *Scheduler) processPendingRequests() error {
	query := `SELECT * FROM scheduled_requests WHERE status = ? AND scheduled_at <= ?`
	rows, err := s.store.Query(ctx, query, model.StatusPending, time.Now())

	for rows.Next() {
		var sq model.ScheduledRequest
		var headerJSON string

		err := rows.Scan(
			&sq.ID,
			&sq.Title,
			&sq.Description,
			&sq.URL,
			&sq.Payload,
			&headerJSON,
			&sq.ScheduledAt,
			&sq.Status,
			&sq.CreatedAt,
			&sq.ExecutedAt,
			&sq.DeletedAt,
		)
		if err != nil {
			s.log.WithError(err).Error("failed to scan scheduled request")
			continue
		}

		if err := json.Unmarshal([]byte(headerJSON), &sq.Header); err != nil {
			s.log.WithError(err).Error("failed to unmarshal scheduled request header")
			continue
		}

		if err := s.executeRequest(&sq); err != nil {
			s.log.WithError(err).WithField("request_id", sq.ID).Error("failed to execute scheduled request")
			_, err := s.db.Exec(`UPDATE scheduled_requests SET status = ?, error = ? WHERE id = ?`, model.StatusFailed, fmt.Sprintf("%v", err), sq.ID)
			if err != nil {
				s.log.WithError(err).Error("failed to update scheduled request status")
			}
		} else {
			_, err := s.db.Exec(`UPDATE scheduled_requests SET status = ?, executed_at = ? WHERE id = ?`, model.StatusResolved, time.Now(), sq.ID)
			if err != nil {
				s.log.WithError(err).Error("failed to update scheduled request status")
			}
		}
	}

	return err
}

func (s *Scheduler) executeRequest(sq *model.ScheduledRequest) error {
	req, err := http.NewRequest(http.MethodPost, sq.URL, strings.NewReader(string(sq.Payload)))
	if err != nil {
		return fmt.Errorf("failed to create http request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	for key, value := range sq.Header {
		req.Header.Set(key, value)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute action: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("webhook returned status code %d", resp.StatusCode)
	}

	return nil
}
