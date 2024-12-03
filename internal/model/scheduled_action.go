package model

import (
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"time"
)

type ScheduledRequestStatus int

var validate *validator.Validate

const (
	StatusPending  ScheduledRequestStatus = 0
	StatusResolved ScheduledRequestStatus = 1
	StatusFailed   ScheduledRequestStatus = 2
)

type ScheduledRequest struct {
	ID          string                 `json:"id" db:"id"`
	Title       string                 `json:"title" db:"title"`
	Description string                 `json:"description" db:"description"`
	ScheduledAt time.Time              `json:"schedule_at" db:"schedule_at"`
	CreatedAt   time.Time              `json:"created_at" db:"created_at"`
	ExecutedAt  sql.NullTime           `json:"executed_at" db:"executed_at"`
	DeletedAt   sql.NullTime           `json:"deleted_at" db:"deleted_at"`
	Suspend     bool                   `json:"suspend" db:"suspend"`
	URL         string                 `json:"url" db:"url"`
	Header      map[string]string      `json:"header" db:"header"`
	Status      ScheduledRequestStatus `json:"status" db:"status"`
	Payload     json.RawMessage        `json:"payload" db:"payload"`
	Error       string                 `json:"error" db:"error,omitempty"`
}

type ScheduledRequestRegisterInput struct {
	Title       string            `json:"title"  validate:"required,min=3,max=100"`
	Description string            `json:"description"  validate:"omitempty,max=500"`
	URL         string            `json:"url"  validate:"required,url"`
	Payload     string            `json:"payload"  validate:"required,json"`
	ScheduledAt time.Time         `json:"scheduled_at" validate:"required,future_time"`
	Header      map[string]string `json:"header"  validate:"omitempty,dive,keys,max=50,endkeys,max=200"`
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

//func (input *ScheduleActionRegisterInput) Validate() []ValidationError {
//	if err := validate.Struct(input); err != nil {
//		var validationErrors []ValidationError
//
//		for _, err := range err.(validator.ValidationErrors) {
//			validationErrors = append(validationErrors, ValidationError{
//				Field: err.Field(),
//				Tag:   err.Tag(),
//				Value: err.Value(),
//			})
//		}
//		return validationErrors
//	}
//	return nil
//}

//func (input *ScheduleActionRegisterInput) ToSchedule() (*ScheduledAction, error) {
//	if errs := input.Validate(); len(errs) > 0 {
//		return nil, fmt.Errorf("validation failed: %v", errs)
//	}
//
//	return &ScheduledAction{
//		ID:          uuid.NewString(),
//		Status:      StatusPending,
//		Title:       input.Title,
//		Payload:     input.Payload,
//		CreatedAt:   time.Now(),
//		Metadata:    input.Metadata,
//		ScheduledAt: input.ScheduledAt,
//		URL:         input.URL,
//	}, nil
//}
