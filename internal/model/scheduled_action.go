package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type ScheduledActionStatus int

var validate *validator.Validate

const (
	StatusPending ScheduledActionStatus = 0
	StatusSent    ScheduledActionStatus = 1
	StatusFailed  ScheduledActionStatus = 2
	StatusDeleted ScheduledActionStatus = 3
)

type ScheduledAction struct {
	ID          string                `json:"id" db:"id"`
	Title       string                `json:"title" db:"title"`
	Description string                `json:"description" db:"description"`
	ScheduledAt time.Time             `json:"schedule_at" db:"schedule_at"`
	CreatedAt   time.Time             `json:"created_at" db:"created_at"`
	UpdateAt    time.Time             `json:"updated_at" db:"updated_at"`
	DeletedAt   time.Time             `json:"deleted_at" db:"deleted_at"`
	URL         string                `json:"url" db:"url"`
	Status      ScheduledActionStatus `json:"status" db:"status"`
	Payload     string                `json:"payload" db:"payload"`
	Metadata    map[string]string     `json:"metadata" db:"metadata"`
	Failures    int                   `json:"failures" db:"failures"`
}

type ScheduleActionRegisterInput struct {
	Title       string            `json:"title"  validate:"required,min=3,max=100"`
	Description string            `json:"description"  validate:"omitempty,max=500"`
	URL         string            `json:"url"  validate:"required,url"`
	Payload     string            `json:"payload"  validate:"required,json"`
	ScheduledAt time.Time         `json:"scheduled_at" validate:"required,future_time"`
	Metadata    map[string]string `json:"metadata"  validate:"omitempty,dive,keys,max=50,endkeys,max=200"`
}

type ValidationError struct {
	Field string
	Tag   string
	Value interface{}
}

func init() {
	validate = validator.New()

	_ = validate.RegisterValidation("future_time", validateFutureTime)
}

func validateFutureTime(fl validator.FieldLevel) bool {
	timeValue, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}
	return timeValue.After(time.Now())
}

func (input *ScheduleActionRegisterInput) Validate() []ValidationError {
	if err := validate.Struct(input); err != nil {
		var validationErrors []ValidationError

		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ValidationError{
				Field: err.Field(),
				Tag:   err.Tag(),
				Value: err.Value(),
			})
		}
		return validationErrors
	}
	return nil
}

func (input *ScheduleActionRegisterInput) ToSchedule() (*ScheduledAction, error) {
	if errs := input.Validate(); len(errs) > 0 {
		return nil, fmt.Errorf("validation failed: %v", errs)
	}

	return &ScheduledAction{
		ID:          uuid.NewString(),
		Status:      StatusPending,
		Title:       input.Title,
		Payload:     input.Payload,
		CreatedAt:   time.Now(),
		Metadata:    input.Metadata,
		ScheduledAt: input.ScheduledAt,
		URL:         input.URL,
	}, nil
}
