package model

import (
	"time"
)

// Token is a domain model of token.
type Token struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Enabled bool   `json:"enabled" db:"enabled"`
}

// AuthStatus is status of authorization.
type AuthStatus string

const (
	// AuthStatusSucceed represents successful authorization.
	AuthStatusSucceed = "succeed"
	// AuthStatusFailed represents failed authorization.
	AuthStatusFailed = "failed"
)

// Authorization represents authorization
type Authorization struct {
	Token
	Status AuthStatus `json:"status"`
	Error  *string    `json:"error,omitempty"`
	IP     string     `json:"ip"`
	Time   time.Time  `json:"timestamp"`
}
