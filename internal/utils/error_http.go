package utils

import (
	"encoding/json"
	"net/http"
)

type ApiError struct {
	StatusCode int `json:"status"`
	Message    any `json:"message"`
}

func NewApiError(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	erro := ApiError{
		StatusCode: status,
		Message:    err.Error(),
	}
	json.NewEncoder(w).Encode(erro)
}
