package errs

import "net/http"

type AppError struct {
	Message string `json:"message"`
	Code int	`json:"omitempty`
}

func NewNotFoundError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusNotFound}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusServiceUnavailable}
}

func(err AppError) AsMessage() *AppError{
	return &AppError{Message: err.Message}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code: http.StatusUnprocessableEntity,
	}
}