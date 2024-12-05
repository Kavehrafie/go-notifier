package handlers

import (
	"github.com/kavehrafie/go-scheduler/internal/repository"
	"github.com/kavehrafie/go-scheduler/pkg/domain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Handler struct {
	repo repository.Repository
	log  *logrus.Logger
}

func NewHandler(repo repository.Repository, log *logrus.Logger) *Handler {
	return &Handler{
		repo: repo,
		log:  log,
	}
}

func (h *Handler) RegisterRoutes(e *echo.Echo) {

	rateLimiter := middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20))

	api := e.Group("/api/v1")
	api.Use(rateLimiter)

	task := e.Group("/tasks")
	task.POST("", h.CreateTask)
	task.GET("", h.ListTasks)
	task.GET("/:id", h.GetTask)

}

func (h *Handler) CreateTask(c echo.Context) error {
	var input domain.TaskCreateInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request body"})
	}

	// todo: validate

	tr := h.repo.GetTaskRepository()
	task := &domain.Task{
		URL:       input.URL,
		Payload:   input.Payload,
		ExecuteAt: time.Now().Add(time.Second * input.After),
	}
	if err := tr.Create(c.Request().Context(), task); err != nil {
		h.log.Errorf("failed to create task: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to create task"})
	}

	return c.JSON(http.StatusCreated, task)
}

func (h *Handler) ListTasks(c echo.Context) error {

}

func (h *Handler) GetTask(c echo.Context) error {

}
