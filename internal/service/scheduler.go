package service

import (
	"context"
	"github.com/kavehrafie/go-scheduler/internal/repository"
	"github.com/kavehrafie/go-scheduler/pkg/domain"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"sync"
	"time"
)

type SchedulerService struct {
	repo repository.TaskRepository
	stop chan struct{}
	log  *logrus.Logger
	wg   sync.WaitGroup
}

func NewSchedulerService(repo repository.TaskRepository, log *logrus.Logger) *SchedulerService {
	return &SchedulerService{
		repo: repo,
		log:  log,
		stop: make(chan struct{}),
	}
}

func (s *SchedulerService) Start() {
	s.log.Info("scheduler service started")
	s.wg.Add(1)
	go s.processPendingTasks()
}

func (s *SchedulerService) Stop() {
	close(s.stop)
	s.wg.Wait()
	s.log.Info("scheduler service stopped")
}

func (s *SchedulerService) processPendingTasks() {
	defer s.wg.Done()
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	s.executePendingTasks() // execute immediately on start

	for {
		select {
		case <-ticker.C:
			s.executePendingTasks()
		case <-s.stop:
			s.log.Info("stopping pending tasks processor")
			return
		}
	}
}

func (s *SchedulerService) executePendingTasks() {
	s.log.Info("executing pending tasks")
	ctx := context.Background()
	tasks, err := s.repo.ListPendingTasks(ctx)
	s.log.Infof("found %d pending tasks", len(tasks))
	if err != nil {
		s.log.Errorf("failed to list pending tasks: %v", err)
		return
	}
	for _, task := range tasks {
		go s.executeTask(ctx, &task)
	}
}

func (s *SchedulerService) executeTask(ctx context.Context, task *domain.Task) {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("POST", task.URL, strings.NewReader(task.Payload))
	if err != nil {
		s.log.Errorf("error creating request for task %v: %v", task.ID, err)
		err = s.repo.UpdateStatus(ctx, task.ID, domain.TaskStatusError)
		s.log.Errorf("error updating status for task %v: %v", task.ID, err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("error executing task %v: %v", task.ID, err)
	}
	defer resp.Body.Close()

	s.log.Infof("task %v executed with status code %v", task.ID, resp.StatusCode)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		s.repo.UpdateStatus(ctx, task.ID, domain.TaskStatusCompleted)
	} else {
		s.repo.UpdateStatus(ctx, task.ID, domain.TaskStatusError)
	}

}
