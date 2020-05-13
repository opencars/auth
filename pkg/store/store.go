package store

import (
	"errors"
)

var (
	// ErrRecordNotFound ...
	ErrRecordNotFound = errors.New("record not found")
)

// Store is an interface for communication with store.
type Store interface {
	Token() TokenRepository
}
