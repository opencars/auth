package store

import (
	"github.com/opencars/auth/pkg/model"
)

type TokenRepository interface {
	FindByID(id string) (*model.Token, error)
}

type BlackListRepository interface {
	FindByIPv4(ipv4 string) (*model.BlackListItem, error)
}
