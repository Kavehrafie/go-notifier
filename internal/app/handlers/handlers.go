package handlers

import (
	"github.com/kavehrafie/go-scheduler/internal/repository"
	"github.com/kavehrafie/go-scheduler/pkg/domain"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"net/http"
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
	e.POST("/tasks", h.CreateTask)

	e.GET("/dummy", h.GetDummy)
}

func (h *Handler) GetDummy(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"msg": "This is a dummy url",
	})
}

func (h *Handler) CreateTask(c echo.Context) error {
	var task domain.TaskCreateInput
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	tr := h.repo.GetTaskRepository()
	if err := tr.Create(c.Request().Context(), &task); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, task)
}
