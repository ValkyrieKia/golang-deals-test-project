package util

import (
	"errors"
	"fmt"
	"net/http"
)

type ErrorType string

const (
	ErrBadRequest   ErrorType = "err/bad_request"
	ErrNotFound     ErrorType = "err/not_found"
	ErrInternal     ErrorType = "err/internal"
	ErrUnauthorized ErrorType = "err/unauthorized"
	ErrForbidden    ErrorType = "err/forbidden"
)

func (e ErrorType) GetHTTPStatus() int {
	switch e {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrUnauthorized:
		return http.StatusUnauthorized
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrForbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

type CommonError struct {
	Source     error
	Message    string
	Code       string
	HTTPStatus int
}

func (err *CommonError) Error() string {
	if err.Message == "" {
		return err.Source.Error()
	}
	return fmt.Sprintf("%v", err.Message)
}

func (err *CommonError) Unwrap() error {
	if err.Source == nil {
		return errors.New(err.Message)
	}
	return err.Source
}

// NewCommonError wrap an error object as `CommonError`.
func NewCommonError[T string | ErrorType](errObject error, errorType T, message string) *CommonError {
	var msg string
	var code string
	var httpStatus int

	msg = message

	switch v := any(errorType).(type) {
	case string:
		code = v
	case ErrorType:
		code = string(v)
		httpStatus = v.GetHTTPStatus()
	default:
		code = "err/unknown"
		httpStatus = http.StatusInternalServerError
	}

	if msg == "" {
		msg = code
	}

	errResult := &CommonError{
		Source:  errObject,
		Message: msg,
	}

	if code != "" {
		errResult.Code = code
	}

	if httpStatus != 0 {
		errResult.HTTPStatus = httpStatus
	}

	return errResult
}
