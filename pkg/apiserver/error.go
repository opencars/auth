package apiserver

import (
	"net/http"
)

var (
	InvalidToken = NewError(http.StatusUnauthorized, "auth.token.is_not_valid")
	TokenRevoked = NewError(http.StatusUnauthorized, "auth.token.revoked")
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
