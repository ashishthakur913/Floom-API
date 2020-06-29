package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	httpCode int
	message  string `json:"message"`
}

func (e *Error) Error() error {
	return fmt.Errorf(e.message)
}

func (e *Error) GetMessage() string {
	return e.message
}

func (e *Error) GetHttpCode() int {
	return e.httpCode
}

func DuplicateError(message string) *Error {
	return &Error{
		httpCode: http.StatusTooManyRequests,
		message:  message,
	}
}

func NotFoundError(message string) *Error {
	return &Error{
		httpCode: http.StatusNotFound,
		message:  message,
	}
}

func InvalidArgumentError(message string) *Error {
	return &Error{
		httpCode: http.StatusBadRequest,
		message:  message,
	}
}

func InternalServerError(message string) *Error {
	return &Error{
		httpCode: http.StatusInternalServerError,
		message:  message,
	}
}

func UnauthorisedError(message string) *Error {
	return &Error{
		httpCode: http.StatusUnauthorized,
		message:  message,
	}
}
