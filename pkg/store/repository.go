package store

import (
	"github.com/opencars/auth/pkg/model"
)

// TokenRepository is responsible for tokens manipulation.
type TokenRepository interface {
	FindByID(id string) (*model.Token, error)
}

// BlackListRepository is responsible for manipulation of blacklisted items IP addresses.
type BlackListRepository interface {
	FindByIPv4(ipv4 string) (*model.BlackListItem, error)
}
