package errors

import "errors"

// custom error types

var (
	ErrNotFound = errors.New("entity not found")
	ErrInvalid  = errors.New("invalid entity")
)
