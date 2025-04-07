package models

import "errors"

var (
	ErrServiceError = errors.New("service error")
	ErrBadRequest   = errors.Join(ErrServiceError, errors.New("bad request"))
	ErrNotFound     = errors.Join(ErrServiceError, errors.New("not found"))
	ErrConflict     = errors.Join(ErrServiceError, errors.New("conflict"))
	ErrUnauthorized = errors.Join(ErrServiceError, errors.New("unauthorized"))
	ErrForbidden    = errors.Join(ErrServiceError, errors.New("forbidden"))
	ErrServerError  = errors.Join(ErrServiceError, errors.New("server error"))
)
