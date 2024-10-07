package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	Astatus  int    `json:"status"`
	Amessage string `json:"message"`
	Aerror   string `json:"error,omitempty"`
}

func (a *apiError) Status() int {
	return a.Astatus
}
func (a *apiError) Message() string {
	return a.Amessage
}
func (a *apiError) Error() string {
	return a.Aerror
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		Astatus:  statusCode,
		Amessage: message,
	}
}

func NewApiErrorFromBytes(body []byte) (ApiError, error) {
	var result apiError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}
	return &result, nil
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		Astatus:  http.StatusNotFound,
		Amessage: message,
	}
}
func NewInternalServerError(message string) ApiError {
	return &apiError{
		Astatus:  http.StatusInternalServerError,
		Amessage: message,
	}
}
func NewBadRequestError(message string) ApiError {
	return &apiError{
		Astatus:  http.StatusBadRequest,
		Amessage: message,
	}
}
