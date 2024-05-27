package errors

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	BadRequest          = errors.New("bad request")
	NotFound            = errors.New("not found")
	Conflict            = errors.New("conflict")
	Forbidden           = errors.New("forbidden")
	Unauthorized        = errors.New("unauthorized")
	InternalServerError = errors.New("internal server error")
)

type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
	Unwrap() error
}

type RestError struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  error       `json:"error,omitempty"`
	ErrCauses interface{} `json:"-"`
}

func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - causes: %v", e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e RestError) Status() int {
	return e.ErrStatus
}

func (e RestError) Causes() interface{} {
	return e.ErrCauses
}

func (e RestError) Unwrap() error {
	return e.ErrError
}

func NewRestError(status int, err error, causes interface{}) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCauses: causes,
	}
}

func NewBadRequestError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  BadRequest,
		ErrCauses: causes,
	}
}

func NewNotFoundError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusNotFound,
		ErrError:  NotFound,
		ErrCauses: causes,
	}
}

func NewConflictError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusConflict,
		ErrError:  Conflict,
		ErrCauses: causes,
	}
}

func NewForbiddenError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusForbidden,
		ErrError:  Forbidden,
		ErrCauses: causes,
	}
}

func NewUnauthorizedError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusUnauthorized,
		ErrError:  Unauthorized,
		ErrCauses: causes,
	}
}

func NewInternalServerError(causes interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  InternalServerError,
		ErrCauses: causes,
	}
}
