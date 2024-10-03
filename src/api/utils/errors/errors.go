package errors

import "net/http"

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	status  int    `json:"status"`
	message string `json:"message"`
	error   string `json:"error, omitempty"`
}

func (a *apiError) Status() int {
	return a.status
}
func (a *apiError) Message() string {
	return a.message
}
func (a *apiError) Error() string {
	return a.error
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		status:  http.StatusNotFound,
		message: message,
	}
}
func NewInternalServerError(message string) ApiError {
	return &apiError{
		status:  http.StatusInternalServerError,
		message: message,
	}
}
func NewBadRequestError(message string) ApiError {
	return &apiError{
		status:  http.StatusBadRequest,
		message: message,
	}
}
