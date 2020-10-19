package handler

import "net/http"

var (
	// ErrInvalidToken returned, if ApiKey/X-Api-Key is not valid.
	ErrInvalidToken = NewError(http.StatusUnauthorized, "auth.token.is_not_valid")
	// ErrTokenRevoked returned, if ApiKey/X-Api-Key was temporary disabled.
	ErrTokenRevoked = NewError(http.StatusUnauthorized, "auth.token.revoked")
	// ErrAccessDenied returned, if ip address is blacklisted.
	ErrAccessDenied = NewError(http.StatusForbidden, "auth.access_denied")
	// ErrUnhealthy returned, if application has an internal error.
	ErrUnhealthy = NewError(http.StatusInternalServerError, "system.unhealthy")
)

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// StatusError represents error with http status code.
type StatusError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error allows StatusError to satisfy the error interface.
func (se StatusError) Error() string {
	return se.Message
}

// Status returns our HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

// NewError creates new error instance.
func NewError(code int, message string) Error {
	return &StatusError{
		Code:    code,
		Message: message,
	}
}
