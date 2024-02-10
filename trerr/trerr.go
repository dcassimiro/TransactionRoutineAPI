package trerr

import (
	"errors"
	"fmt"
	"net/http"
)

// Error custom error type
type Error struct {
	HTTPCode int         `json:"-"`
	Message  string      `json:"message"`
	Detail   interface{} `json:"detail,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %v - message: %v - detail: %v", e.HTTPCode, e.Message, e.Detail)
}

// New creates a new error
func New(httpCode int, message string, detail interface{}) error {
	return &Error{
		HTTPCode: httpCode,
		Message:  message,
		Detail:   detail,
	}
}

// GetHTTPCode returns the error http code
func GetHTTPCode(err error) int {
	var e *Error
	if !errors.As(err, &e) {
		return http.StatusInternalServerError
	}
	return e.HTTPCode
}

// NewError creates a new struct Error
func NewError(httpCode int, message string, detail interface{}) *Error {
	return &Error{
		HTTPCode: httpCode,
		Message:  message,
		Detail:   detail,
	}
}
