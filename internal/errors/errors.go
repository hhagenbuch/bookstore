package errors

import (
	"net/http"
)

type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	return e.Message
}

func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func NewBadRequestError(message string) *AppError {
	return NewAppError(http.StatusBadRequest, message, nil)
}

func NewInternalServerError(err error) *AppError {
	return NewAppError(http.StatusInternalServerError, "Internal server error", err)
}

func NewNotFoundError(message string) *AppError {
	return NewAppError(http.StatusNotFound, message, nil)
}
