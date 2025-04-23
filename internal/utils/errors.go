package utils

import (
	"errors"
	"net/http"
)

var (
	ErrUnauthorized     = errors.New("unauthorized")
	ErrForbidden        = errors.New("forbidden")
	ErrNotFound         = errors.New("not found")
	ErrInvalidInput     = errors.New("invalid input")
	ErrInternal         = errors.New("internal server error")
	ErrDeadlineInPast   = errors.New("deadline cannot be in the past")
	ErrNotOwner         = errors.New("you are not the owner of this resource")
	ErrAlreadyCompleted = errors.New("microtask is already completed")
)

func ErrorToStatusCode(err error) int {
	switch {
	case errors.Is(err, ErrUnauthorized):
		return http.StatusUnauthorized
	case errors.Is(err, ErrForbidden):
		return http.StatusForbidden
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrInvalidInput):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
