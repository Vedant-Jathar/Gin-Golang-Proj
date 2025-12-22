package Utils

import "fmt"

type AppError struct {
	StatusCode int
	Message    string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

func NewError(status int, msg string) *AppError {
	return &AppError{
		StatusCode: status,
		Message:    msg,
	}
}
