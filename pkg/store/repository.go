package store

import (
	"github.com/opencars/auth/pkg/model"
)

// Store is an interface for communication with tokens.
type TokenRepository interface {
	FindByID(id string) (*model.Token, error)
}
