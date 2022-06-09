package util

import (
	"errors"
	"net/http"
)

var (
	MapErrorTypeToHTTPStatus = mapErrorTypeToHTTPStatus
	IsError                  = isError
	NewError                 = newError
)

var (
	ErrBadRequest       = errors.New("bad request")
	ErrInternal         = errors.New("internal error")
	ErrInvalidAPICall   = errors.New("invalid api call")
	ErrNotAuthenticated = errors.New("not authenticated")
	ErrResourceNotFound = errors.New("resource not found")
)

type ErrorResponse struct {
	ErrorCode int
	Cause     string
}

const (
	ErrorCodeInternal           = 0
	ErrorCodeInvalidJSONBody    = 30
	ErrorCodeInvalidCredentials = 201
	ErrorCodeEntityNotFound     = 404
	ErrorCodeValidation         = 500
)

type serverError struct {
	code      int
	cause     string
	errorType error
}

func (e serverError) Error() string {
	return e.cause
}

func mapErrorTypeToHTTPStatus(err error) int {
	switch err {
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrInternal:
		return http.StatusInternalServerError
	case ErrInvalidAPICall, ErrResourceNotFound:
		return http.StatusNotFound
	case ErrNotAuthenticated:
		return http.StatusNonAuthoritativeInfo
	default:
		return http.StatusInternalServerError
	}
}

func isError(errorType error) (bool, int, string, error) {
	err, isErr := errorType.(serverError)
	if !isErr {
		return false, 0, "", errorType
	}
	return true, err.code, err.cause, err.errorType
}

func newError(cause string, code int, errorType, err error) {

}
