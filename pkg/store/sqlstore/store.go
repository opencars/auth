package sqlstore

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/opencars/auth/pkg/store"
)

type Store struct {
	db *sqlx.DB

	tokenRepository     *TokenRepository
	blackListRepository *BlacklistRepository
}

func (s *Store) Token() store.TokenRepository {
	if s.tokenRepository != nil {
		return s.tokenRepository
	}

	s.tokenRepository = &TokenRepository{
		store: s,
	}

	return s.tokenRepository
}

func (s *Store) Blacklist() store.BlackListRepository {
	if s.blackListRepository != nil {
		return s.blackListRepository
	}

	s.blackListRepository = &BlacklistRepository{
		store: s,
	}

	return s.blackListRepository
}

func New(host string, port int, user, password, dbname string) (*Store, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	return &Store{
		db: db,
	}, nil
}
