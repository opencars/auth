package domain

import (
	"context"

	"github.com/opencars/auth/pkg/domain/command"
	"github.com/opencars/auth/pkg/domain/model"
	"github.com/opencars/auth/pkg/domain/query"
)

type UserService interface {
	CreateToken(context.Context, *command.CreateToken) (*model.Token, error)
	TokenDetails(context.Context, *query.TokenDetails) (*model.Token, error)
	ResetToken(context.Context, *command.ResetToken) (*model.Token, error)
	DeleteToken(context.Context, *command.DeleteToken) error
	ListTokens(context.Context, *query.ListTokens) ([]model.Token, error)
}

type SessionChecker interface {
	CheckSession(ctx context.Context, sessionToken string) (*model.User, error)
}

// TokenRepository is responsible for tokens manipulation.
type TokenRepository interface {
	Create(ctx context.Context, token *model.Token) error
	Update(ctx context.Context, token *model.Token) error
	FindByID(ctx context.Context, id string) (*model.Token, error)
	FindBySecret(ctx context.Context, secret string) (*model.Token, error)
	DeleteByID(ctx context.Context, id string) error
	List(ctx context.Context, q *query.ListTokens) ([]model.Token, error)
}

// BlackListRepository is responsible for manipulation of blacklisted items IP addresses.
type BlackListRepository interface {
	FindByIPv4(ipv4 string) (*model.BlackListItem, error)
}

// Store is an interface for communication with store.
type Store interface {
	Token() TokenRepository
	Blacklist() BlackListRepository
}
