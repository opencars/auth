package model

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Token is a domain model of token.
type Token struct {
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Secret    string    `json:"secret,omitempty" db:"secret"`
	Name      string    `json:"name" db:"name"`
	Enabled   bool      `json:"enabled" db:"enabled"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewToken(userID, name string) *Token {
	return &Token{
		ID:        uuid.NewV4().String(),
		UserID:    userID,
		Secret:    GenerateSecret(16),
		Name:      name,
		Enabled:   true,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func (t *Token) ResetSecret() {
	t.Secret = GenerateSecret(16)
}

func (t *Token) ClearSecret() {
	t.Secret = ""
}

// AuthStatus is status of authorization.
type AuthStatus string

const (
	// AuthStatusSucceed represents successful authorization.
	AuthStatusSucceed AuthStatus = "succeed"
	// AuthStatusFailed represents failed authorization.
	AuthStatusFailed AuthStatus = "failed"
)

// Authorization represents authorization
type Authorization struct {
	Token
	Status AuthStatus `json:"status"`
	Error  *string    `json:"error,omitempty"`
	IP     string     `json:"ip"`
	Time   time.Time  `json:"timestamp"`
}

func GenerateSecret(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}

	return hex.EncodeToString(b)
}
