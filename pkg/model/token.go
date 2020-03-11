package model

import (
	"time"
)

type Token struct {
	ID      string `json:"id" db:"id"`
	Name    string `json:"name" db:"name"`
	Enabled bool   `json:"enabled" db:"enabled"`
}

type Authorization struct {
	Token
	Status string    `json:"status"`
	Error  *string   `json:"error,omitempty"`
	IP     string    `json:"ip"`
	Time   time.Time `json:"timestamp"`
}
