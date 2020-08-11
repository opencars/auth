package apiserver

import (
	"net/http"
)

var (
	ErrInvalidToken = NewError(http.StatusUnauthorized, "auth.token.is_not_valid")
	ErrTokenRevoked = NewError(http.StatusUnauthorized, "auth.token.revoked")
	ErrAccessDenied = NewError(http.StatusForbidden, "auth.access_denied")
	ErrUnhealthy    = NewError(http.StatusInternalServerError, "system.unhealthy")
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

func NewError(code int, message string) Error {
	return Error{
		Code:    code,
		Message: message,
	}
}
