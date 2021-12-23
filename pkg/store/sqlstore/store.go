package sqlstore

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/opencars/auth/pkg/domain"
)

// Store is an implementation of store.Store interface based on PostgreSQL.
type Store struct {
	db *sqlx.DB

	tokenRepository     *TokenRepository
	blackListRepository *BlacklistRepository
}

// Token returns repository responsible for tokens.
func (s *Store) Token() domain.TokenRepository {
	if s.tokenRepository != nil {
		return s.tokenRepository
	}

	s.tokenRepository = &TokenRepository{
		store: s,
	}

	return s.tokenRepository
}

// Blacklist returns repository responsible for blacklisted items.
func (s *Store) Blacklist() domain.BlackListRepository {
	if s.blackListRepository != nil {
		return s.blackListRepository
	}

	s.blackListRepository = &BlacklistRepository{
		store: s,
	}

	return s.blackListRepository
}

// New returns newly allocated store.
func New(host string, port int, user, password, dbname, sslmode string) (*Store, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s", host, port, user, dbname, sslmode, password)

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	return &Store{
		db: db,
	}, nil
}
