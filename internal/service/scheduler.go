package service

import (
	"context"
	"github.com/kavehrafie/go-scheduler/internal/repository"
	"github.com/kavehrafie/go-scheduler/pkg/domain"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"sync"
	"time"
)

type SchedulerService struct {
	repo     repository.TaskRepository
	stop     chan struct{}
	workers  chan struct{} // Semaphore for worker goroutines
	log      *logrus.Logger
	stopOnce sync.Once
}

const (
	maxWorkers    = 10
	taskTimeout   = 30 * time.Second
	checkInterval = 10 * time.Second // config
)

func NewSchedulerService(repo repository.TaskRepository, log *logrus.Logger) *SchedulerService {
	return &SchedulerService{
		repo:    repo,
		log:     log,
		stop:    make(chan struct{}),
		workers: make(chan struct{}, maxWorkers), // limit number of workers
	}
}

func (s *SchedulerService) Start(ctx context.Context) {
	s.log.Info("scheduler service started")
	go s.processPendingTasks(ctx)
}

func (s *SchedulerService) Stop() {
	s.stopOnce.Do(func() {
		close(s.stop)
		s.log.Info("scheduler service stopped")
	})
}

func (s *SchedulerService) processPendingTasks(ctx context.Context) {
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	s.executePendingTasks(ctx) // initial check
	for {
		select {
		case <-ticker.C:
			nextCheck := s.executePendingTasks(ctx)
			ticker.Reset(nextCheck)
		case <-s.stop:
			s.log.Info("stopping pending tasks processor")
			return
		case <-ctx.Done():
			return
		}
	}
}

func (s *SchedulerService) executePendingTasks(ctx context.Context) time.Duration {
	tasks, err := s.repo.ListPendingTasks(ctx)
	s.log.Infof("found %d pending tasks", len(tasks))
	if err != nil {
		s.log.Errorf("failed to list pending tasks: %v", err)
		return checkInterval
	}

	if len(tasks) == 0 {
		return checkInterval
	}

	var nextCheck time.Duration = checkInterval
	for _, task := range tasks {
		select {
		case s.workers <- struct{}{}: // acquire a worker
			go func(t domain.Task) {
				defer func() { <-s.workers }() // release the worker
				s.executeTask(ctx, &t)
			}(task)
		default:
			s.log.Warn("max workers reached, waiting for next cycle.")
			return checkInterval / 2 // check sooner next time
		}
	}

	return nextCheck
}

func (s *SchedulerService) executeTask(ctx context.Context, task *domain.Task) {
	taskCtx, cancel := context.WithTimeout(ctx, taskTimeout)
	defer cancel()

	client := &http.Client{
		Timeout: taskTimeout,
	}

	req, err := http.NewRequestWithContext(taskCtx, "POST", task.URL, strings.NewReader(task.Payload))
	if err != nil {
		s.handleTaskError(ctx, task, err, "error creating request")
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		s.handleTaskError(ctx, task, err, "error executing request")
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if err := s.repo.UpdateStatus(ctx, task.ID, domain.TaskStatusCompleted); err != nil {
			s.log.Errorf("failed to update task status: %v", err)
		}
		s.log.Infof("task %v completed successfully", task.ID)
	} else {
		s.handleTaskError(ctx, task, nil,
			"task failed with status code: %d", resp.StatusCode)
	}
}

func (s *SchedulerService) handleTaskError(ctx context.Context, task *domain.Task, err error, format string, args ...interface{}) {
	if err != nil {
		s.log.Errorf(format+": %v", append(args, err)...)
	} else {
		s.log.Errorf(format, args...)
	}

	if updateErr := s.repo.UpdateStatus(ctx, task.ID, domain.TaskStatusError); updateErr != nil {
		s.log.Errorf("failed to update task status: %v", updateErr)
	}
}
