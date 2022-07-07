package model

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	code    int
	message string
}

func NewHTTPError(code int, message string, args ...interface{}) error {
	return &HTTPError{
		code:    code,
		message: fmt.Sprintf(message, args...),
	}
}

func (e *HTTPError) Code() int {
	return e.code
}

func (e *HTTPError) Error() string {
	return e.message
}

var (
	ErrMessageInternalServerError = NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
)
