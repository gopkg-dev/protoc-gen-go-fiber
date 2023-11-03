package errors

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// UnknownCode is unknown code for error info.
	UnknownCode = 500
	// UnknownReason is unknown reason for error info.
	UnknownReason = "UnknownReason"
)

type Error struct {
	Code     int                    `json:"code"`               // http status code
	Reason   string                 `json:"reason"`             // error reason
	Message  string                 `json:"message"`            // error message
	Metadata map[string]interface{} `json:"metadata,omitempty"` // extra data
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// Is matches each error in the chain with the target value.
func (e *Error) Is(err error) bool {
	if se := new(Error); errors.As(err, &se) {
		return se.Reason == e.Reason
	}
	return false
}

// WithMetadata with an MD formed by the mapping of key, value.
func (e *Error) WithMetadata(md map[string]interface{}) *Error {
	err := *e
	err.Metadata = md
	return &err
}

// WithMessage with ....
func (e *Error) WithMessage(format string, args ...interface{}) *Error {
	_err := *e
	_err.Message = fmt.Sprintf(format, args...)
	return &_err
}

// New returns an error object for the code, message.
func New(code int, reason, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Reason:  reason,
	}
}

// Newf New(code fmt.Sprintf(format, a...))
func Newf(code int, reason, format string, a ...interface{}) *Error {
	return New(code, reason, fmt.Sprintf(format, a...))
}

// FromError try to convert an error to *Error.
// It supports wrapped errors.
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	if se := new(Error); errors.As(err, &se) {
		return se
	}
	return New(UnknownCode, UnknownReason, err.Error())
}
