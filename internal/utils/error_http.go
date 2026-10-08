package utils

import "fmt"

type ApiError struct {
	StatusCode int `json:"status"`
	Message    any `json:"message"`
}

func (a ApiError) Error() string {
	return fmt.Sprintf("api error: %w", a.StatusCode)
}

func NewApiError(statusCode int, err error) ApiError {
	return ApiError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}
