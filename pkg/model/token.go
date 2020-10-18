package model

import (
	"time"
)

type Token struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Enabled bool   `json:"enabled" db:"enabled"`
}

type AuthStatus string

const (
	AuthStatusSucceed = "succeed"
	AuthStatusFailed  = "failed"
)

type Authorization struct {
	Token
	Status AuthStatus `json:"status"`
	Error  *string    `json:"error,omitempty"`
	IP     string     `json:"ip"`
	Time   time.Time  `json:"timestamp"`
}
