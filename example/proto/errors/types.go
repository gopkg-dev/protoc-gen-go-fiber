package errors

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const (
	DefaultBadRequestReason            = "BadRequest"
	DefaultUnauthorizedReason          = "Unauthorized"
	DefaultForbiddenReason             = "Forbidden"
	DefaultNotFoundReason              = "NotFound"
	DefaultMethodNotAllowedReason      = "MethodNotAllowed"
	DefaultTooManyRequestsReason       = "TooManyRequests"
	DefaultRequestEntityTooLargeReason = "RequestEntityTooLarge"
	DefaultInternalServerErrorReason   = "InternalServerError"
	DefaultConflictReason              = "Conflict"
	DefaultRequestTimeoutReason        = "RequestTimeout"
)

// BadRequest generates a 400 error.
func BadRequest(reason, format string, a ...interface{}) *Error {
	if reason == "" {
		reason = DefaultBadRequestReason
	}
	return &Error{
		Code:    fiber.StatusBadRequest,
		Message: fmt.Sprintf(format, a...),
		Reason:  reason,
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(reason, format string, a ...interface{}) *Error {
	if reason == "" {
		reason = DefaultUnauthorizedReason
	}
	return &Error{
		Code:    fiber.StatusUnauthorized,
		Message: fmt.Sprintf(format, a...),
		Reason:  reason,
	}
}

// Forbidden generates a 403 error.
func Forbidden(reason, format string, a ...interface{}) *Error {
	if reason == "" {
		reason = DefaultForbiddenReason
	}
	return &Error{
		Code:    fiber.StatusForbidden,
		Message: fmt.Sprintf(format, a...),
		Reason:  reason,
	}
}

// NotFound generates a 404 error.
func NotFound(reason, format string, a ...interface{}) *Error {
	if reason == "" {
		reason = DefaultNotFoundReason
	}
	return &Error{
		Code:    fiber.StatusNotFound,
		Message: fmt.Sprintf(format, a...),
		Reason:  reason,
	}
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(reason, format string, a ...interface{}) error {
	if reason == "" {
		reason = DefaultMethodNotAllowedReason
	}
	return &Error{
		Code:    fiber.StatusMethodNotAllowed,
		Message: fmt.Sprintf(format, a...),
		Reason:  reason,
	}
}

// TooManyRequests generates a 429 error.
func TooManyRequests(reason, format string, a ...interface{}) error {
	if reason == "" {
		reason = DefaultTooManyRequestsReason
	}
	return &Error{
		Code:    fiber.StatusTooManyRequests,
		Message: fmt.Sprintf(format, a...),
		Reason:  reason,
	}
}

// Timeout generates a 408 error.
func Timeout(reason, format string, a ...interface{}) error {
	if reason == "" {
		reason = DefaultRequestTimeoutReason
	}
	return &Error{
		Code:    fiber.StatusRequestTimeout,
		Message: fmt.Sprintf(format, a...),
		Reason:  reason,
	}
}

// Conflict generates a 409 error.
func Conflict(reason, format string, a ...interface{}) error {
	if reason == "" {
		reason = DefaultConflictReason
	}
	return &Error{
		Code:    fiber.StatusConflict,
		Message: fmt.Sprintf(format, a...),
		Reason:  reason,
	}
}

// RequestEntityTooLarge generates a 413 error.
func RequestEntityTooLarge(reason, format string, a ...interface{}) error {
	if reason == "" {
		reason = DefaultRequestEntityTooLargeReason
	}
	return &Error{
		Code:    fiber.StatusRequestEntityTooLarge,
		Message: fmt.Sprintf(format, a...),
		Reason:  reason,
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(reason, format string, a ...interface{}) error {
	if reason == "" {
		reason = DefaultInternalServerErrorReason
	}
	return &Error{
		Code:    fiber.StatusInternalServerError,
		Message: fmt.Sprintf(format, a...),
		Reason:  reason,
	}
}
