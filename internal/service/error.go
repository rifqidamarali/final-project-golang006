package service

import (
	"errors"
	"fmt"
)

var (
    ErrNotFound   = errors.New("resource not found")
    ErrForbidden  = errors.New("forbidden")
    ErrValidation = errors.New("validation error")
)

type CustomError struct {
	Base    error
	Message string
}

// Error implements the error interface.
func (e *CustomError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("%s: %s", e.Base.Error(), e.Message)
	}
	return e.Base.Error()
}

// NewCustomError creates a new CustomError.
func NewCustomError(base error, message string) *CustomError {
	return &CustomError{Base: base, Message: message}
}

func (e *CustomError) Unwrap() error {
    return e.Base
}

// Is checks if an error matches the base error.
// func Is(err, target error) bool {
// 	if customErr, ok := err.(*CustomError); ok {
// 		return customErr.Base == target
// 	}
// 	return err == target
// }